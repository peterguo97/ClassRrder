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
	r.TplName = "zhuce.html"
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
		r.Ctx.WriteString("<h1>注册成功</h1>")
	} else {
		r.Ctx.WriteString("<h1>The user is existed!</h1>")
		fmt.Println("The user is existed")
	}
}
