package models

// GitUser Github账号
type GitUser struct {
	ID        int64  `json:"id"`
	NodeID    string `json:"node_id"`
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
	URL       string `json:"url"`
	HTMLURL   string `json:"html_url"`
	ReposURL  string `json:"repos_url"`
}
