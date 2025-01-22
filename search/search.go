package search

import (
	"log"
	"os"
	"strings"
)

// reads entire file into memory and uses strings.Contains to
// search for text
func SearchTextInFile(fname, text string) (bool, error) {
	log.Printf("searching for text '%v' in file %v\n", text, fname)

	data, err := os.ReadFile(fname)
	if err != nil {
		return false, err
	}

	found := strings.Contains(string(data), text)
	return found, nil
}
