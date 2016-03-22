package controllers

type HomeController struct {
	BaseController
}

func (hb *HomeController) Get() {
	hb.TplName = "home.html"
}
