package controllers

import (
	// "fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplNames = "login.html"
}

func (this *LoginController) Post() {
	//this.Ctx.WriteString(fmt.Sprint(this.Input()))
	account := this.Input().Get("account")
	pwd := this.Input().Get("pwd")
	autologin := (this.Input().Get("autologin") == "on")

	//this.Ctx.WriteString(account)
	//this.Ctx.WriteString(pwd)
	// this.Ctx.WriteString(autologin)
	beego.Debug("login", autologin)
	maxAge := 0
	if autologin {
		maxAge = 1<<31 - 1
	}
	this.Ctx.SetCookie("account", account, maxAge, "/")
	this.Ctx.SetCookie("pwd", pwd, maxAge, "/")

	this.Redirect("/", 301)
	return
}
