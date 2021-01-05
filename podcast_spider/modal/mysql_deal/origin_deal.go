package mysql_deal

import (
	"podcast_spider/dao"
	"podcast_spider/lib"
	"podcast_spider/modal/common"
	"log"
	"sync"
)

var (
	OriginAlbumLock    sync.Mutex
	OriginProgramLock  sync.Mutex
)

//插入单个数据
func InsertOrUpdateValueToPodcastList(list dao.PodcastTask){
	podcastList := common.GetOriginAlbumDataByPodcastList(list.PodcastList)
	OriginAlbumLock.Lock()
	searchPodcastList := podcastList
	if created, id, err := lib.DbObject.ReadOrCreate(&searchPodcastList, "CollectionId"); err == nil {
		if !created {
			podcastList.Id = int(id)
			_,err = lib.DbObject.Update(&podcastList)
			if err != nil{
				log.Fatal(err)
			}
		}
	} else {
		log.Fatal(err)
	}
	OriginAlbumLock.Unlock()
}

func InsertOrUpdateValueToPodcastContent(content dao.XmlItemCont){
	podcastContent := common.GetOriginProgramDataByXmlItem(content.XmlItem)
	OriginProgramLock.Lock()
	searchPodcastContent := podcastContent
	if created, id, err := lib.DbObject.ReadOrCreate(&searchPodcastContent, "Guid"); err == nil {
		if !created {
			podcastContent.Id = int(id)
			_,err = lib.DbObject.Update(&podcastContent)
			if err != nil{
				log.Fatal(err)
			}
		}
	} else {
		log.Fatal(err)
	}
	OriginProgramLock.Unlock()
}

