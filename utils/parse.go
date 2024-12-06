package utils

import (
	"log"
	"os"
)

func ReadInput(fileName string) string {
	dat, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	// String manupulations
	input := string(dat)
	return input
}
