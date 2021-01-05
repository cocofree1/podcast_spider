package es_deal

import (
	"podcast_spider/lib"
	"podcast_spider/modal/common"
	"context"
	"log"
)

func init(){
	for _,value := range common.IndexNameMap{
		flag,_ := lib.EsClient.IndexExists(value).Do(context.Background())
		if !flag{
			_,err := lib.EsClient.CreateIndex(value).Do(context.Background())
			if err != nil{
				log.Fatal(err)
			}
		}
	}
}
