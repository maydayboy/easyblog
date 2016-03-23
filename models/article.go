package models

import (
	"github.com/astaxie/beego/orm"
	//"strconv"
	"time"
)

type Article struct {
	ID          int64     `orm:"auto"`
	Title       string    `orm:"null"`
	Content     string    `orm:"null"`
	CategoryStr string    `orm:"null"`
	CreateTime  time.Time `orm:"null"`
}

func GetArticleCount() (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("article")
	count, err := qs.Count()
	return count, err
}

func GetArticle() ([]*Article, error) {
	o := orm.NewOrm()
	article := make([]*Article, 0)
	qs := o.QueryTable("article")
	_, err := qs.All(&article)
	return article, err
}
