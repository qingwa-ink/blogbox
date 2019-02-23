package services

import (
	"blog/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

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
	httpClient := &http.Client{
		Timeout: 20 * time.Second,
	}
	res, err := httpClient.Get(gitAPIURL)
	if err != nil {
		logs.Warn("something wrong here: %s", err.Error())
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		logs.Warn("something wrong here: %s", err.Error())
		return
	}

	var project models.GitProject
	json.Unmarshal(body, &project)

	fmt.Printf("project : %+v\n", project)
}
