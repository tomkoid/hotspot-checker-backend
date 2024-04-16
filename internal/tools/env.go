package tools

import (
	"fmt"

	"github.com/joho/godotenv"
)

func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env: no env file doesn't exist")
	}
}
