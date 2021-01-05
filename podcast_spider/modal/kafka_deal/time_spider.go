package kafka_deal

import (
	"podcast_spider/dao"
	"podcast_spider/lib"
	"podcast_spider/modal/common"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func TimerSpiderTask(){
	t := time.NewTicker(time.Minute)
	for{
		if _,ok := <- t.C;ok{
			for i := 0; i < 20;i++{
				go SpiderTask(<-PodcastTagChannel)
			}
		}
	}
}

// 爬取任务
func SpiderTask(albumTag dao.AlbumTag){
	url := fmt.Sprintf("https://itunes.apple.com/search?country=cn&media=podcast&entity=podcast&term=%d",albumTag.AlbumId)
	// 获取数据
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("http get error", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("read error", err)
		return
	}
	// 解析
	var podcastLists []dao.PodcastList
	var podcastResult dao.PodcastResult
	err = json.Unmarshal(body, &podcastResult)
	if err != nil{
		log.Fatal(err)
		return
	}
	podcastLists = podcastResult.Results
	if len(podcastLists) > 0{
		podcastTask := dao.PodcastTask{
			PodcastList: podcastLists[0],
			AlbumTag: albumTag,
		}
		// 写入kafka
		task,err := json.Marshal(podcastTask)
		if err != nil{
			log.Fatal(err)
			return
		}
		TaskChannel <- podcastTask
		lib.Send(common.IndexNameMap[common.ORIGIN_ALBUM_DATA],string(task))
		lib.Send(common.IndexNameMap[common.ALBUM_DATA],string(task))
	}
	return
}

// 处理数据
func DealData(task dao.PodcastTask)(result []dao.XmlItem){
	resp, err := http.Get(task.PodcastList.FeedUrl)
	if err != nil {
		log.Fatal("http get error", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("read error", err)
		return
	}
	result = UnmarshalPodcast(body)


	for _,item := range result{
		var xmlItemCont dao.XmlItemCont
		xmlItemCont.XmlItem = item
		xmlItemCont.CollectionId = task.PodcastList.CollectionId
		// 写入kafka
		program, err := json.Marshal(xmlItemCont)
		if err != nil{
			log.Fatal(err)
			continue
		}
		lib.Send(common.IndexNameMap[common.ORIGIN_PROGRAM_DATA],string(program))
		lib.Send(common.IndexNameMap[common.PROGRAM_DATA],string(program))
	}
	return
}

// 解析数据
func UnmarshalPodcast(body []byte)(result []dao.XmlItem){
	xmlRss := &dao.XmlRss{}
	err := xml.Unmarshal(body, &xmlRss)
	if err != nil{
		log.Fatal(err)
		return
	}
	if len(xmlRss.Channel) > 0{
		result = xmlRss.Channel[0].Item
	}
	return
}


