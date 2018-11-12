package controllers

import (
	"fmt"
	"test/models"
	"test/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//LoginController is a controller to handle the login
type LoginController struct {
	beego.Controller
}

//Get is to hadle the method Get
func (l *LoginController) Get() {
	l.TplName = "login.html"
}

//Post is to handle method Post
func (l *LoginController) Post() {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable("user")
	if !(qs.Filter("name", l.GetString("username")).Exist()) {
		fmt.Println("查无此人！")
	} else {
		user := new(models.User)
		qs.Filter("name", l.GetString("username")).One(user)
		if user.Pass == utils.CryptoPass(l.GetString("userpass")) {
			v := l.GetSession("hello")
			if v == nil {
				l.SetSession("test", int(11))
			}
			fmt.Printf("hello %s, welcome home\n", user.Name)
		} else {
			fmt.Println("get away, faker")
		}
		// pass := l.GetString("userpass")
	}
	l.TplName = "login.html"
}
