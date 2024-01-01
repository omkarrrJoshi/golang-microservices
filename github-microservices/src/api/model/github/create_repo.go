package github

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssue    bool   `json:"has_issue"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

type CreateRepoResponse struct {
	Id           int64           `json:"id"`
	Name         string          `json:"name"`
	FullName     string          `json:"full_name"`
	Owner        RepoOwner       `json:"owner"`
	Permisssions RepoPermissions `json:"permissions"`
}

type RepoOwner struct {
	Id      int64  `json:"id"`
	Login   string `json:"login"`
	Url     string `json:"url"`
	HtmlUrl string `json:"html_url"`
}

type RepoPermissions struct {
	IsAdmin bool `json:"is_admin"`
	HasPush bool `json:"has_push"`
	HasPull bool `json:"hsa_pull"`
}
