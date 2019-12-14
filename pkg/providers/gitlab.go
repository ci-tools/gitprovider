package providers

import (
	"fmt"
	"gitprovider/pkg/requests"
)

// ProjectNamespace gitlab api model
type ProjectNamespace struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Kind     string `json:"kind"`
	FullPath string `json:"full_path"`
}

// Project gilab api model
type Project struct {
	ID                int              `json:"id"`
	Description       string           `json:"description"`
	DefaultBranch     string           `json:"default_branch"`
	SSHURLToRepo      string           `json:"ssh_url_to_repo"`
	HTTPURLToRepo     string           `json:"http_url_to_repo"`
	WebURL            string           `json:"web_url"`
	ReadmeURL         string           `json:"readme_url"`
	Name              string           `json:"name"`
	NameWithNamespace string           `json:"name_with_namespace"`
	Path              string           `json:"path"`
	PathWithNamespace string           `json:"path_with_namespace"`
	CreatedAt         string           `json:"created_at"`
	LastActivityAt    string           `json:"last_activity_at"`
	ForksCount        int              `json:"forks_count"`
	AvatarURL         string           `json:"avatar_url"`
	StarCount         int              `json:"star_count"`
	TagList           []string         `json:"tag_list"`
	Namespace         ProjectNamespace `json:"namespace"`
}

// GitlabHook ...
type GitlabHook struct {
	ID                       int    `json:"id"`
	URL                      string `json:"url"`
	ProjectID                int    `json:"project_id"`
	PushEvent                bool   `json:"push_events"`
	PushEventsBranchFilter   string `json:"push_events_branch_filter"`
	IssuesEvents             bool   `json:"issues_events"`
	ConfidentialIssuesEvents bool   `json:"confidential_issues_events"`
	MergeRequestsEvents      bool   `json:"merge_requests_events"`
	TagPushEvents            bool   `json:"tag_push_events"`
	NoteEvents               bool   `json:"note_events"`
	JobEvents                bool   `json:"job_events"`
	PipelineEvents           bool   `json:"pipeline_events"`
	WikiPageEvents           bool   `json:"wiki_page_events"`
	EnableSSLVerification    bool   `json:"enable_ssl_verification"`
	CreatedAt                string `json:"created_at"`
}

// Gitlab provider
type Gitlab struct {
	Server string
	Token  string
}

// Repos get list of repositories
func (gitlab Gitlab) Repos() ([]Repo, error) {
	req := &requests.Request{
		URL: fmt.Sprintf("%s/api/v4/projects", gitlab.Server),
		Headers: map[string]string{
			"PRIVATE-TOKEN": gitlab.Token,
		},
	}
	projects := []Project{}
	if err := req.Get(&projects); err != nil {
		return nil, err
	}
	repos := []Repo{}
	for _, project := range projects {
		repos = append(repos, Repo{
			ID:       project.ID,
			Name:     project.Name,
			WebURL:   project.WebURL,
			SSHURL:   project.SSHURLToRepo,
			FullName: project.PathWithNamespace,
		})
	}
	return repos, nil
}

// Hooks get project hooks
func (gitlab Gitlab) Hooks(repo Repo) ([]Hook, error) {
	req := &requests.Request{
		URL: fmt.Sprintf("%s/api/v4/projects/%d/hooks", gitlab.Server, repo.ID),
		Headers: map[string]string{
			"PRIVATE-TOKEN": gitlab.Token,
		},
	}
	gitlabHooks := []GitlabHook{}
	if err := req.Get(&gitlabHooks); err != nil {
		return nil, err
	}
	hooks := []Hook{}
	for _, gitlabHook := range gitlabHooks {
		hooks = append(hooks, Hook{
			ID:       gitlabHook.ID,
			RepoID:   gitlabHook.ProjectID,
			URL:      gitlabHook.URL,
			CheckSSL: gitlabHook.EnableSSLVerification,
			Events:   []string{},
		})
	}
	return hooks, nil
}

// SetHook get project hooks
// func (gitlab Gitlab) SetHook(hook GitlabHook) error {
// 	req := &requests.Request{
// 		URL: fmt.Sprintf("%s/api/v4/projects/%d/hooks", gitlab.Server, project.ID),
// 		Headers: map[string]string{
// 			"PRIVATE-TOKEN": gitlab.Token,
// 		},
// 	}
// 	var hooks []GitlabHook
// 	req.Get(hooks)
// 	return nil
// }
