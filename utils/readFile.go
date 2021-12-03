package utils

import (
	// "fmt"
	"io/ioutil"
	// "log"
)

func ReadFile(filename string) (string, error) {

	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	// Convert []byte to string and print to screen
	text := string(content)
	// fmt.Println(text)

	return text, nil
}
