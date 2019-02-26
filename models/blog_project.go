package models

// BlogProject 本地保存的博客项目
type BlogProject struct {
	ID           int64  `json:"id"`
	NodeID       string `json:"node_id"`
	Name         string `json:"name"`
	FullName     string `json:"full_name"`
	Description  string `json:"description"`
	ContentsURL  string `json:"contents_url"`
	EventsURL    string `json:"events_url"`
	UpdatedAt    string `json:"updated_at"`
	PushedAt     string `json:"pushed_at"`
	AuthorID     int64  `json:"author_id"`
	AuthorNodeID string `json:"author_node_id"`
	AvatarURL    string `json:"avatar_url"`
	ReposURL     string `json:"repos_url"`
	PushEventID  string `json:"push_event_id"`
}
