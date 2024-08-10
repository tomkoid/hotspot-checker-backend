package tools

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func setEnv(key, value string) {
	if strings.TrimSpace(os.Getenv(key)) != "" {
		return
	} else {
		fmt.Printf("%s: %s (%s)\n", key, os.Getenv(key), value)
	}

	os.Setenv(key, value)
}

func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env: no env file doesn't exist")
	}

	// set default values
	setEnv("PASSWORD", "esp32")
	setEnv("HOST", "0.0.0.0")
	setEnv("PORT", "3000")
	setEnv("TIMEOUT", "60")
	setEnv("NTFY_HOST", "ntfy.sh")
	setEnv("NTFY_ROOM", "esp32-alerts")
	setEnv("NTFY_TITLE", "Warning")
	setEnv("NTFY_MSG", "Hotspot is down! Turn it back on.")
}
