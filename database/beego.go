package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"go/types"
)

type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

func init() {
	types.Array{}
	// 注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/go_practice?charset=utf8", 30)
	// 注册定义的model
	orm.RegisterModel(new(User))

	// 创建table
	orm.RunSyncdb("default", false, true)
}
