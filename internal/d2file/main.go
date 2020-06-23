package d2file

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/tlentz/d2modmaker/internal/util"
)

type D2File struct {
	FileName string     `json:"filename,omitempty"`
	Headers  []string   `json:"headers,omitempty"`
	Rows     [][]string `json:"records,omitempty"`
}

func ReadD2File(fname string, filePath string) (*D2File, error) {
	// create new D2File pointer with fname
	d2file := &D2File{FileName: fname}

	// open csvfile
	csvfile, err := os.Open(filePath + fname)
	checkD2FileErr(d2file, err)

	defer csvfile.Close()

	// set up reader
	reader := csv.NewReader(csvfile)
	reader.Comma = '\t'

	headers, err := reader.Read()
	checkD2FileErr(d2file, err)

	rows, err := reader.ReadAll()
	checkD2FileErr(d2file, err)

	// set the headers/records on D2File
	d2file.Headers = headers
	d2file.Rows = rows

	return d2file, nil
}

func WriteD2File2(d2file *D2File, filePath string) {
	file, err := os.Create(filePath)
	checkD2FileErr(d2file, err)
	defer file.Close()

	w := csv.NewWriter(file)
	w.Comma = '\t'
	w.UseCRLF = true
	w.Write(d2file.Headers)
	e := w.WriteAll(d2file.Rows)
	checkD2FileErr(d2file, e)
}

func checkD2FileErr(d2file *D2File, err error) {
	util.CheckError(fmt.Sprintf("Filename: %s", d2file.FileName), err)
}
