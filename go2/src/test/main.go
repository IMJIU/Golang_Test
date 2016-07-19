package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"test/controllers"
	_ "test/models"
	_ "test/routers"
)

func main() {
	orm.Debug = true
	// beego.EnableAdmin = true
	// beego.AdminHttpAddr = "localhost"
	// beego.AdminHttpPort = 8088
	//orm.RunSyncdb("default", false, true)
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Run()
}
