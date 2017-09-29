package datalaundry

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestCleanFile(t *testing.T) {
	inputFileName := "input.txt"
	resultFileName := "result.csv"

	err := cleanFile(inputFileName, resultFileName)
	if err != nil {
		t.Error(err)
	}

	resultData, rerr := ioutil.ReadFile(resultFileName)
	if rerr != nil {
		t.Error(rerr)
	}

	testData, terr := ioutil.ReadFile("testResult.csv")
	if terr != nil {
		t.Error(terr)
	}

	if !bytes.Equal(resultData, testData) {
		t.Errorf("Files do not match! Awful!")
	}
}
func TestNewIface(t *testing.T) {
	iname := "name"
	inet := "inet"
	status := "up"

	test := iface{name: iname, inet: inet, status: status}

	if test.name != iname {
		t.Errorf("Expect %v, got %v", test.name, iname)
	}
	if test.inet != inet {
		t.Errorf("Expect %v, got %v", inet, test.inet)
	}
	if test.status != status {
		t.Errorf("Expect %v, got %v", status, test.status)
	}
}

func TestReadFile(t *testing.T) {
	// Test data
	tData := &iface{name: "lo0", inet: "127.0.0.1"}

	myFile := "/path/to/file.txt"
	myData := &iface{}

	err := readFile(myFile, myData)

	if err != nil {
		t.Error(err)
	}
	if !match(myData, tData) {
		t.Errorf("Expected %v, got %v", tData, myData)
	}
}

func match(ia *iface, ib *iface) bool {
	if ia.name != ib.name {
		return false
	}
	if ia.inet != ib.inet {
		return false
	}
	if ia.status != ib.status {
		return false
	}
	return true
}
