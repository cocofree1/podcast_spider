package lib

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/config"
	"log"
	"strings"
)

var Consumer sarama.Consumer
type ConsumerCallback func(data []byte)

func init(){
	conf, err := config.NewConfig("ini", "conf/db.conf")
	if err != nil {
		log.Fatal("new config failed, err:", err)
		return
	}
	// 获取配置信息
	addrs := conf.String("kafka::addrs")

	kafkaConfig := sarama.NewConfig()
	client, err := sarama.NewClient(strings.Split(addrs, ","), kafkaConfig)
	if err != nil {
		log.Fatal("unable to create kafka client: ", err)
	}

	Consumer, err = sarama.NewConsumerFromClient(client)
	if err != nil {
		log.Fatal(err)
	}
}


func LoopConsumer(topic string, callback ConsumerCallback) {
	partitionConsumer, err := Consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer partitionConsumer.Close()

	for {
		msg := <-partitionConsumer.Messages()
		if callback != nil {
			callback(msg.Value)
		}
	}
}