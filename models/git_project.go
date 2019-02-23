package models

// GitProject Github项目
type GitProject struct {
	ID           int64   `json:"id"`
	NodeID       string  `json:"node_id"`
	Name         string  `json:"name"`
	FullName     string  `json:"full_name"`
	Private      bool    `json:"private"`
	Owner        GitUser `json:"owner"`
	Description  string  `json:"description"`
	HTMLURL      string  `json:"html_url"`
	URL          string  `json:"url"`
	ContentsURL  string  `json:"contents_url"`
	DownloadsURL string  `json:"downloads_url"`
	EventsURL    string  `json:"events_url"`
	LanguagesURL string  `json:"languages_url"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
	PushedAt     string  `json:"pushed_at"`
	GitURL       string  `json:"git_url"`
	SSHURL       string  `json:"ssh_url"`
	SvnURL       string  `json:"svn_url"`
}
