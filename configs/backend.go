package configs

import (
	"fmt"
	"os"
	"strconv"
)

func intEnv(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return fallback
}

var Port = func() string {
	if v := os.Getenv("PORT"); v != "" {
		return v
	}
	return "8080"
}

var BaseURL = func() string {
	if v := os.Getenv("BASE_URL"); v != "" {
		return v
	}
	return fmt.Sprintf("http://localhost:%s", Port())
}()

var RedisURL = func() string {
	if v := os.Getenv("REDIS_URL"); v != "" {
		return v
	}
	return "localhost:6379"
}()

var RedisPassword = os.Getenv("REDIS_PASSWORD")

var RateLimitRPS = intEnv("RATE_LIMIT_RPS", 2)
var MaxURLs = intEnv("MAX_URLS", 10)
var URLTTLMinutes = intEnv("URL_TTL_MINUTES", 1)
