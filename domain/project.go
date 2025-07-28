package domain

type Project struct {
	Id   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}
