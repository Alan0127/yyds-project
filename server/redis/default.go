package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"yyds-pro/model"
)

var DefaultRedisClient *redis.Client

func InitRedis(config model.AppConfig) (err error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		Network:  "tcp",
		PoolSize: 50,
	})
	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}
	DefaultRedisClient = client
	return
}

func PipelineGetHashField(keyList []string, filed string) (err error, valList []string) {
	pipeClient := DefaultRedisClient.Pipeline()
	for _, key := range keyList {
		pipeClient.HGet(key, filed)
	}
	res, err := pipeClient.Exec()
	if err != nil {
		if err != redis.Nil {
			return
		}
		/********** ！！！！！！！！！！*************/
		// 注意这里如果某一次获取时出错（常见的redis.Nil），返回的err即不为空
		// 如果需要处理redis.Nil为默认值，此处不能直接return
	}
	for _, cmdRes := range res {
		var val string
		// 此处断言类型为在for循环内执行的命令返回的类型,上面HGet返回的即为*redis.StringCmd类型
		// 处理方式和直接调用同样处理即可
		cmd, ok := cmdRes.(*redis.StringCmd)
		if ok {
			val, err = cmd.Result()
			if err != nil {
				fmt.Println(err)
			}
		}
		valList = append(valList, val)
	}
	return
}

func PipelineDelHashField(keyList []string, filed string) (err error, valList []string) {
	pipeClient := DefaultRedisClient.Pipeline()
	for _, key := range keyList {
		pipeClient.HDel(key, filed)
	}
	res, err := pipeClient.Exec()
	if err != nil {
		if err != redis.Nil {
			return
		}
	}
	for _, cmdRes := range res {
		var val string
		cmd, ok := cmdRes.(*redis.StringCmd)
		if ok {
			val, err = cmd.Result()
			if err != nil {
				fmt.Println(err)
			}
		}
		valList = append(valList, val)
	}
	return
}
