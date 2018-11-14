package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"test/models"
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
	var rooms []*models.Room
	building := this.GetString("selectBuilding")
	mybuilding, _ := strconv.Atoi(building)
	datetime := this.GetString("datetime")
	classtiming := this.GetString("classtiming")
	o.QueryTable("room").Filter("Build", mybuilding).RelatedSel().All(&rooms)
	for _, v := range rooms {
		fmt.Println(v.Id)
	}
	fmt.Println(building, datetime, classtiming)
	this.TplName = "result.html"
}
