package lib

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 注册数据库
var DbObject orm.Ormer

func init(){
	conf, err := config.NewConfig("ini", "conf/db.conf")
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}
	// 获取配置信息
	hostname := conf.String("mysql::hostname")
	port, err := conf.Int("mysql::port")
	if err != nil{
		return
	}
	databaseName := conf.String("mysql::database_name")
	userName := conf.String("mysql::user_name")
	password := conf.String("mysql::password")
	charset := conf.String("mysql::charset")

	// 连接数据库
	confString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		userName,password,hostname,port, databaseName,charset)

	err = orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil{
		return
	}

	err = orm.RegisterDataBase("default", "mysql", confString)
	if err != nil{
		return
	}
	// 打开调试模式，开发的时候方便查看orm生成什么样子的sql语句
	orm.Debug = true
	// 自动创建表 参数二为是否开启创建表   参数三是否更新表
	_ = orm.RunSyncdb("default", false, true)
	DbObject = orm.NewOrm()
}