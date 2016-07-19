package controllers

import (
	// "fmt"
	"github.com/astaxie/beego"
	"test/models"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	cname := c.Input().Get("cname")
	op := c.Input().Get("op")

	switch op {
	case "add":
		err := models.AddCategory(cname)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 301)
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
	}

	beego.Debug("cname", cname)
	//c.Redirect("/category", 301)
	c.Data["isC"] = true
	c.TplNames = "category.html"
	var err error
	c.Data["list"], err = models.GetAllCategory()
	if err != nil {
		beego.Error(err)
	}
}

func (c *CategoryController) Post() {
	cname := c.Input().Get("cname")
	op := c.Input().Get("op")

	switch op {
	case "add":
		err := models.AddCategory(cname)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 301)
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
	}

	beego.Debug("cname", cname)
	c.Redirect("/category", 301)
	c.Data["isC"] = true
	c.TplNames = "category.html"
	return
}
