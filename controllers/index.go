package controllers

import (
	"blog/models"

	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {

	var content models.BlogContent
	var category models.BlogCategory
	var project models.BlogProject

	bps, _ := project.FindAll()
	cat := c.Input().Get("cat")

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	if cat == "" {
		c.Data["Contents"], _ = content.FindContents(1, "")
	} else {
		c.Data["Contents"], _ = content.FindContents(1, cat)
	}
	c.Data["Categories"], _ = category.FindCategories(0, "")
	c.Data["Project"] = bps[0]
	c.Data["Nickname"] = beego.AppConfig.String("nickname")
	c.Data["Information"] = beego.AppConfig.String("information")
	c.TplName = "index.html"
}
