package dao

import "github.com/astaxie/beego/orm"

func init(){
	orm.RegisterModel(
		new(Albums),
		new(Programs),
		new(HomeConfig),
		new(Anchors),
		new(PodcastTag),
		new(AlbumTag),
	)
}