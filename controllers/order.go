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
	qs := o.QueryTable("OrderRoom")
	type order struct {
		Build    int
		Room     int
		Datetime string
	}

	type classes struct {
		Id []int `json:"id"`
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
	orderqs := qs.Filter("Build", or.Build).Filter("Room", or.Room).Filter("OrderDate", t)
	timing := [6]int{0, 1, 2, 3, 4, 5}
	class := new(classes)
	if orderqs.Exist() {
		var myclass []*models.OrderRoom
		fmt.Println("found it")
		orderqs.All(&myclass, "ClassTiming")
		type st struct {
			start int
			end   int
		}
		myst := &st{
			start: 1,
			end:   0,
		}
		var res = make([]int, 0)
		for _, item := range myclass {
			myst.end = int(item.ClassTiming)
			fmt.Println(myst.start, myst.end)
			res = append(res, timing[myst.start:myst.end]...)
			myst.start = myst.end + 1
		}
		res = append(res, timing[myst.start:]...)
		fmt.Println(res)
		class.Id = res
	} else {
		room := &models.OrderRoom{
			Build:     &models.Building{Id: or.Build},
			Room:      &models.Room{Id: or.Room},
			OrderDate: t,
		}
		class.Id = timing[:]
		o.Insert(room)
	}
	this.Data["json"] = class
	this.ServeJSON()
}
