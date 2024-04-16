package tools

import "github.com/joho/godotenv"

func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		panic(".env: Error loading .env file")
	}
}
