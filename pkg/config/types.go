package config

// Config is the root configuration
type Config struct {
	General GeneralConfig `json:"general"`
	JWT     JWTConfig     `json:"jwt"`
}

type GeneralConfig struct {
	DatabaseURI string `json:"databaseURI"`
	ListenAddr  string `json:"listenAddr"`
	BlockList   string `json:"blockList"`
}

type JWTConfig struct {
	Secret       string `json:"secret"`
	ValidMinutes int64  `json:"validMinutes"`
	Realm        string `json:"realm"`
}
