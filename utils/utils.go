package utils

import (
	"os"
	"strings"
)

func LoadFileAsLines(filename string) ([]string, error) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(fileContent), "\n"), nil
}
