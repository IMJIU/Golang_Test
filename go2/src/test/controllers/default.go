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
	// c.TplNames = "index.tpl"
	c.TplNames = "home.html"
	c.Data["IsGood"] = true
	c.Data["IsFalse"] = false

	//c.Ctx.WriteString("appname:" + beego.AppName)
	//c.Ctx.WriteString("appname:" + beego.AppConfig.String("appname"))

	// beego.Trace("hello")
	// beego.Debug("debug")
	// beego.SetLevel(beego.LevelDebug)
	// beego.Info("info hello")

	type Person struct {
		Name string
		Age  int
	}
	c.Data["p"] = &Person{"zlf", 18}
	nums := []int{1, 2, 3, 4, 5}
	c.Data["ns"] = nums
	c.Data["v1"] = "v1"
	c.Data["htmlVar"] = "<font color='red'>red color</font>"
}
