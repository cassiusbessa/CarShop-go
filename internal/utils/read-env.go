package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func ReadEnv(env string) string {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	uri := os.Getenv(env)
	if uri == "" {
		fmt.Println("You must set your 'MONGODB URI' environmental variable.")
	}
	return uri
}
