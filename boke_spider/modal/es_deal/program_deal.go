package es_deal

import (
	"boke_spider/dao"
	"boke_spider/lib"
	"boke_spider/modal/common"
	"context"
	"log"
	"sync"
	"time"
)

var ProgramLock sync.Mutex

func InsertOrUpdateProgram(content dao.XmlItemCont){
	now := time.Now()
	pubDate := now
	if content.XmlItem.PubDate != ""{
		pubDate = common.DateFormToTime(content.XmlItem.PubDate)
	}

	description := common.DescriptionRegexp(content.XmlItem.Description)
	programs := &dao.Programs{
		AlbumId:      content.CollectionId,
		Name:         content.XmlItem.Title,
		Introduction: description,
		AudioUrl:     content.XmlItem.Guid,
		CreatedAt:    pubDate,
		UpdatedAt:    now,
	}

	data,_ := lib.EsClient.Get().Index(common.IndexNameMap[common.PROGRAM_DATA]).Id(programs.AudioUrl).Do(context.Background())
	if data == nil{
		// 插入
		ProgramLock.Lock()
		_,err := lib.EsClient.Index().
			Index(common.IndexNameMap[common.PROGRAM_DATA]).
			Id(programs.AudioUrl).
			BodyJson(programs).
			Refresh("wait_for").
			Do(context.Background())
		if err != nil{
			log.Fatal(err)
		}
		ProgramLock.Unlock()
	}else{
		ProgramLock.Lock()
		_,err := lib.EsClient.Update().
			Index(common.IndexNameMap[common.PROGRAM_DATA]).
			Id(programs.AudioUrl).
			Doc(programs).
			Do(context.Background())
		if err != nil{
			log.Fatal(err)
		}
		ProgramLock.Unlock()
	}
}
