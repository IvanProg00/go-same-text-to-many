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

func TestTransformData(t *testing.T) {
	data := config.AllLinesConfigs{
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
	result := data.TransformData("Hello {name}. I'm {age} years old.")

	expectedData := "Hello Ivan. I'm 20 years old.\nHello John. I'm 42 years old.\nHello Fred. I'm 38 years old.\n"

	if result != expectedData {
		t.Error("Result not equals to expected data.")
	}
}

func TestKeyInContent(t *testing.T) {
	data := []string{
		"one",
		"two",
	}
	result := []string{}
	for _, item := range data {
		result = append(result, config.KeyInContent(item))
	}

	expectedData := []string{
		"{one}",
		"{two}",
	}

	if len(result) != len(expectedData) {
		t.Errorf("Length of result <%d> not equals to length of expected data <%d>", len(result), len(expectedData))
		return
	}

	for i, item := range result {
		if item != expectedData[i] {
			t.Errorf("Result Key %s not equals to expected key %s", item, expectedData[i])
		}
	}
}
