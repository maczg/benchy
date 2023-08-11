package mixin

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func GetEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func MustGetEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("missing environment variable %s", key)
	}
	return value, nil
}

func MustMapEnv(target *string, key string) {
	value := os.Getenv(key)
	if value == "" {
		logrus.Fatalf("Missing environment variable %s", key)
	}
	*target = value
}
