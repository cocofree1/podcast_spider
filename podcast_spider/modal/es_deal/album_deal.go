package es_deal

import (
	"podcast_spider/dao"
	"podcast_spider/lib"
	"podcast_spider/modal/common"
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

var AlbumLock sync.Mutex

func InsertOrUpdateEsAlbum(list dao.PodcastTask){
	now := time.Now()
	classTag := strconv.Itoa(list.AlbumTag.FirstTagId)
	if list.AlbumTag.SecondTagId > 0{
		classTag = fmt.Sprintf("%s,%d",classTag,list.AlbumTag.SecondTagId)
	}
	albums := &dao.Albums{
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
	data,_ := lib.EsClient.Get().Index(common.IndexNameMap[common.ALBUM_DATA]).Id(strconv.Itoa(albums.Id)).Do(context.Background())
	if data == nil{
		// 插入
		AlbumLock.Lock()
		_,err := lib.EsClient.Index().
			Index(common.IndexNameMap[common.ALBUM_DATA]).
			Id(strconv.Itoa(albums.Id)).
			BodyJson(albums).
			Refresh("wait_for").
			Do(context.Background())
		if err != nil{
			log.Fatal(err)
		}
		AlbumLock.Unlock()
	}else{
		AlbumLock.Lock()
		_,err := lib.EsClient.Update().
			Index(common.IndexNameMap[common.ALBUM_DATA]).
			Id(strconv.Itoa(albums.Id)).
			Doc(albums).
			Do(context.Background())
		if err != nil{
			log.Fatal(err)
		}
		AlbumLock.Unlock()
	}
}

