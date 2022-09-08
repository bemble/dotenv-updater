package config

import (
	"os"
)

func IsDeleteAllowed() bool {
	return os.Getenv("DELETE_ALLOWED") == "1" || os.Getenv("DELETE_ALLOWED") == "true" || os.Getenv("DELETE_ALLOWED") == "True"
}
