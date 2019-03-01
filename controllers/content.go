package controllers

import (
	"blog/models"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/astaxie/beego"
)

type ContentController struct {
	beego.Controller
}

func (c *ContentController) Get() {

	var content models.BlogContent
	var category models.BlogCategory
	var project models.BlogProject

	bps, _ := project.FindAll()

	path := c.Input().Get("path")
	content.FindContentByPath(path)

	gitPath := beego.AppConfig.String("git_path")
	mdPath := fmt.Sprintf("%s/%s", gitPath, path)
	f, err := os.Open(mdPath)
	if err != nil {
		return
	}
	datas, _ := ioutil.ReadAll(f)

	c.Data["Content"] = string(datas)
	c.Data["Categories"], _ = category.FindCategories(0, "")
	c.Data["Project"] = bps[0]
	c.Data["Nickname"] = beego.AppConfig.String("nickname")
	c.Data["Information"] = beego.AppConfig.String("information")
	c.TplName = "content.html"
}
