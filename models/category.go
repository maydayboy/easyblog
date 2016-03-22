package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Category struct {
	ID          int64     `orm:"auto"`
	Title       string    `orm:null`
	ViewCount   int64     `orm:"index;null"`
	TopicCount  int64     `orm:"index:null"`
	CreatedTime time.Time `orm:"index:null"`
}

func GetCategoryCount() (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("category")
	count, err := qs.Count()
	return count, err

}
func GetAllCategory() ([]*Category, error) {
	o := orm.NewOrm()
	category := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&category)
	return category, err
}

func AddCategory(title string) error {
	o := orm.NewOrm()
	cate := &Category{Title: title, CreatedTime: time.Now()}
	qs := o.QueryTable("category")
	err := qs.Filter("title", title).One(cate)

	if err == nil {
		return err
	}

	_, err = o.Insert(cate)
	if err != nil {
		return err
	}
	return nil

}

func DelteCatergory(id string) error {
	o := orm.NewOrm()
	idInt, _ := strconv.Atoi(id)
	cat := Category{ID: (int64)(idInt)}
	_, err := o.Delete(&cat)
	return err
}
