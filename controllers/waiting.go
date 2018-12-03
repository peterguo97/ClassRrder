package controllers

import (
	"fmt"
	"test/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type WaitingController struct {
	beego.Controller
}

func (w *WaitingController) Get() {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	var waiting []*models.OrderRoom
	o.QueryTable("OrderRoom").Filter("Waiting", true).All(&waiting, "Id", "Room", "Build", "ClassTiming", "OrderDate")
	for _, item := range waiting {
		fmt.Println(item)
	}
	w.Ctx.WriteString("hello world")
}
