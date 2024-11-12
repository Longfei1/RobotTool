package config

type ServerConfig struct {
	Name     string `json:"name"`
	Addr     string `json:"addr"`
	UserId   string `json:"userId"`
	Password string `json:"password"`
}

type AppModeInfo struct {
	Name    string `json:"name"`
	UseWeb  bool   `json:"useWeb"`
	WebPort int    `json:"webPort"`
	UiUrl   string `json:"uiUrl"`
}

type AppConfig struct {
	Mode     string         `json:"mode"`
	ModeList []*AppModeInfo `json:"modeList"`
}
