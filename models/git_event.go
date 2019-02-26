package models

// GitEvent Github事件，包括上传，拉取等
type GitEvent struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Public    bool   `json:"public"`
	CreatedAt string `json:"created_at"`
}
