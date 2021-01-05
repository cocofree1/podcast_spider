package dao

import "time"

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
