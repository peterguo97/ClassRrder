package controllers

import (
	"fmt"
	"test/models"

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
	fmt.Println(this.GetString("datatime"))
	fmt.Println(this.GetString("classtiming"))
	this.TplName = "search.html"
}
