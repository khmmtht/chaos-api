package domain

type Token struct {
	Id        string `json:"id"`
	ProjectId string `json:"project_id"`
	Value     string `json:"value"`
	Name      string `json:"name"`
}
