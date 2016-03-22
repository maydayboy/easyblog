package controllers

type LogoutController struct {
	BaseController
}

func (lc *LogoutController) Get() {
	lc.Logout()
}

func (lc *LogoutController) Post() {
	lc.Logout()
}

func (lc *LogoutController) Logout() {
	uid := lc.Ctx.GetCookie("uid")
	lc.DelSession(uid)
	lc.Ctx.SetCookie("uid", "")
	lc.Redirect("/login", 302)

}
