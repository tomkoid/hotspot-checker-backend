package tools

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env: no env file doesn't exist")

		// set default values
		os.Setenv("PASSWORD", "esp32")
		os.Setenv("NTFY_ROOM", "esp32-alerts")
		os.Setenv("NTFY_TITLE", "Warning")
		os.Setenv("NTFY_MSG", "Hotspot is down! Turn it back on.")

		return
	}
}
