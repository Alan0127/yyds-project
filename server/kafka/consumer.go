package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

//
//  GetConsumerData
//  @Description: kafka消费者
//  @param ch
//  @param addr
//  @param topic
//  @return res
//  @return err
//
func GetConsumerData(ch chan struct{}, addr, topic string) (res []byte, err error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.Retry.Max = 3
	consumer, err := sarama.NewConsumer([]string{addr}, config)
	if err != nil {
		return
	}
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		return
	}
	defer consumer.Close()
	for partition := range partitionList {
		partitionConsumer, err0 := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetOldest)
		if err0 != nil {
			err = err0
			return
		}
		defer partitionConsumer.Close()
		for {
			select {
			case msg := <-partitionConsumer.Messages():
				fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
					msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
				res = msg.Value
				if string(res) != "" {
					ch <- struct{}{}
				}
				return
			case err = <-partitionConsumer.Errors():
				fmt.Printf("err :%s\n", err.Error())
				return
			}
		}
	}
	return
}
