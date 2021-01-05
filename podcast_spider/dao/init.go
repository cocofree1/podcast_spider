package dao

import (
	"github.com/astaxie/beego/orm"
)

func init(){
	orm.RegisterModel(
		new(OriginAlbumData),
		new(OriginProgramData),
		new(Albums),
		new(Programs),
		new(AlbumTag),
	)
}
