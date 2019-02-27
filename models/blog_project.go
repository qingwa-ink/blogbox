package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// BlogProject 本地保存的博客项目
type BlogProject struct {
	ID           int64  `orm:"column(id)"`
	NodeID       string `orm:"column(node_id)"`
	Name         string `json:"name"`
	FullName     string `json:"full_name"`
	Description  string `json:"description"`
	ContentsURL  string `orm:"column(contents_url)"`
	EventsURL    string `orm:"column(events_url)"`
	UpdatedAt    string `json:"updated_at"`
	PushedAt     string `json:"pushed_at"`
	AuthorID     int64  `orm:"column(author_id)"`
	AuthorNodeID string `orm:"column(author_node_id)"`
	AvatarURL    string `orm:"column(avatar_url)"`
	ReposURL     string `orm:"column(repos_url)"`
	PushEventID  string `orm:"column(push_event_id)"`
}

// TableName 获取本表的表名
func (project *BlogProject) TableName() string {
	return "blog_project"
}

// Insert 插入数据
func (project *BlogProject) Insert() {

	o := orm.NewOrm()
	o.Using("default")

	fmt.Println(o.Insert(project))
}

// Update 修改数据
func (project *BlogProject) Update() {

	o := orm.NewOrm()
	o.Using("default")

	fmt.Println(o.Update(project, "description", "updated_at", "pushed_at", "avatar_url", "push_event_id"))
}

// FindAll 查找所有数据
func (project *BlogProject) FindAll() (cs []BlogProject, err error) {

	o := orm.NewOrm()
	o.Using("default")
	_, err = o.QueryTable(project.TableName()).All(&cs)

	return cs, err
}

// DeleteAll 删除所有数据
func (project *BlogProject) DeleteAll() (err error) {

	o := orm.NewOrm()
	o.Using("default")
	_, err = o.QueryTable(project.TableName()).Filter("id__gt", "0").Delete()

	return err
}
