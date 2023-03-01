package configuration

import (
	"os"
)

func getenv(key string) string {
	// Return empty string to all variables to unit tests
	if v, e := os.LookupEnv("SETUP_ENV"); !e || v != "PROD" {
		return ""
	}

	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}

	panic("environment variable with key \"" + key + "\" was not set")
}
