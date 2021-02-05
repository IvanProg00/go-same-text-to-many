package file_test

import (
	"fmt"
	"go-same-text-to-many/pkg/file"
	"testing"
)

func TestExistsFile(t *testing.T) {
	existsFile := "../../test/exists_file_1.txt"
	notExistsFile := "../../test/not_exists.txt"

	exists := file.ExistsFile(existsFile)
	notExists := file.ExistsFile(notExistsFile)

	expectedExists := true
	expectedNotExists := false

	if exists != expectedExists {
		t.Errorf("File <%s> is not exists, but must exists", existsFile)
	}

	if notExists != expectedNotExists {
		t.Errorf("File <%s> is exists, but must not exists", notExistsFile)
	}
}

func TestExistsFiles(t *testing.T) {
	existFiles := []string{
		"../../test/exists_file_1.txt",
		"../../test/exists_file_2.txt",
	}
	notExistFiles := []string{
		"../../test/not_exists_file_1.txt",
		"../../test/not_exists_file_2.txt",
	}

	existSearchFiles := ListSearchFiles(existFiles)
	notExistSearchFiles := ListSearchFiles(notExistFiles)

	existsError := file.ExistsFiles(existSearchFiles)
	notExistsError := file.ExistsFiles(notExistSearchFiles)

	if existsError != nil {
		t.Errorf(existsError.Error())
	}

	if notExistsError == nil {
		t.Errorf("Files must Exists")
	}
}

func ListSearchFiles(searchFiles []string) []file.SearchFile {
	result := []file.SearchFile{}

	for _, searchFile := range searchFiles {
		result = append(result, file.SearchFile{
			FileName:    searchFile,
			IfNotExists: fmt.Sprintf("File <%s> not exists", searchFile),
		})
	}

	return result
}

func TestReadFile(t *testing.T) {
	fileToRead := "../../test/read_file.txt"
	data, err := file.ReadFile(fileToRead)
	if err != nil {
		t.Errorf("Can't read file <%s>", fileToRead)
		return
	}

	expectedData := "Hello Go Tests"

	if data != expectedData {
		t.Errorf("Result is not equal to expected data.")
	}
}

func TestWriteFile(t *testing.T) {
	fileName := "../../test/write_file.txt"
	content := "Content"

	err := file.WriteFile(fileName, content)
	if err != nil {
		t.Error(err.Error())
		return
	}

	expectedExistsFile := true

	if file.ExistsFile(fileName) != expectedExistsFile {
		t.Errorf("File <%s> not exists", fileName)
		return
	}

	data, err := file.ReadFile(fileName)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if data != content {
		t.Error("File Contains not equals to expected data")
	}

}
