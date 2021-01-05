package es_deal

import (
	"boke_spider/dao"
	"boke_spider/lib"
	"boke_spider/modal/common"
	"context"
	"log"
	"strconv"
	"sync"
)

var OriginAlbumLock sync.Mutex
var OriginProgramLock sync.Mutex

func InsertOrUpdateEsPodcastList(list dao.PodcastTask){
	podcastList := common.GetOriginAlbumDataByPodcastList(list.PodcastList)
	data,_ := lib.EsClient.Get().Index(common.IndexNameMap[common.ORIGIN_ALBUM_DATA]).Id(strconv.Itoa(podcastList.CollectionId)).Do(context.Background())
	if data == nil{
		// 插入
		OriginAlbumLock.Lock()
		_,err := lib.EsClient.Index().
			Index(common.IndexNameMap[common.ORIGIN_ALBUM_DATA]).
			Id(strconv.Itoa(podcastList.CollectionId)).
			BodyJson(podcastList).
			Refresh("wait_for").
			Do(context.Background())
		if err != nil{
			log.Fatal(err)
		}
		OriginAlbumLock.Unlock()
	}else{
		OriginAlbumLock.Lock()
		_,err := lib.EsClient.Update().
			Index(common.IndexNameMap[common.ORIGIN_ALBUM_DATA]).
			Id(strconv.Itoa(podcastList.CollectionId)).
			Doc(podcastList).
			Do(context.Background())
		if err != nil{
			log.Fatal(err)
		}
		OriginAlbumLock.Unlock()
	}
}

func InsertOrUpdateEsPodcastContent(content dao.XmlItemCont){
	podcastContent := common.GetOriginProgramDataByXmlItem(content.XmlItem)
	data,_ := lib.EsClient.Get().Index(common.IndexNameMap[common.ORIGIN_PROGRAM_DATA]).Id(podcastContent.Guid).Do(context.Background())
	if data == nil{
		// 插入
		OriginProgramLock.Lock()
		_,err := lib.EsClient.Index().
			Index(common.IndexNameMap[common.ORIGIN_PROGRAM_DATA]).
			Id(podcastContent.Guid).
			BodyJson(podcastContent).
			Refresh("wait_for").
			Do(context.Background())
		if err != nil{
			log.Fatal(err)
		}
		OriginProgramLock.Unlock()
	}else{
		OriginProgramLock.Lock()
		_,err := lib.EsClient.Update().
			Index(common.IndexNameMap[common.ORIGIN_PROGRAM_DATA]).
			Id(podcastContent.Guid).
			Doc(podcastContent).
			Do(context.Background())
		if err != nil{
			log.Fatal(err)
		}
		OriginProgramLock.Unlock()
	}
}
