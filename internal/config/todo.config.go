package config

type Config struct {
	Port string
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
