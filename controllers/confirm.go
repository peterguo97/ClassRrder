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
	if err != nil {
		fmt.Println(err)
	}
	t, err := time.Parse("2006-01-02", or.Datetime)
	if err != nil {
		fmt.Println(err)
	}
	res := new(models.OrderRoom)
	myerr := o.QueryTable("OrderRoom").Filter("Build", or.Build).Filter("Room", or.Room).Filter("OrderDate", t).Filter("ClassTiming", or.Timing).One(res)
	message := new(msg)
	if myerr == orm.ErrNoRows || !res.Waiting {
		mydate := &models.OrderRoom{
			Build:       &models.Building{Id: or.Build},
			Room:        &models.Room{Id: or.Room},
			OrderDate:   t,
			ClassTiming: or.Timing,
			Waiting:     true,
		}
		if myerr == orm.ErrNoRows {
			_, err := o.Insert(mydate)
			if err != nil {
				message.Err = "插入出错"
			} else {
				message.Success = "insert success"
			}
		} else {
			res.Waiting = true
			_, err := o.Update(res)
			if err != nil {
				message.Err = "update error"
			} else {
				message.Success = "update success"
			}
		}
	} else {
		message.Err = "exist"
	}
	c.Data["json"] = message
	c.ServeJSON()
}
