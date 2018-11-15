package controllers

import (
	"fmt"
	"strconv"
	"test/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type SearchController struct {
	beego.Controller
}

func (this *SearchController) Get() {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	var buildings []*models.Building
	qs := o.QueryTable("building")
	qs.Limit(100).All(&buildings)
	this.Data["Pages"] = buildings
	this.TplName = "search.html"
}

func (this *SearchController) Post() {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	const shortForm = "2018-11-11"
	var rooms []*models.Room
	building := this.GetString("selectBuilding")
	mybuilding, _ := strconv.Atoi(building)
	datetime := this.GetString("datetime")
	classtiming := this.GetString("classtiming")
	t, _ := time.Parse(shortForm, datetime)
	room := new(models.OrderRoom)

	o.QueryTable("room").Filter("Build", mybuilding).RelatedSel().All(&rooms)

	o.QueryTable("order_room").Filter("Build", mybuilding).Filter("Orderdate", t).One(room)
	fmt.Println(room)
	this.Data["Rooms"] = rooms
	fmt.Println(building, datetime, classtiming)
	this.TplName = "search.html"
}
