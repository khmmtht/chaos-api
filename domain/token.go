package domain

type Token struct {
	Id        string `json:"id" bson:"id"`
	ProjectId string `json:"project_id" bson:"project_id"`
	Value     string `json:"value" bson:"value"`
	Name      string `json:"name" bson:"name"`
}
