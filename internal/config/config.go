package config

import (
	"github.com/joho/godotenv"
)

// Load loads environment
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

// ClientConfig interface for ClientConfig
type ClientConfig interface {
	AuthServerAddress() string
	ChatServerAddress() string
	TLSCertFile() string
	TLSKeyFile() string
}
