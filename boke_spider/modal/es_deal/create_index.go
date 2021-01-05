package es_deal

import (
	"boke_spider/lib"
	"boke_spider/modal/common"
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
