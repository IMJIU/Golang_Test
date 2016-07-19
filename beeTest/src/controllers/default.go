package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
	c.Ctx.WriteString("appname2:"+beego.AppName)
		c.Ctx.WriteString("appname2:"+beego.AppConfig.String("appname"))
}
