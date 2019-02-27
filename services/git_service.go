package services

import (
	"blog/models"
	"blog/tools"
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// RefreshGitData : 刷新github的博客数据
func RefreshGitData() {

	gitPath := beego.AppConfig.String("git_path")
	sqlitePath := beego.AppConfig.String("sqlite_path")
	gitName := beego.AppConfig.String("git_name")

	// TODO 第一步，检测项目是否存在，是否公有
	gitAPIURL := fmt.Sprintf("https://api.github.com/repos/%s", gitName)
	body, err := tools.HTTPGet(gitAPIURL)
	if err != nil {
		logs.Error("Load Project Error : %s", err.Error())
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

	// TODO 第二步，检测当前版本
	pushID := "0"
	gitEventURL := fmt.Sprintf("https://api.github.com/repos/%s/events", gitName)
	events, err := tools.HTTPGet(gitEventURL)
	if err != nil {
		logs.Error("Load Project Events Error : %s", err.Error())
		return
	}
	var es []models.GitEvent
	json.Unmarshal(events, &es)
	fmt.Printf("%+v", es)
	for _, e := range es {
		if e.Type == "PushEvent" || e.Type == "CreateEvent" {
			pushID = e.ID
			break
		}
	}

	// TODO 第三步，检测数据库，查看项目是否存在
	var bp models.BlogProject
	bps, _ := bp.FindAll()
	if len(bps) > 0 && bps[0].FullName == gitName {
		bps[0].Description = project.Description
		bps[0].UpdatedAt = project.UpdatedAt
		bps[0].PushedAt = project.PushedAt
		bps[0].AvatarURL = project.Owner.AvatarURL
		// 检测项目是否有新的提交
		if pushID == bps[0].PushEventID {
			bps[0].Update()
			return
		}
		// 更新版本号
		bps[0].PushEventID = pushID
		bps[0].Update()
		bp = bps[0]
	} else {
		// 保存项目
		bp.DeleteAll()
		bp.ID = project.ID
		bp.NodeID = project.NodeID
		bp.Name = project.Name
		bp.FullName = project.FullName
		bp.Description = project.Description
		bp.ContentsURL = project.ContentsURL
		bp.EventsURL = project.EventsURL
		bp.UpdatedAt = project.UpdatedAt
		bp.PushedAt = project.PushedAt
		bp.AvatarURL = project.Owner.AvatarURL
		bp.AuthorID = project.Owner.ID
		bp.ReposURL = project.Owner.ReposURL
		bp.AuthorNodeID = project.Owner.NodeID
		bp.PushEventID = pushID
		bp.Insert()
	}

	// 第四步，下载更新数据
	var category models.BlogCategory
	var content models.BlogContent
	category.DeleteAll()
	content.DeleteAll()

	rd, err := ioutil.ReadDir(gitPath)
	for _, fi := range rd {
		fileName := fmt.Sprintf("%s/%s", gitPath, fi.Name())
		fmt.Println(fileName)
		if fileName != sqlitePath {
			if fi.IsDir() {
				err := os.RemoveAll(fileName)
				if err != nil {
					logs.Warning("Remove file failed : ", err.Error())
				}
			} else {
				err := os.Remove(fileName)
				if err != nil {
					logs.Warning("Remove file failed : ", err.Error())
				}
			}
		}
	}

	download(bp, gitName, "", gitPath)
}

func download(bp models.BlogProject, gitName, filePath, savePath string) {

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

	for _, gitFile := range files {

		if gitFile.Type == "file" {
			httpBody, err := tools.HTTPGet(gitFile.DownloadsURL)
			if err != nil {
				logs.Error("Download File Error: %s", err.Error())
				return
			}
			path := fmt.Sprintf("%s/%s", savePath, gitFile.Path)
			file, err := os.Create(path)
			if err != nil {
				logs.Error("Save File Error: %s", err.Error())
				return
			}
			writer := bufio.NewWriter(file)
			writer.Write(httpBody)
			writer.Flush()

			content := &models.BlogContent{
				ProjectID:    bp.ID,
				Name:         gitFile.Name,
				Path:         gitFile.Path,
				Deep:         strings.Count(gitFile.Path, "/"),
				Size:         gitFile.Size,
				DownloadsURL: gitFile.DownloadsURL,
			}
			content.Insert()
		} else if gitFile.Type == "dir" {
			category := &models.BlogCategory{
				ProjectID: bp.ID,
				Name:      gitFile.Name,
				Path:      gitFile.Path,
				Deep:      strings.Count(gitFile.Path, "/"),
			}
			category.Insert()
			download(bp, gitName, gitFile.Path, savePath)
		}
	}
}
