package models

import (
	"time"

	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
)

type Building struct {
	Id         int
	Name       string       `orm:"size(20);unique"`
	Room       []*Room      `orm:"reverse(many)"`
	OrderRooms []*OrderRoom `orm:"reverse(many)"`
}

type Room struct {
	Id         int
	Name       string       `orm:"size(20);unique"`
	Build      *Building    `orm:"rel(fk)"`
	OrderRooms []*OrderRoom `orm:"reverse(many)"`
}

type User struct {
	Id       int
	Name     string `orm:"size(20)"`
	Pass     string
	Register time.Time `orm:"auto_now_add;type(datetime)"`
}

type OrderRoom struct {
	Id          int
	Build       *Building `orm:"rel(fk)"`
	Room        *Room     `orm:"rel(fk)"`
	HasOrdered  bool      `orm:"default(true)"`
	ClassTiming uint8     `orm:"defalut(0)"`
	OrderDate   time.Time `orm:"type(date)"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:gjj721385@tcp(127.0.0.1)/test?charset=utf8")
	orm.RegisterModel(new(Building), new(Room), new(User), new(OrderRoom))
	orm.RunSyncdb("default", false, true)
}
