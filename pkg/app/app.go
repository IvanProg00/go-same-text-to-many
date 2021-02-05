package app

import (
	"fmt"
	"go-same-text-to-many/pkg/config"
	"go-same-text-to-many/pkg/file"
	"os"
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

	configFile := arguments[0]
	contentFile := arguments[1]
	outputFile := arguments[2]

	searchFiles := []file.SearchFile{
		{
			FileName:    configFile,
			IfNotExists: fmt.Sprintf("File with configs <%s> not exists", configFile),
		},
		{
			FileName:    contentFile,
			IfNotExists: fmt.Sprintf("File with contents <%s> not exists", contentFile),
		},
	}

	if err := file.ExistsFiles(searchFiles); err != nil {
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
