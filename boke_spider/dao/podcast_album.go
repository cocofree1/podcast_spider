package dao

import "time"

// 专辑
type Albums struct{
	Id               int		`orm:"column(id);auto;pk"`
	AnchorId         int        `description:"主播ID"`
	AnchorName       string     `description:"主播名"`
	Name             string     `description:"专辑名"`
	PictureUrl       string     `description:"封面图片"`
	ClassTag         string     `description:"分类(一级，二级)"`
	Introduction     string     `orm:"type(text)" description:"专辑简介"`
	Note             string     `orm:"type(text)" description:"专辑备注"`
	RssUrl           string     `description:"Rss链接"`
	CreatedAt        time.Time  `description:"创建日期"`
	UpdatedAt        time.Time  `description:"最后修改日期"`
}

type AlbumTag struct {
	Id               int		   `orm:"column(id);auto;pk"`
	FirstTagId       int           `description:"一级分类id"`
	SecondTagId      int           `description:"二级分类id"`
	AlbumId          int           `description:"专辑ID"`
	Description      string        `orm:"type(text)" description:"专辑简介"`
	CreatedAt        time.Time     `description:"创建日期"`
	UpdatedAt        time.Time     `description:"最后修改日期"`
}

