package mysql_deal

import (
	"boke_spider/dao"
	"boke_spider/lib"
	"boke_spider/modal/common"
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
	programs := dao.Programs{
		AlbumId:       content.CollectionId,
		Name:          content.XmlItem.Title,
		Introduction:  description,
		AudioUrl:      content.XmlItem.Guid,
		CreatedAt:     pubDate,
		UpdatedAt:     now,
	}
	ProgramLock.Lock()
	searchProgram := programs
	if created, id, err := lib.DbObject.ReadOrCreate(&searchProgram, "AudioUrl"); err == nil {
		if !created {
			programs.Id = int(id)
			_,err = lib.DbObject.Update(&programs)
			if err != nil{
				log.Fatal(err)
			}
		}
	} else {
		log.Fatal(err)
	}
	ProgramLock.Unlock()
}