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
	Myerr := o.QueryTable("OrderRoom").Filter("Build", or.Build).Filter("Room", or.Room).Filter("OrderDate", t).Filter("ClassTiming", or.Timing).One(res)
	fmt.Println(Myerr)
	message := new(msg)
	if Myerr == orm.ErrNoRows || !res.HasOrdered {
		mydate := &models.OrderRoom{
			Build:       &models.Building{Id: or.Build},
			Room:        &models.Room{Id: or.Room},
			OrderDate:   t,
			ClassTiming: or.Timing,
			HasOrdered:  true,
		}
		if Myerr == orm.ErrNoRows {
			_, err := o.Insert(mydate)
			if err != nil {
				message.Err = "insert failed"
			} else {
				message.Success = "insert success"
			}
		} else {
			res.HasOrdered = true
			_, err := o.Update(res)
			if err != nil {
				message.Err = "update error"
			} else {
				message.Success = "update success"
			}
		}

	} else {
		message.Err = "has existed"
	}
	c.Data["json"] = message
	c.ServeJSON()
}
