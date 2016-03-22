package controllers

import (
	"easyblog/models"
	"fmt"
	"strconv"
	//"github.com/astaxie/beego"
)

type AjaxController struct {
	BaseController
}

func (ac *AjaxController) Get() {
	do := ac.GetString("do")
	switch do {
	case "baseinfo":
		ac.GetBaseInfo()
		break
	case "getcategory":
		ac.GetCategory()
		break
	case "getcategorylen":
		ac.GetCategoryLen()
		break
	case "deletecategory":
		err := ac.DeleteCategory()
		if err != nil {
			fmt.Println(err)
		}

		break

	}

}

func (ac *AjaxController) Post() {
	if ac.CheckLogin() == false { //check the Is Login
		result := map[string]string{"status": "1"}
		ac.Ctx.Output.JSON(result, false, false)
		return
	}
	do := ac.GetString("do")
	title := ac.GetString("categorytitle")
	result := map[string]string{"status": "1"}
	switch do {
	case "createCategory":
		if title == "" {
			result["status"] = "2"
			ac.Ctx.Output.JSON(result, false, false)
			return
		}
		err := ac.CreateCategory(title)
		if err != nil {
			result["status"] = "2"
		} else {
			result["status"] = "0"
		}
		ac.Ctx.Output.JSON(result, false, false)

		break
	}

}

func (ac *AjaxController) GetBaseInfo() {
	var flag string = "0"
	uid := ac.Ctx.GetCookie("uid")
	if uid == "" {
		flag = "0"
		baseinfo := []string{flag, ""}
		ac.Ctx.Output.JSON(baseinfo, false, false)
		return
	}
	uidsession := ac.GetSession(uid)

	if uidsession == "" {
		flag = "0"
		baseinfo := []string{flag, ""}
		ac.Ctx.Output.JSON(baseinfo, false, false)
		return
	}
	flag = "1"
	value, ok := uidsession.(string)
	if !ok {
		fmt.Println("It's not ok for type string")
	}
	fmt.Println("flag is ", flag, "uidsession is ", uidsession)
	baseinfo := []string{flag, value}
	ac.Ctx.Output.JSON(baseinfo, false, false)
}

func (ac *AjaxController) CheckLogin() bool {
	uid := ac.Ctx.GetCookie("uid")
	if uid == "" {
		return false
	}
	uidsession := ac.GetSession(uid)
	if uidsession == "" {
		return false
	}
	return true

}

func (ac *AjaxController) GetCategory() {
	cate, err := models.GetAllCategory()
	if err != nil {
		ac.Ctx.Output.JSON(nil, false, false)
		return
	}
	current := ac.GetString("current")
	showCount := ac.GetString("showCount")
	iCurrent, err := strconv.Atoi(current)
	iShowCount, err := strconv.Atoi(showCount)
	startFlag := (iCurrent - 1) * iShowCount
	endFlag := len(cate)
	if len(cate) > iCurrent*iShowCount {
		endFlag = iCurrent * iShowCount
	}
	newCate := cate[startFlag:endFlag]
	ac.Ctx.Output.JSON(newCate, false, false)

}

func (ac *AjaxController) GetCategoryLen() {
	count, err := models.GetCategoryCount()
	if err != nil {
		ac.Ctx.Output.JSON(nil, false, false)
	}
	ac.Ctx.Output.JSON(count, false, false)
}

//Create new category
func (ac *AjaxController) CreateCategory(title string) error {

	err := models.AddCategory(title)
	return err
}

//Delete Ccategory
func (ac *AjaxController) DeleteCategory() error {
	ID := ac.GetString("ID")
	err := models.DelteCatergory(ID)
	ac.Redirect("/", 302)
	return err
}
