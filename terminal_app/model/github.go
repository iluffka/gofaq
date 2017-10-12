package model

type GitHubResponse struct {
	ID			int64	`json:"id"`
	Name		string	`json:"name"`
	FullName	string	`json:"full_name"`
	Owner		map[string]interface{}	`json:"owner"`
}
