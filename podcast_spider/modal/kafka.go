package modal

import (
	"podcast_spider/modal/kafka_deal"
)

func RunKafkaSpider(){
	exit := make(chan int)
	// 任务处理
	go kafka_deal.GetAlbumTag()
	go kafka_deal.TimerSpiderTask()
	go kafka_deal.DealTaskContents()
	// 处理数据
	go kafka_deal.DealKafkaToOriginAlbum()
	go kafka_deal.DealKafkaToOriginProgram()
	go kafka_deal.DealKafkaToAlbum()
	go kafka_deal.DealKafkaToProgram()
	<-exit
}
