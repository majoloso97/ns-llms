package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("Error loading .env file: %s", err)
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
