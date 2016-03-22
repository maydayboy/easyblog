package models

import (
	"github.com/astaxie/beego/orm"
	//"strconv"
	//"time"
)

type User struct {
	Id       int64  `orm:"auto"`
	UserName string `orm:"not null"`
	Password string `orm:"not null"`
	Others   string `orm:"not null"`
}

func GetAllUser() ([]*User, error) {
	o := orm.NewOrm()
	users := make([]*User, 0)
	qs := o.QueryTable("user")
	_, err := qs.All(&users)
	return users, err
}
