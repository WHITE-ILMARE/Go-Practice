package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

func init() {
	// 注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/go_practice?charset=utf8", 30)
	// 注册定义的model
	orm.RegisterModel(new(User))
	// 创建table
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()
	user := User{Name: "slene"}
	// 插入表
	id, err := o.Insert(&user)
	fmt.Println("ID: %d, ERR: %v\n", id, err)

	// 更新表
	user.Name = "astaxie"
	num, err := o.Update(&user)
	fmt.Println("NUM: %d, ERR: %v\n", num, err)

	// 读取one
	u := User{Id: user.Id}
	err = o.Read(&u)
	fmt.Println("ERR: %v\n", err)

	// 删除表
	num, err = o.Delete(&u)
	fmt.Println("NUM: %d, ERR: %v\n", num, err)
}
