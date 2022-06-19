package config

// NewConfig creates a new configuration
func NewConfig() *Config {
	return &Config{
		General: GeneralConfig{
			DatabaseURI: GetDatabaseConnectionUri(),
			ListenAddr:  ":8080",
		},
	}
}
