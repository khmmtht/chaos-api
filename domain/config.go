package domain

type ChaosConfig struct {
	Id        string `json:"id" bson:"id"`
	ProjectId string `json:"project_id" bson:"project_id"`
	Name      string `json:"name" bson:"name"`
	Mode      Mode   `json:"mode" bson:"mode"`
	Value     string `json:"value" bson:"value"`
	Response  string `json:"response" bson:"response"`
}

type Mode string

const (
	Latency   Mode = "latency"
	Response  Mode = "response"
	ErrorRate Mode = "error-rate"
)
