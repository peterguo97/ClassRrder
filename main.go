package main

import (
	_ "test/routers"

	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/mysql"
)

// func init() {
// 	orm.RegisterDriver("mysql", orm.DRMySQL)

// 	orm.RegisterDataBase("default", "mysql", "root:gjj721385@tcp(127.0.0.1)/test?charset=utf8")

// 	orm.RunSyncdb("default", false, true)
// }

func main() {
	// fmt.Println(o.Insert(order))
	beego.Run()
}
