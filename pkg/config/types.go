package config

// Config is the root configuration
type Config struct {
	General GeneralConfig `json:"general"`
}

type GeneralConfig struct {
	DatabaseURI string `json:"databaseURI"`
	ListenAddr  string `json:"listenAddr"`
	BlockList   string `json:"blockList"`
}
