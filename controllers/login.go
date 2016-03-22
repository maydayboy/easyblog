package controllers

import (
	"easyblog/models"

	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
	//"github.com/astaxie/beego"
)

type LoginController struct {
	BaseController
}

func (lc *LoginController) Get() {
	if lc.CheckLogin() {
		lc.Redirect("/home", 302)
	}
	lc.TplName = "login.html"
}

func (lc *LoginController) Post() {
	userName := lc.GetString("name")
	password := lc.GetString("pswd")
	Users, err := models.GetAllUser()
	if err != nil {
		lc.Redirect("/login", 302)
		return
	}
	for _, v := range Users {
		if v.UserName == userName && v.Password == password {
			uid := makeHash(userName, "")
			lc.SetSession(uid, userName)
			lc.Ctx.SetCookie("uid", uid)
			lc.Redirect("/home", 302)
		}
	}
	lc.TplName = "login.html"
}

func (lc *LoginController) CheckLogin() bool {
	uid := lc.Ctx.GetCookie("uid")
	if uid == "" {
		return false
	}
	uidsession := lc.GetSession(uid)
	if uidsession == "" {
		return false
	}
	return true

}

func makeHash(userID string, securt string) (cipherStr string) {
	t := fmt.Sprintf("%d", time.Now().Unix())
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(t + userID))
	if securt == "" {
		cipherStr = hex.EncodeToString(md5Ctx.Sum(nil))
	} else {
		cipherStr = hex.EncodeToString(md5Ctx.Sum([]byte(securt)))
	}
	return
}
