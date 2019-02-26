package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {

	sqlitePath := beego.AppConfig.String("sqlite_path")

	orm.RegisterDriver("sqlite3", orm.DRSqlite)

	orm.RegisterDataBase("default", "sqlite3", sqlitePath)

	orm.RegisterModel(new(BlogProject), new(BlogCategory), new(BlogContent))

}
