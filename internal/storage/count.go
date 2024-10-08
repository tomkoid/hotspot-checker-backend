package storage

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func SaveCount(count *int) error {
	if os.Getenv("DEBUG") == "true" {
		log.Printf("Saving count: %d\n", *count)
	}

	configDir, err := os.UserCacheDir()

	if err != nil {
		fmt.Println("save: could not get user config dir")
		configDir = "/"
	}

	// check if file already exists
	_, err = os.Stat(configDir + "/http-server-count.txt")
	if err != nil {
		// create file
		file, err := os.Create(configDir + "/http-server-count.txt")
		if err != nil {
			fmt.Println("save: could not create file")
			return err
		}
		// write count
		file.WriteString(fmt.Sprintf("%d", *count))
		file.Close()

		fmt.Println("save: created new file")
	} else {
		// write count
		file, _ := os.OpenFile(configDir+"/http-server-count.txt", os.O_RDWR, 0700)
		file.WriteString(fmt.Sprintf("%d", *count))
		file.Close()
	}

	return nil
}

func GetCount() int {
	configDir, _ := os.UserCacheDir()

	// check if file already exists
	_, err := os.Stat(configDir + "/http-server-count.txt")
	if err != nil {
		fmt.Println("save: file does not exist")
		return 0
	}

	output, err := os.ReadFile(configDir + "/http-server-count.txt")
	if err != nil {
		fmt.Println("save: could not read file")
		return 0
	}

	outputInt, err := strconv.Atoi(string(output))
	if err != nil {
		fmt.Println("save: could not convert string to int")
		return 0
	}

	return outputInt
}
