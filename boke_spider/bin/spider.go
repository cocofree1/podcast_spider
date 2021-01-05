package main

import (
	_ "boke_spider/dao"
	_ "boke_spider/lib"
	"boke_spider/modal"
)

func main() {
	modal.RunKafkaSpider()
}

