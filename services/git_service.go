package services

import (
	"blog/models"
	"blog/tools"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// RefreshGitData : 刷新github的博客数据
func RefreshGitData() {

	gitPath := beego.AppConfig.String("git_path")
	sqlitePath := beego.AppConfig.String("sqlite_path")
	gitName := beego.AppConfig.String("git_name")

	fmt.Println(gitPath)
	fmt.Println(sqlitePath)
	fmt.Println(gitName)

	gitAPIURL := fmt.Sprintf("https://api.github.com/repos/%s", gitName)
	body, err := tools.HTTPGet(gitAPIURL)
	if err != nil {
		logs.Warn("something wrong here: %s", err.Error())
		return
	}

	var project models.GitProject
	json.Unmarshal(body, &project)

	logs.Debug("project : %+v", project)
	if project.ID == 0 {
		logs.Error("Init Error : this project is not found.")
		return
	} else if project.Private == true {
		logs.Error("Init Error : this project is not public.")
		return
	}

	download(gitName, "", gitPath)
}

func download(gitName, filePath, savePath string) {

	gitAPIURL := fmt.Sprintf("https://api.github.com/repos/%s/contents/%s", gitName, filePath)
	body, err := tools.HTTPGet(gitAPIURL)
	if err != nil {
		logs.Warn("something wrong here: %s", err.Error())
		return
	}

	logs.Warn("body : %s", string(body))
}
