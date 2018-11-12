package controllers

import (
	"fmt"
	"test/models"
	"test/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Get() {
	r.TplName = "register.min.html"
}

func (r *RegisterController) Post() {
	orm.Debug = true
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	username := r.GetString("username")
	if !(qs.Filter("name", username).Exist()) {
		var pass string = utils.CryptoPass(r.GetString("userpass"))
		fmt.Printf("the len of the pass %d\n", len(pass))
		user := &models.User{Name: username, Pass: pass}
		o.Insert(user)
	} else {
		fmt.Println("The user is existed")
	}
	r.TplName = "register.min.html"
}
