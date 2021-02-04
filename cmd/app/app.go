package main

import (
	"errors"
	"fmt"
	"go-same-text-to-many/pkg/config"
	"go-same-text-to-many/pkg/file"
	"log"
	"os"
	"path/filepath"
)

const (
	// NumArguments ...
	NumArguments = 3
)

var (
	// AssetsPath ...
	AssetsPath = "assets"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) < NumArguments {
		fmt.Printf("Minimum %d arguments\n", NumArguments)
		os.Exit(0)
	}

	configFile, err := filepath.Abs(arguments[0])
	contentFile, err := filepath.Abs(arguments[1])
	outputFile, err := filepath.Abs(arguments[2])

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	if err = ExistsFiles(configFile, contentFile); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	if !file.ExistsFile(contentFile) {
		fmt.Println("File with contents not exists")
		os.Exit(0)
	}

	configs, err := file.ReadFile(configFile)
	if err != nil {
		log.Fatalln(err)
		os.Exit(0)
	}

	configList, err := config.New(configs)
	if err != nil {
		log.Fatalln(err)
		os.Exit(0)
	}

	content, err := file.ReadFile(contentFile)
	if err != nil {
		log.Fatalln(err)
		os.Exit(0)
	}

	modifiedContent := configList.TransformData(content)

	file.WriteFile(outputFile, modifiedContent)
}

// ExistsFiles ...
func ExistsFiles(configFile, contentFile string) error {
	if !file.ExistsFile(configFile) {
		return errors.New("File with configs not exists")
	}
	if !file.ExistsFile(contentFile) {
		return errors.New("File with contents not exists")
	}

	return nil
}
