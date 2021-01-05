package lib

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/config"
	"log"
	"strings"
)

var Producer sarama.AsyncProducer

func init(){
	conf, err := config.NewConfig("ini", "conf/db.conf")
	if err != nil {
		log.Fatal("new config failed, err:", err)
		return
	}
	// 获取配置信息
	addrs := conf.String("kafka::addrs")

	// 获取配置文件
	kafkaConfig := sarama.NewConfig()
	client, err := sarama.NewClient(strings.Split(addrs, ","), kafkaConfig)
	if err != nil {
		log.Fatal("unable to create kafka client: ", err)
	}
	Producer, err = sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		log.Fatal(err)
	}
}

func Send(topic, data string) {
	Producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(data)}
}