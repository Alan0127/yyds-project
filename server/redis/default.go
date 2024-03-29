package redis

import (
	"github.com/go-redis/redis"
	"strconv"
	"time"
	_const "yyds-pro/core/const"
	"yyds-pro/log"
	"yyds-pro/model"
	"yyds-pro/trace"
)

var DefaultRedisClient *redis.Client

//初始化redis连接
func InitRedis(config model.AppConfig) (err error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		Network:  "tcp",
		PoolSize: 50,
	})
	if _, err = client.Ping().Result(); err != nil {
		panic(err)
	}
	DefaultRedisClient = client
	return
}

//
//  HashGetWithCtx
//  @Description: hashSet
//  @param ctx
//  @param filed
//  @param key
//  @return res
//  @return err
//
func HashGetWithCtx(ctx *trace.Trace, filed, key string) (res string, err error) {
	res, err = DefaultRedisClient.HGet(filed, key).Result()
	ctx.Redis.Flag = true
	if err != nil {
		ctx.Redis.Flag = false
		ctx.Redis.Error = err
	}
	return
}

func HashSetWithContext(ctx *trace.Trace, filed, key string, value interface{}) (err error) {
	_, err = DefaultRedisClient.HSet(filed, key, value).Result()
	ctx.Redis.Flag = true
	if err != nil {
		ctx.Redis.Flag = false
		ctx.Redis.Error = err
	}
	return
}

//
//  PipelineSetHashField
//  @Description: 管道批量set value
//  @param ctx
//  @param keyList
//  @param filed
//  @return err
//  @return valList
//
func PipelineSetHashField(ctx trace.Trace, keymap map[string]interface{}, filed string) (err error, valList []string) {
	ctx.Redis.Flag = true
	var errList []error
	pipeClient := DefaultRedisClient.Pipeline()
	for key, val := range keymap {
		pipeClient.HSet(key, filed, val)
	}
	res, err := pipeClient.Exec()
	if err != nil {
		if err != redis.Nil {
			ctx.Redis.Error = err
			ctx.Redis.Flag = false
			return
		}
	}
	for _, cmdRes := range res {
		var val string
		// 此处断言类型为在for循环内执行的命令返回的类型,上面HGet返回的即为*redis.StringCmd类型
		// 处理方式和直接调用同样处理即可
		cmd, ok := cmdRes.(*redis.StringCmd)
		if ok {
			val, err = cmd.Result()
			if err != nil {
				errList = append(errList, err)
			}
		}
		valList = append(valList, val)
	}
	ctx.Redis.Error = errList
	return
}

//
//  PipelineGetHashField
//  @Description: 使用管道批量获取value
//  @param keyList
//  @param filed
//  @return err
//  @return valList
//
func PipelineGetHashField(ctx trace.Trace, keyList []string, filed string) (err error, valList []string) {
	ctx.Redis.Flag = true
	var errList []error
	pipeClient := DefaultRedisClient.Pipeline()
	for _, key := range keyList {
		pipeClient.HGet(key, filed)
	}
	res, err := pipeClient.Exec()
	if err != nil {
		if err != redis.Nil {
			ctx.Redis.Error = err
			ctx.Redis.Flag = false
			return
		}
	}
	for _, cmdRes := range res {
		var val string
		// 此处断言类型为在for循环内执行的命令返回的类型,上面HGet返回的即为*redis.StringCmd类型
		// 处理方式和直接调用同样处理即可
		cmd, ok := cmdRes.(*redis.StringCmd)
		if ok {
			val, err = cmd.Result()
			if err != nil {
				errList = append(errList, err)
			}
		}
		valList = append(valList, val)
	}
	ctx.Redis.Error = errList
	return
}

