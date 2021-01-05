package kafka_deal

import (
	"podcast_spider/dao"
	"podcast_spider/lib"
	"podcast_spider/modal/common"
	"podcast_spider/modal/es_deal"
	"podcast_spider/modal/mysql_deal"
	"encoding/json"
)

// 专辑原始数据
func DealKafkaToOriginAlbum() {
	lib.LoopConsumer(common.IndexNameMap[common.ORIGIN_ALBUM_DATA], OriginAlbumDeal)
}

func OriginAlbumDeal(data []byte) {
	var podcastTask dao.PodcastTask
	err := json.Unmarshal(data,&podcastTask)
	if err != nil{
		return
	}

	// 写入mysql,es
	go mysql_deal.InsertOrUpdateValueToPodcastList(podcastTask)
	go es_deal.InsertOrUpdateEsPodcastList(podcastTask)
}

// 内容原始数据
func DealKafkaToOriginProgram() {
	lib.LoopConsumer(common.IndexNameMap[common.ORIGIN_PROGRAM_DATA], OriginProgramDeal)
}

func OriginProgramDeal(data []byte) {
	var content dao.XmlItemCont
	err := json.Unmarshal(data,&content)
	if err != nil{
		return
	}
	//写入mysql,es
	go mysql_deal.InsertOrUpdateValueToPodcastContent(content)
	go es_deal.InsertOrUpdateEsPodcastContent(content)
}

// 专辑
func DealKafkaToAlbum() {
	lib.LoopConsumer(common.IndexNameMap[common.ALBUM_DATA], AlbumDeal)
}

func AlbumDeal(data []byte) {
	var podcastTask dao.PodcastTask
	err := json.Unmarshal(data,&podcastTask)
	if err != nil{
		return
	}
	//写入mysql,es
	go mysql_deal.InsertOrUpdateAlbum(podcastTask)
	go es_deal.InsertOrUpdateEsAlbum(podcastTask)
}


// 节目
func DealKafkaToProgram() {
	lib.LoopConsumer(common.IndexNameMap[common.PROGRAM_DATA], ProgramDeal)
}

func ProgramDeal(data []byte) {
	var content dao.XmlItemCont
	err := json.Unmarshal(data,&content)
	if err != nil{
		return
	}
	//写入mysql,es
	go mysql_deal.InsertOrUpdateProgram(content)
	go es_deal.InsertOrUpdateProgram(content)
}