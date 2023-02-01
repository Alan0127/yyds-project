package kafka

import (
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"strings"
	"time"
)

const (
	kafkaTimeOut = time.Second * 5

	// kafka生产者发送信息方式
	Sync  = "Sync"
	Async = "Async"
)

var (
	kafkaAddressError = errors.New("kafka address is error")
)

func CreateProducer(address string, topic string, duration time.Duration, syncOrAsync string) (AbsKafkaProducer, error) {
	var producer AbsKafkaProducer
	switch syncOrAsync {
	case Sync:
		producer = &SyncKafkaProducer{}
	case Async:
		producer = &AsyncKafkaProducer{}
	default:
		return nil, errors.New("kafka mode error please choose sync or async")
	}
	err := producer.NewKafkaProducer(address, topic, duration)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

// kafka生产者配置struct
type KafkaConfig struct {
	addressList []string       //地址列表
	topic       string         //kafka topic
	config      *sarama.Config //kafka配置信息
}

// 创建kafka生产者config返回
func NewProducer(address, topic string, duration time.Duration) (*KafkaConfig, error) {
	addressList := strings.Split(address, ",")
	if len(addressList) < 1 || addressList[0] == "" {
		return nil, errors.New("kafka addr error")
	}
	//配置producer
	sendConfig := sarama.NewConfig()
	sendConfig.Producer.Return.Successes = true
	sendConfig.Producer.Timeout = kafkaTimeOut
	if duration != 0 {
		sendConfig.Producer.Timeout = duration
	}
	return &KafkaConfig{
		addressList: addressList,
		topic:       topic,
		config:      sendConfig,
	}, nil
}

// KafkaProducer 生产者抽象
type AbsKafkaProducer interface {
	NewKafkaProducer(address, topic string, duration time.Duration) error
	Send(value []byte) error
}

// SyncKafkaProducer 同步kafka生产者
type SyncKafkaProducer struct {
	KafkaConfig *KafkaConfig
}

// 同步发送生产者实例
func (k *SyncKafkaProducer) NewKafkaProducer(address, topic string, duration time.Duration) (err error) {
	if len(address) == 0 {
		return kafkaAddressError
	}
	k.KafkaConfig, err = NewProducer(address, topic, duration)
	return
}

// send同步发送消息
func (k *SyncKafkaProducer) Send(value []byte) (err error) {
	if k == nil || k.KafkaConfig == nil || value == nil {
		return errors.New("kafka init error")
	}
	p, err := sarama.NewSyncProducer(k.KafkaConfig.addressList, k.KafkaConfig.config)
	if err != nil {
		err = errors.New("create SyncProducer error")
		return
	}
	defer p.Close()
	msg := &sarama.ProducerMessage{
		Topic: k.KafkaConfig.topic,
		Value: sarama.ByteEncoder(value),
	}
	_, _, err = p.SendMessage(msg)
	if err != nil {
		err = errors.New("send kafka msg error")
		return
	}
	return
}

// AsyncKafkaProducer kafka异步生产者
type AsyncKafkaProducer struct {
	KafkaConfig *KafkaConfig         //共有的kafka生产者配置在这个里面
	producer    sarama.AsyncProducer //异步生产者
	isClose     chan struct{}        // 监听producer是否可以关闭
}

//异步发送生产者实例
func (k *AsyncKafkaProducer) NewKafkaProducer(address, topic string, duration time.Duration) (err error) {
	if len(address) == 0 {
		return kafkaAddressError
	}
	//配置异步发送启动参数
	k.KafkaConfig, err = NewProducer(address, topic, duration)
	if err != nil {
		return
	}
	k.isClose = make(chan struct{}, 2)
	//启动kafka异步producer
	k.Run()
	return nil
}

func (k *AsyncKafkaProducer) Send(value []byte) (err error) {
	//如果实例或者配置为空，直接返回。如果发送数据为空也直接返回
	if k == nil || k.KafkaConfig == nil || value == nil {
		err = errors.New("kafka init error")
		return err
	}
	//封装消息
	msg := &sarama.ProducerMessage{
		Topic: k.KafkaConfig.topic,
		Value: sarama.ByteEncoder(value),
	}
	//一般不会出现，producer实例为空时，表示异步创建producer失败
	if k.producer == nil {
		k.Run()
	}
	select {
	//producer出现error时重新启动
	case <-k.isClose:
		if k.producer != nil {
			// 收到可以关闭producer的消息isClose,关闭producer并重启
			k.producer.Close()
			k.Run()
		}
		return
	default:
		k.producer.Input() <- msg
	}
	return
}

// kafka异步生产者
func (k *AsyncKafkaProducer) Run() {
	if k.KafkaConfig == nil || k == nil {
		return
	}
	//创建异步producer
	producer, err := sarama.NewAsyncProducer(k.KafkaConfig.addressList, k.KafkaConfig.config)
	//如果创建失败主动置空k.producer，否则producer不为空，在重启的时候k.producer是会有值的
	if err != nil {
		k.isClose <- struct{}{}
		k.producer = nil
		fmt.Println("create NewAsyncProducer error!")
		return
	}
	if producer == nil {
		k.isClose <- struct{}{}
		k.producer = nil
		fmt.Println("create NewAsyncProducer error!")
		return
	}
	//如果创建成功为实例k的prodcer赋值
	k.producer = producer
	go func(p sarama.AsyncProducer) {
		errChan := p.Errors()
		success := p.Successes()
		for {
			select {
			//出现错误
			case rc := <-errChan:
				if rc != nil {
					//标记producer出现error，在send时会监听到这个标记
					k.isClose <- struct{}{}
					fmt.Println("send kafka data error")
				}
				return
			case res := <-success:
				data, _ := res.Value.Encode()
				fmt.Printf("发送成功，value=%s \n", string(data))
			}
		}
	}(producer)
}
