package threelevelbreak

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestImportFile(t *testing.T) {
	_, err := loadFile("input.csv")

	if err != nil {
		t.Error(err)
	}
}

func TestReport(t *testing.T) {
	input := "input.csv"
	output := "output.txt"

	err := generateReport(input, output)

	resFile, err := ioutil.ReadFile("output.txt")
	testFile, err := ioutil.ReadFile("test.txt")

	if bytes.Compare(resFile, testFile) != 0 {
		t.Error("Files don't match! Lame!")
	}

	if err != nil {
		t.Error(err)
	}
}
