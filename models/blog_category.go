package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// BlogCategory 本地保存的博客专栏
type BlogCategory struct {
	ID        int64  `orm:"column(id)"`
	ProjectID int64  `orm:"column(project_id)"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Deep      int    `json:"deep"`
	AvatarURL string `orm:"column(avatar_url)"`
}

// TableName 获取本表的表名
func (category *BlogCategory) TableName() string {
	return "blog_category"
}

// Insert 插入数据
func (category *BlogCategory) Insert() {

	o := orm.NewOrm()
	o.Using("default")

	fmt.Println(o.Insert(category))
}

// FindAll 查找所有数据
func (category *BlogCategory) FindAll() (cs []BlogCategory, err error) {

	o := orm.NewOrm()
	o.Using("default")
	_, err = o.QueryTable(category.TableName()).All(&cs)

	return cs, err
}

// DeleteAll 删除所有数据
func (category *BlogCategory) DeleteAll() (err error) {

	o := orm.NewOrm()
	o.Using("default")
	_, err = o.QueryTable(category.TableName()).Filter("id__gt", "0").Delete()

	return err
}
