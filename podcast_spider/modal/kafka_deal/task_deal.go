package kafka_deal

import (
	"podcast_spider/dao"
	"podcast_spider/modal/mysql_deal"
)

var TaskChannel chan dao.PodcastTask
var PodcastTagChannel chan dao.AlbumTag

func init(){
	TaskChannel = make(chan dao.PodcastTask, 100)
	PodcastTagChannel = make(chan dao.AlbumTag, 5000)
}

// 处理任务
func DealTaskContents(){
	for{
		if task,ok := <-TaskChannel; ok{
			go DealData(task)
		}
	}
}

func GetAlbumTag(){
	for {
		tags := mysql_deal.GetAllPodcastTag()
		for _,item := range tags{
			PodcastTagChannel <- item
		}
	}
}