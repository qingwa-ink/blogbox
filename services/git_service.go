package services

import (
	"blog/models"
	"blog/tools"
	"bufio"
	"encoding/json"
	"fmt"
	"os"

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

	// gitEventURL := fmt.Sprintf("https://api.github.com/repos/%s/events", gitName)
	// events, err := tools.HTTPGet(gitEventURL)
	// if err != nil {
	// 	logs.Error("Load Project Error : %s", err.Error())
	// 	return
	// }
	// var es []models.GitEvent
	// json.Unmarshal(events, &es)
	// fmt.Printf("%+v", es)

	// gitAPIURL := fmt.Sprintf("https://api.github.com/repos/%s", gitName)
	// body, err := tools.HTTPGet(gitAPIURL)
	// if err != nil {
	// 	logs.Error("Load Project Error : %s", err.Error())
	// 	return
	// }

	// var project models.GitProject
	// json.Unmarshal(body, &project)

	// logs.Debug("project : %+v", project)
	// if project.ID == 0 {
	// 	logs.Error("Init Error : this project is not found.")
	// 	return
	// } else if project.Private == true {
	// 	logs.Error("Init Error : this project is not public.")
	// 	return
	// }

	// download(gitName, "", gitPath)

	content := &models.BlogContent{
		Name:      "Kate",
		Path:      "good/Kate",
		Deep:      1,
		Size:      99,
		ProjectID: 10086,
	}
	// content.Insert()
	cs, err := content.FindAll()
	if err != nil {
		logs.Error("FindAll Error : %s", err.Error())
		return
	}
	logs.Warning("%+v", cs)
}

func download(gitName, filePath, savePath string) {

	path := fmt.Sprintf("%s/%s", savePath, filePath)
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("temp dir is not exist")
			err := os.Mkdir(path, os.ModePerm)
			if err != nil {
				fmt.Printf("mkdir failed![%v]\n", err)
			}
		}
	}

	gitAPIURL := fmt.Sprintf("https://api.github.com/repos/%s/contents/%s", gitName, filePath)
	body, err := tools.HTTPGet(gitAPIURL)
	if err != nil {
		logs.Error("Load Contents Error : %s", err.Error())
		return
	}

	var files []models.GitFile
	json.Unmarshal(body, &files)

	logs.Debug("body : %+v", files)

	for _, file := range files {

		if file.Type == "file" {
			httpBody, err := tools.HTTPGet(file.DownloadsURL)
			if err != nil {
				logs.Error("Download File Error: %s", err.Error())
				return
			}
			path := fmt.Sprintf("%s/%s", savePath, file.Path)
			file, err := os.Create(path)
			if err != nil {
				logs.Error("Save File Error: %s", err.Error())
				return
			}
			writer := bufio.NewWriter(file)
			writer.Write(httpBody)
			writer.Flush()
		} else if file.Type == "dir" {
			download(gitName, file.Path, savePath)
		}
	}
}
