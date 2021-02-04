package file

import (
	"errors"
	"io/ioutil"
	"os"
)

// SearchFile ...
type SearchFile struct {
	FileName    string
	IfNotExists string
}

// ExistsFile ...
func ExistsFile(fileName string) bool {
	file, err := os.Stat(fileName)
	if err != nil {
		return false
	}
	return !file.IsDir()
}

// ExistsFiles ...
func ExistsFiles(searchFiles []SearchFile) error {
	for _, searchFile := range searchFiles {
		if !ExistsFile(searchFile.FileName) {
			return errors.New(searchFile.IfNotExists)
		}
	}

	return nil
}

// ReadFile ...
func ReadFile(fileName string) (string, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// WriteFile ...
func WriteFile(fileName, content string) error {
	err := ioutil.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}
