package models

// BlogCategory 本地保存的博客专栏
type BlogCategory struct {
	ProjectID int64  `json:"project_id"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Deep      int    `json:"deep"`
}
