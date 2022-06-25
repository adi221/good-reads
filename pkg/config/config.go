package config

// NewConfig creates a new configuration
func NewConfig() *Config {
	return &Config{
		General: GeneralConfig{
			DatabaseURI: GetDatabaseConnectionUri(),
			ListenAddr:  ":8080",
			BlockList:   "https://raw.githubusercontent.com/anudeepND/blacklist/master/adservers.txt",
		},
		JWT: GetJWTConfig(),
	}
}
