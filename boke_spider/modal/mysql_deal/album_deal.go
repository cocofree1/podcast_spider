package mysql_deal

import (
	"boke_spider/dao"
	"boke_spider/lib"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

var AlbumLock sync.Mutex

func InsertOrUpdateAlbum(list dao.PodcastTask){
	now := time.Now()
	classTag := strconv.Itoa(list.AlbumTag.FirstTagId)
	if list.AlbumTag.SecondTagId > 0{
		classTag = fmt.Sprintf("%s,%d",classTag,list.AlbumTag.SecondTagId)
	}
	albums := dao.Albums{
		Id:               list.PodcastList.CollectionId,
		AnchorId:         list.PodcastList.ArtistId,
		AnchorName:       list.PodcastList.ArtistName,
		Name:             list.PodcastList.CollectionName,
		PictureUrl:       list.PodcastList.ArtworkUrl60,
		ClassTag:         classTag,
		Introduction:     list.AlbumTag.Description,
		RssUrl:           list.PodcastList.FeedUrl,
		CreatedAt:        list.PodcastList.ReleaseDate,
		UpdatedAt:        now,
	}
	AlbumLock.Lock()
	searchAlbums := albums
	if created, _, err := lib.DbObject.ReadOrCreate(&searchAlbums, "Id"); err == nil {
		if !created {
			_,err = lib.DbObject.Update(&albums)
			if err != nil{
				log.Fatal(err)
			}
		}
	} else {
		log.Fatal(err)
	}
	AlbumLock.Unlock()
}


func GetAllPodcastTag()(tags []dao.AlbumTag){
	_,err := lib.DbObject.Raw("select * from album_tag").QueryRows(&tags)
	if err != nil{
		log.Fatal("get album_tag fail",err)
	}
	return
}
