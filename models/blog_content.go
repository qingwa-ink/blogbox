package models

// BlogContent 本地保存的博客内容
type BlogContent struct {
	ID           int64  `json:"id"`
	ProjectID    int64  `json:"project_id"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	Deep         int    `json:"deep"`
	Size         int64  `json:"size"`
	DownloadsURL string `json:"download_url"`
}
