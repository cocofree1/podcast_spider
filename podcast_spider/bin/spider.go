package main

import (
	_ "podcast_spider/dao"
	_ "podcast_spider/lib"
	"podcast_spider/modal"
)

func main() {
	modal.RunKafkaSpider()
}

