module podcast_spider

go 1.15

require (
	github.com/Shopify/sarama v1.19.0
	github.com/astaxie/beego v1.12.3
	github.com/eapache/go-resiliency v1.2.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/olivere/elastic v6.2.35+incompatible
	github.com/olivere/elastic/v7 v7.0.22 // indirect
	github.com/pierrec/lz4 v2.6.0+incompatible // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
)

replace github.com/olivere/elastic v6.2.35+incompatible => ../github.com/olivere/elastic
