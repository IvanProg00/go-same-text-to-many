package config_test

import (
	"go-same-text-to-many/pkg/config"
	"testing"
)

func TestNew(t *testing.T) {
	data := `name;age
Ivan;20
John;42
Fred;38`

	expectedData := config.AllLinesConfigs{
		{
			{
				Key:   "name",
				Value: "Ivan",
			},
			{
				Key:   "age",
				Value: "20",
			},
		},
		{
			{
				Key:   "name",
				Value: "John",
			},
			{
				Key:   "age",
				Value: "42",
			},
		},
		{
			{
				Key:   "name",
				Value: "Fred",
			},
			{
				Key:   "age",
				Value: "38",
			},
		},
	}

	result, err := config.New(data)
	if err != nil {
		t.Error(err)
		return
	}

	if len(result) != len(expectedData) {
		t.Error("Expected value is not equal to result.")
		return
	}

	for i, configLine := range result {
		for k, config := range configLine {
			if config.Key != expectedData[i][k].Key {
				t.Errorf("Result Key \"%s\" on line %d is not equal to expected key \"%s\"", config.Key, i, expectedData[i][k].Key)
			}
			if config.Value != expectedData[i][k].Value {
				t.Errorf("Result Value \"%s\" on line %d is not equal to expected value \"%s\"", config.Value, i, expectedData[i][k].Value)
			}
		}
	}
}
