package serviceimpl

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"yyds-pro/core/const"
	"yyds-pro/log"
	"yyds-pro/model"
	"yyds-pro/repository"
	"yyds-pro/repository/repoimpl"
	redisRate "yyds-pro/server/redis"
	"yyds-pro/trace"
)

type RateLimiterService struct {
	RateLimiterRepo repository.RateLimiterRepository
}

func NewRateLimiterService() RateLimiterService {
	return RateLimiterService{
		RateLimiterRepo: repoimpl.NewRateLimiterRepository(),
	}
}

//
//  InitIntegralDb
//  @Description: 初始化积分相关db和redis
//  @receiver r
//  @param ctx
//  @param integral
//  @return err
//
func (r RateLimiterService) InitIntegralDb(ctx *trace.Trace, integral model.Integral) (err error) {
	for i := 0; integral.RemainCount > 0; i++ {
		v := GrabIntegral(&integral)
		rand.Seed(time.Now().UnixNano())
		randomId := rand.Intn(100) // 生成0~9的随机数
		id := randomId + v
		//把积分初始化信息放入mysql
		err = r.RateLimiterRepo.InitIntegralDb(ctx, id, v)
		if err != nil {
			log.L.ErrorCtx(ctx, err)
			return
		}
		//放入未消费的积分队列(list),存放的是id
		_, err = redisRate.SetIntegralList(ctx, _const.NIntegralList, strconv.Itoa(id))
		if err != nil {
			log.L.ErrorCtx(ctx, err)
			return
		}
		//把积分具体信息放入redis，用hash表储存(id+money)
		_, err = redisRate.SetIntegralInfo(ctx, _const.IntegralInfo, strconv.Itoa(id), strconv.Itoa(v))
		if err != nil {
			log.L.ErrorCtx(ctx, err)
			return
		}
	}
	return

}

//
//  GetIntegralByUserInfo
//  @Description: 抢积分限流具体逻辑
//  @receiver r
//  @param ctx
//  @param userInfo
//  @return msg
//  @return val
//  @return err
//
func (r RateLimiterService) GetIntegralByUserInfo(ctx *trace.Trace, userInfo model.UserInfo) (msg string, val int, err error) {
	//限流处理
	flag, err := redisRate.GetFilterBucket(ctx)
	if err != nil {
		return
	}
	if flag {
		//抢积分具体业务逻辑
		msg, val, err = redisRate.DoIntegral(ctx, _const.UserKey, userInfo.UserName, _const.NIntegralList, _const.YIntegralList)
		if err != nil {
			return
		}
		return
	}
	return
}

//
//  GrabReward
//  @Description: 随机生成integral返回
//  @param integral
//  @return int
//
func GrabIntegral(integral *model.Integral) int {
	if integral.RemainCount <= 0 {
		panic("RemainCount <= 0")
	}
	//最后一个
	if integral.RemainCount-1 == 0 {
		amount := integral.RemainIntegral
		integral.RemainCount = 0
		integral.RemainIntegral = 0
		return amount
	}
	//是否可以直接0.01
	if (integral.RemainIntegral / integral.RemainCount) == 1 {
		fmt.Println(integral.RemainIntegral / integral.RemainCount)
		amount := 1
		integral.RemainIntegral -= amount
		integral.RemainCount--
		return amount
	}

	//最大可领积分 = 剩余积分的平均值x2 = (剩余积分 / 剩余数量) * 2
	//领取积分范围 = 0.01 ~ 最大可领积分
	maxAmount := (integral.RemainIntegral / integral.RemainCount) * 2
	rand.Seed(time.Now().UnixNano())
	amount := rand.Intn(maxAmount)
	for amount == 0 {
		//防止零
		amount = rand.Intn(maxAmount)
	}
	integral.RemainIntegral -= amount
	//防止剩余积分负数
	if integral.RemainIntegral < 0 {
		amount += integral.RemainIntegral
		integral.RemainIntegral = 0
		integral.RemainCount = 0
	} else {
		integral.RemainCount--
	}
	return amount
}

//
//  UpdateUserIntegral
//  @Description: 更新用户积分
//  @receiver r
//  @param ctx
//  @param val
//  @return error
//
func (r RateLimiterService) UpdateUserIntegral(ctx *trace.Trace, val int, userName string) error {
	return r.RateLimiterRepo.UpdateUserIntegral(ctx, val, userName)
}
