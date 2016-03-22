package main

import (
	"easyblog/controllers"

	_ "easyblog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/home", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.RESTRouter("/ajax", &controllers.AjaxController{})
	beego.Run()
}
