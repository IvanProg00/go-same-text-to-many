package file

import (
	"io/ioutil"
	"os"
)

// ExistsFile ...
func ExistsFile(fileName string) bool {
	file, err := os.Stat(fileName)
	if err != nil {
		return false
	}
	return !file.IsDir()
}

// ExistsDir ...
func ExistsDir(dirName string) bool {
	dir, err := os.Stat(dirName)
	if err != nil {
		return false
	}
	return dir.IsDir()
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
