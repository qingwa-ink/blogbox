package controllers

import (
	"blog/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	var content models.BlogContent
	var category models.BlogCategory
	var project models.BlogProject

	bps, _ := project.FindAll()

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Contents"], _ = content.FindContents(1, "")
	c.Data["Categories"], _ = category.FindCategories(0, "")
	c.Data["Project"] = bps[0]
	c.Data["Nickname"] = beego.AppConfig.String("nickname")
	c.Data["Information"] = beego.AppConfig.String("information")
	c.TplName = "index.html"
}
