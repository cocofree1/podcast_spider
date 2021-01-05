package lib

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/olivere/elastic"
	"log"
)

var EsClient  *elastic.Client

func init(){
	conf, err := config.NewConfig("ini", "conf/db.conf")
	if err != nil {
		log.Fatal("new config failed, err:", err)
		return
	}
	// 获取配置信息
	hostname := conf.String("es::hostname")
	port, err := conf.Int("es::port")
	if err != nil{
		log.Fatal(err)
	}
	url := fmt.Sprintf("http://%s:%d",hostname,port)
	c, err := elastic.NewClient(elastic.SetSniff(false),elastic.SetURL(url))
	if err != nil{
		log.Fatal(err)
	}
	EsClient = new(elastic.Client)
	EsClient = c
}