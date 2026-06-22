package configs

import (
	"fmt"
	"os"
)

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
