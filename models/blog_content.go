package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// BlogContent 本地保存的博客内容
type BlogContent struct {
	ID           int64  `orm:"column(id)"`
	ProjectID    int64  `orm:"column(project_id)"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	Deep         int    `json:"deep"`
	Size         int64  `json:"size"`
	Ctime        int64  `json:"ctime"`
	Description  string `json:"description"`
	DownloadsURL string `orm:"column(download_url)"`
}

func (content *BlogContent) TableName() string {
	return "blog_content"
}

func (content *BlogContent) Insert() {

	o := orm.NewOrm()
	o.Using("default")

	fmt.Println(o.Insert(content))
}

func (content *BlogContent) FindAll() (cs []BlogContent, err error) {

	o := orm.NewOrm()
	o.Using("default")
	_, err = o.QueryTable(content.TableName()).All(&cs)

	return cs, err
}
