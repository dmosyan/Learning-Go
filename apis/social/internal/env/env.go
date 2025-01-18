package env

import (
	"os"
	"strconv"
)

func GetEnv(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return val
}

func GetIntEnv(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	valAsint, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return valAsint
}
