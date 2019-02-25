package models

// GitFile Github文件
type GitFile struct {
	Name         string `json:"name"`
	Path         string `json:"path"`
	Size         int64  `json:"size"`
	Sha          string `json:"sha"`
	HTMLURL      string `json:"html_url"`
	URL          string `json:"url"`
	DownloadsURL string `json:"download_url"`
	GitURL       string `json:"git_url"`
	Type         string `json:"type"`
}
