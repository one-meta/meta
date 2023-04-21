package config

// Auth 包含jwt 和 casbin
type Auth struct {
	Enable bool   `json:"enable"`
	JWT    JWT    `json:"JWT"`
	Casbin Casbin `json:"Casbinx"`
}
type JWT struct {
	Hmac string `json:"hmac"`
	Key  string `json:"key"`
	TTL  int    `json:"TTL"`
}
type Casbin struct {
	ModelPath string `json:"modelPath"`
}
