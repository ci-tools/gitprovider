package providers

// Exported models

// Repo ...
type Repo struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	Name     string `json:"name"`
	WebURL   string `json:"webURL"`
	SSHURL   string `json:"sshURL"`
}

// HookEvent type required
type HookEvent string

// Enum<HookEvent>
const (
	HookEventPush        HookEvent = "push"
	HookEventPullRequest HookEvent = "pull-request"
)

// Hook ...
type Hook struct {
	ID       int      `json:"id"`
	URL      string   `json:"url"`
	RepoID   int      `json:"repoID"`
	Events   []string `json:"events"`
	CheckSSL bool     `json:"checkSSL"`
}
