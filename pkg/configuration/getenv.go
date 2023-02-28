package configuration

import (
	"os"
)

func getenv(key string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}

	panic("environment variable with key \"" + key + "\" was not set")
}
