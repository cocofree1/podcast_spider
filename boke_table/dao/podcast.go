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

//节目
type Programs struct{
	Id            int		 `orm:"column(id);auto;pk"`
	AlbumId       int        `description:"专辑ID"`
	Name          string     `description:"节目名"`
	Introduction  string     `orm:"type(text)" description:"节目简介"`
	AudioUrl      string     `description:"音频链接"`
	CreatedAt     time.Time  `description:"创建日期"`
	UpdatedAt     time.Time  `description:"最后修改日期"`
}

// 页面配置
type HomeConfig struct{
	Id              int			 `orm:"column(id);auto;pk"`
	EffectStatus    int          `description:"生效状态"`
	LoopType        int          `description:"循环状态"`
	Words           string       `description:"配置文案"`
	SkipType        string       `description:"跳转地址"`
	SkipUrl         string       `description:"跳转地址"`
	StartTime       time.Time    `description:"生效开始时间"`
	EndTime         time.Time    `description:"生效结束时间"`
	OffOn           bool         `description:"推荐是否打开"`
	CreatedAt       time.Time    `description:"创建日期"`
	UpdatedAt       time.Time    `description:"最后修改日期"`
}

// 主播/用户表
type Anchors struct{
	Id            int		 `orm:"column(id);auto;pk"`
	Name          string     `description:"主播名"`
	CertType      int        `description:"认证类型"`
	HeadPicture   string     `description:"头像"`
	Signature     string     `description:"签名"`
	Status        string     `description:"账户状态"`
	CreatedAt     time.Time  `description:"创建日期"`
	UpdatedAt     time.Time  `description:"最后修改日期"`
}

// 专辑对应标签
type AlbumTag struct {
	Id               int		   `orm:"column(id);auto;pk"`
	FirstTagId       int           `description:"一级分类id"`
	SecondTagId      int           `description:"二级分类id"`
	AlbumId          int           `description:"专辑ID"`
	Description      string        `orm:"type(text)" description:"专辑简介"`
	CreatedAt        time.Time     `description:"创建日期"`
	UpdatedAt        time.Time     `description:"最后修改日期"`
}

// 播客标签
type PodcastTag struct {
	Id               int		   `description:"播客中的标签ID"`
	TagName          string        `description:"播客中的标签名"`
	CreatedAt        time.Time     `description:"创建日期"`
	UpdatedAt        time.Time     `description:"最后修改日期"`
}