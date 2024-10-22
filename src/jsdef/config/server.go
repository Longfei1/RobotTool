package config

type ServerConfig struct {
	Name     string `json:"name"`
	Addr     string `json:"addr"`
	UserId   string `json:"userId"`
	Password string `json:"password"`
}
