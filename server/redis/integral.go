package redis

import (
	"fmt"
	"strconv"
	"yyds-pro/core/const"
	"yyds-pro/trace"
)

const (
	filterScript = `
	--利用redis的hash结构，存储key所对应令牌桶的上次获取时间和上次获取后桶中令牌数量
	local bucket_info = redis.call("HMGET", KEYS[1], "last_time", "current_token_num");
	local last_time = tonumber(bucket_info[1]);
	local current_token_num = tonumber(bucket_info[2]);
	redis.replicate_commands();
	redis.call("pexpire", KEYS[1], 1000);
	local now = redis.call("TIME");
	redis.call("SET", "now", tonumber(now[1]));	

	--tonumber是将value转换为数字，此步是取出桶中最大令牌数、生成令牌的速率(每秒生成多少个)、当前时间
	local max_token_num = 30;
	local token_rate = 20;
	local current_time = tonumber(now[1]);
	--reverse_time 即多少毫秒生成一个令牌
	local reverse_time = 1000/token_rate;
	local past_time
	local reverse_token_num
	--如果current_token_num不存在则说明令牌桶首次获取或已过期，即说明它是满的
	if current_token_num == nil then
		current_token_num = max_token_num;
		last_time = current_time;
	else
		--计算出距上次获取已过去多长时间
		past_time = current_time - last_time;
		--在这一段时间内可产生多少令牌
		reverse_token_num = math.floor(past_time/reverse_time);
		current_token_num = current_token_num + reverse_token_num;
		last_time = reverse_time * reverse_token_num + last_time;
		if current_token_num > max_token_num then
			current_token_num = max_token_num;
		end
	end
	if (current_token_num > 0) then
		current_token_num = current_token_num -1;
	end
	-- 将最新得出的令牌获取时间和当前令牌数量进行存储,并设置过期时间
	redis.call('HMSET', KEYS[1], "last_time", last_time, "current_token_num", current_token_num);
	return current_token_num
`

	integralScript = `
	-- KEYS[1]: 用户去重的hash key，用于检测是否已经抢过积分
	-- KEYS[2]:	用户名
	-- KEYS[3]: 未消费的积分队列
	-- KEYS[4]: 已消费的积分队列
	-- 如果用户已抢过积分，则返回nil
	-- 0:用户已经抢过积分，不能再抢第二次  -1:积分已经抢完   1:获取积分成
	-- 检查用户是否已经抢过积分
	locala userReward = redis.call("SISMEMBER", KEYS[1], KEYS[2]);
	if (userReward == 1) 
	then
		return 0;
	end
	-- 判断积分是否已经抢完了
	local len = redis.call("llen", KEYS[3]);
	if (len == 0)
	then
		return -1;
	end
	-- 还有剩余，从未消费的积分队列中取一个id
	local id = redis.call("RPOP", KEYS[3]);
	-- 将取走的积分id放入已消费的队列中
	redis.call("RPUSH", KEYS[4], id);
	--将已抢过积分的user放入redis中
	redis.call("SADD",KEYS[1], KEYS[2]);
	-- 返回id
	return tonumber(id);
	`
)

//
//  SetIntegralList
//  @Description: 未消费积分队列放入redis
//  @param key
//  @param id
//  @return int64
//  @return error
//
func SetIntegralList(ctx *trace.Trace, key string, id string) (i int64, err error) {
	i, err = DefaultRedisClient.LPush(key, id).Result()
	if err != nil {
		ctx.Redis.Error = err
		return
	}
	return
}

//
//  SetIntegralInfo
//  @Description: 积分详情存入redis
//  @param key
//  @param id
//  @param money
//  @return bool
//  @return error
//
func SetIntegralInfo(ctx *trace.Trace, key string, id, money string) (flag bool, err error) {
	flag, err = DefaultRedisClient.HSet(key, id, money).Result()
	if err != nil {
		ctx.Redis.Error = err
		return
	}
	return
}

//
//  GetFilterBucket
//  @Description: 限流判断
//  @return flag
//  @return err
//
func GetFilterBucket(ctx *trace.Trace) (flag bool, err error) {
	//限流处理
	res, err := FilterEval(ctx, filterScript, []string{_const.Key1})
	if err != nil {
		return
	}
	if res.(int64) == 0 {
		fmt.Println("令牌已用完")
		return
	}
	if res.(int64) >= 1 {
		flag = true
		return
	}
	return
}

//
//  CacheGetReward
//  @Description: 抢积分逻辑
//  @param key1
//  @param username
//  @param key2
//  @param key3
//  @return int
//
func DoIntegral(ctx *trace.Trace, key1, key2, key3, key4 string) (msg string, val int, err error) {
	var str string
	//执行抢积分的lua脚本
	res, err := EvalSHA(ctx, integralScript, []string{key1, key2, key3, key4})
	if err != nil {
		return
	}
	switch {
	case res.(int64) == 0:
		msg = "用户已经抢过积分，不能再抢"
		return
	case res.(int64) == -1:
		msg = "积分已经抢完"
		return
	case res.(int64) > 0:
		msg = "您已抢到积分"
		id := int(res.(int64))
		str, err = GetIntegralByIdFromCache(ctx, id)
		if err != nil {
			return
		}
		val, _ = strconv.Atoi(str)
		return
	default:
		msg = "发生预期之外的错误"
	}
	return
}
