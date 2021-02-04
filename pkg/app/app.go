package app

import (
	"fmt"
	"go-same-text-to-many/pkg/config"
	"go-same-text-to-many/pkg/file"
	"os"
	"path/filepath"
)

const (
	// NumArguments ...
	NumArguments = 3
)

// Run ...
func Run() {
	arguments := os.Args[1:]
	if len(arguments) < NumArguments {
		fmt.Printf("Minimum %d arguments\n", NumArguments)
		return
	}

	configFile, err := filepath.Abs(arguments[0])
	contentFile, err := filepath.Abs(arguments[1])
	outputFile, err := filepath.Abs(arguments[2])

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	searchFiles := []file.SearchFile{
		{
			FileName:    configFile,
			IfNotExists: "File with configs not exists",
		},
		{
			FileName:    contentFile,
			IfNotExists: "File with contents not exists",
		},
	}

	if err = file.ExistsFiles(searchFiles); err != nil {
		fmt.Println(err)
		return
	}

	configs, err := file.ReadFile(configFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	configList, err := config.New(configs)
	if err != nil {
		fmt.Println(err)
		return
	}

	content, err := file.ReadFile(contentFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	modifiedContent := configList.TransformData(content)

	file.WriteFile(outputFile, modifiedContent)
}
