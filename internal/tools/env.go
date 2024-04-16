package tools

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func setEnv(key, value string) {
	if os.Getenv(key) != "" {
		return
	}

	os.Setenv(key, value)
}

func EnvLoad() {
	// set default values
	setEnv("PASSWORD", "esp32")
	setEnv("NTFY_ROOM", "esp32-alerts")
	setEnv("NTFY_TITLE", "Warning")
	setEnv("NTFY_MSG", "Hotspot is down! Turn it back on.")

	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env: no env file doesn't exist")

		return
	}
}
