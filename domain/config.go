package domain

type ChaosConfig struct {
	Id        string `json:"id"`
	ProjectId string `json:"project_id"`
	Name      string `json:"name"`
	Mode      Mode   `json:"mode"`
	Value     string `json:"value"`
	Response  string `json:"response"`
}

type Mode string

const (
	Latency   Mode = "latency"
	Response  Mode = "response"
	ErrorRate Mode = "error-rate"
)