//
//  PipelineDelHashField
//  @Description: 使用管道批量删除value
//  @param keyList
//  @param filed
//  @return err
//  @return valList
//
func PipelineDelHashField(ctx trace.Trace, keyList []string, filed string) (err error, valList []string) {
	ctx.Redis.Flag = true
	var errList []error
	pipeClient := DefaultRedisClient.Pipeline()
	for _, key := range keyList {
		pipeClient.HDel(key, filed)
	}
	res, err := pipeClient.Exec()
	if err != nil {
		if err != redis.Nil {
			ctx.Redis.Error = err
			return
		}
	}
	for _, cmdRes := range res {
		var val string
		cmd, ok := cmdRes.(*redis.StringCmd)
		if ok {
			val, err = cmd.Result()
			if err != nil {
				errList = append(errList, err)
			}
		}
		valList = append(valList, val)
	}
	ctx.Redis.Error = errList
	return
}

//
//  SetRedisCtx
//  @Description: redis set string operation
//  @param ctx
//  @param key
//  @param value
//  @param expireTime
//  @return err
//
func SetRedisCtx(ctx *trace.Trace, key, value string, expireTime time.Duration) (err error) {
	_, err = DefaultRedisClient.Set(key, value, expireTime).Result()
	ctx.Redis.Flag = true
	if err != nil {
		ctx.Redis.Error = err
		ctx.Redis.Flag = false
		return
	}
	return
}

//
//  GetRedisCtx
//  @Description: redis get string operation
//  @param ctx
//  @param key
//  @return err
//
func GetRedisCtx(ctx *trace.Trace, key string) (err error) {
	_, err = DefaultRedisClient.Get(key).Result()
	ctx.Redis.Flag = true
	if err != nil {
		ctx.Redis.Error = err
		ctx.Redis.Flag = false
		return
	}
	return
}

//
//  FilterEval
//  @Description: 执行限流lua脚本
//  @param sha
//  @param args
//  @return val
//  @return err
//
func FilterEval(ctx *trace.Trace, sha string, args []string) (val interface{}, err error) {
	val, err = DefaultRedisClient.Eval(sha, args).Result()
	if err != nil {
		log.L.ErrorCtx(ctx, err)
		ctx.Redis.Error = err
		return
	}
	return
}

//
//  EvalSHA
//  @Description: 执行抢积分lua脚本
//  @param sha
//  @param args
//  @return interface{}
//  @return error
//
func EvalSHA(ctx *trace.Trace, sha string, args []string) (val interface{}, err error) {
	val, err = DefaultRedisClient.Eval(sha, args).Result()
	if err != nil {
		log.L.ErrorCtx(ctx, err)
		ctx.Redis.Error = err
		return
	}
	return
}

//
//  GetIntegralByidFromCache
//  @Description: 查询缓存的积分
//  @param id
//  @return val
//  @return err
//
func GetIntegralByIdFromCache(ctx *trace.Trace, id int) (val string, err error) {
	val, err = DefaultRedisClient.HGet(_const.IntegralInfo, strconv.Itoa(id)).Result()
	if err != nil {
		ctx.Redis.Error = err
	}
	return
}

//
//  RollingWindowLimiter
//  @Description: 基于redis的滑动窗口限流
//  @param queueName
//  @param count
//  @param timeWindow
//  @return bool
//
func RollingWindowLimiter(queueName string, count uint, timeWindow int64) bool {
	currTime := time.Now().Unix()
	length := uint(ListLen(queueName))
	if length < count {
		ListLpush(queueName, currTime)
		return true
	}
	//队列满了,取出最早访问的时间
	earlyTime, _ := strconv.ParseInt(GetElemByIndex(queueName, int64(length)-1), 10, 64)
	//说明最早期的时间还在时间窗口内,还没过期,所以不允许通过
	if currTime-earlyTime <= timeWindow {
		return false
	} else {
		//说明最早期的访问应该过期了,去掉最早期的
		ListRPop(queueName)
		ListLPush(queueName, currTime)
	}
	return true
}

func ListLen(key string) int64 {
	res := DefaultRedisClient.LLen(key)
	return res.Val()
}

func ListLpush(key string, currTime int64) {
	DefaultRedisClient.LPush(key, currTime)
}

func GetElemByIndex(key string, len int64) string {
	res := DefaultRedisClient.LIndex(key, len)
	return res.Val()
}

func ListRPop(key string) {
	DefaultRedisClient.LPop(key)
}

func ListLPush(key string, val int64) {
	DefaultRedisClient.LPush(key, val)
}
