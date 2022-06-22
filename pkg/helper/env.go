package helper

import "os"

// Get env variable or fallback value if env variable is empty
func GetEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
