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

// TableName 获取本表的表名
func (content *BlogContent) TableName() string {
	return "blog_content"
}

// Insert 插入数据
func (content *BlogContent) Insert() {

	o := orm.NewOrm()
	o.Using("default")

	fmt.Println(o.Insert(content))
}

// FindContentByPath 根据路径查找数据
func (content *BlogContent) FindContentByPath(path string) (err error) {

	o := orm.NewOrm()
	o.Using("default")
	err = o.QueryTable(content.TableName()).Filter("path", path).One(content)

	return err
}

// FindAll 查找所有数据
func (content *BlogContent) FindAll() (cs []BlogContent, err error) {

	o := orm.NewOrm()
	o.Using("default")
	_, err = o.QueryTable(content.TableName()).All(&cs)

	return cs, err
}

// FindContents 查找所有数据
func (content *BlogContent) FindContents(deep int, path string) (cs []BlogContent, err error) {

	o := orm.NewOrm()
	o.Using("default")
	_, err = o.QueryTable(content.TableName()).Filter("deep__gte", deep).Filter("path__istartswith", path).All(&cs)

	return cs, err
}

// DeleteAll 删除所有数据
func (content *BlogContent) DeleteAll() (err error) {

	o := orm.NewOrm()
	o.Using("default")
	_, err = o.QueryTable(content.TableName()).Filter("id__gt", "0").Delete()

	return err
}
