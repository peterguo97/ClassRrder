package controllers

import (
	"encoding/json"
	"fmt"
	"test/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ConfirmController struct {
	beego.Controller
}

func (c *ConfirmController) Post() {
	orm.Debug = true
	type Order struct {
		Build    int    `json:"build"`
		Room     int    `json:"room"`
		Datetime string `json:"datetime"`
		Timing   uint8  `json:"timing"`
	}

	type msg struct {
		Success string `json:"success"`
		Err     string `json:"err"`
	}

	o := orm.NewOrm()
	o.Using("default")
	var or Order
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &or)
	fmt.Println(or.Timing)
	if err != nil {
		fmt.Println(err)
	}
	t, err := time.Parse("2006-01-02", or.Datetime)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(t)
	}
	res := new(models.OrderRoom)
	qs := o.QueryTable("OrderRoom").Filter("Build", or.Build).Filter("Room", or.Room).Filter("OrderDate", t).Filter("ClassTiming", or.Timing)
	fmt.Println(res)
	qs.One(res)
	messsge := new(msg)
	if !qs.Exist() || !res.HasOrdered {
		mydate := &models.OrderRoom{
			Build:       &models.Building{Id: or.Build},
			Room:        &models.Room{Id: or.Room},
			OrderDate:   t,
			ClassTiming: or.Timing,
			HasOrdered:  true,
		}
		if !qs.Exist() {
			_, err := o.Insert(mydate)
			if err != nil {
				messsge.Err = "插入出错"
			}
		} else {
			res.HasOrdered = true
			_, err := o.Update(res)
			if err != nil {
				messsge.Err = "update error"
			}
		}
		messsge.Success = "1"

	} else {
		messsge.Err = "exist"
	}
	c.Data["json"] = messsge
	c.ServeJSON()
}
