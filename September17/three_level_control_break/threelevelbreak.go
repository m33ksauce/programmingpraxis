package threelevelbreak

import (
	"io/ioutil"
	"os"
)

type inputFile struct {
	pos int
	lines []byte
}

func InputFile(name string) (i *inputFile) {

}

func loadFile(fname string) ([]byte, error) {
	return nil, nil
}

func generateReport(in string, out string) error {
	inFile := os.Open(in)

	inFile.

	err := ioutil.WriteFile(out, []byte("1"), os.ModeAppend)
	return err
}
