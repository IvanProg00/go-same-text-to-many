package config

import (
	"errors"
	"fmt"
	"strings"
)

const (
	// MinLines ...
	MinLines = 2
	// Separator ...
	Separator = ";"
)

// AllLinesConfigs ...
type AllLinesConfigs []LineConfigs

// LineConfigs ...
type LineConfigs []Config

// Config ...
type Config struct {
	Key   string
	Value string
}

// New ...
func New(data string) (AllLinesConfigs, error) {
	dataLines := strings.Split(data, "\n")

	if len(dataLines) < MinLines {
		err := fmt.Sprintf("In File with variables must be more %v than lines.", MinLines)
		return nil, errors.New(err)
	}

	var configList []LineConfigs
	keys := strings.Split(dataLines[0], Separator)

	for i := 1; i < len(dataLines); i++ {
		values := strings.Split(dataLines[i], Separator)

		var configs []Config
		for i, value := range values {
			configs = append(configs, Config{
				Key:   keys[i],
				Value: value,
			})
		}

		configList = append(configList, configs)
	}

	return configList, nil
}

// TransformData ...
func (configs AllLinesConfigs) TransformData(content string) string {
	outputContent := ""
	for _, lineConfigs := range configs {
		tmpContent := content
		for _, config := range lineConfigs {
			tmpContent = strings.ReplaceAll(tmpContent, KeyInContent(config.Key), config.Value)
		}
		outputContent += tmpContent + "\n"
	}

	return outputContent
}

// KeyInContent ....
func KeyInContent(key string) string {
	return "{" + key + "}"
}
