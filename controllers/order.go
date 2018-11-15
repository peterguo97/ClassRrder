package controllers

import (
	"encoding/json"
	"fmt"
	"test/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type OrderController struct {
	beego.Controller
}

func (this *OrderController) Post() {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	type order struct {
		Build    int
		Room     int
		Datetime string
	}

	var or order
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &or)
	if err != nil {
		fmt.Println(err)
	}
	t, err := time.Parse("2006-01-02", or.Datetime)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(t)
	}
	var room models.OrderRoom
	exist := o.QueryTable("OrderRoom").Filter("Build", or.Build).Filter("Room", or.Room).Filter("OrderDate", t).Exist()
	if exist {
		fmt.Println("found it")
	} else {
		room := &models.OrderRoom{
			Build:     &models.Building{Id: or.Build},
			Room:      &models.Room{Id: or.Room},
			OrderDate: t,
		}
		o.Insert(room)
		fmt.Println("cannot found it")
	}
	this.Data["json"] = &room
	this.ServeJSON()
}
