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
	CheckD2FileErr(d2file, err)

	defer csvfile.Close()

	// set up reader
	reader := csv.NewReader(csvfile)
	reader.Comma = '\t'

	headers, err := reader.Read()
	CheckD2FileErr(d2file, err)

	rows, err := reader.ReadAll()
	CheckD2FileErr(d2file, err)

	// set the headers/records on D2File
	d2file.Headers = headers
	d2file.Rows = rows

	return d2file, nil
}

func WriteD2File(d2file *D2File, filePath string) {
	file, err := os.Create(filePath + d2file.FileName)
	CheckD2FileErr(d2file, err)
	defer file.Close()

	w := csv.NewWriter(file)
	w.Comma = '\t'
	w.UseCRLF = true
	w.Write(d2file.Headers)
	e := w.WriteAll(d2file.Rows)
	CheckD2FileErr(d2file, e)
}

func WriteFiles(d2files *map[string]D2File, outDir string) {
	fmt.Println("removing " + outDir)
	os.RemoveAll(outDir)

	fmt.Println("creating " + outDir)
	err := os.MkdirAll(outDir, 0755)
	util.Check(err)

	for _, file := range *d2files {
		fmt.Println("writing " + outDir + file.FileName)
		WriteD2File(&file, outDir)
	}
}

func GetOrCreateFile(dataDir string, d2files *map[string]D2File, filename string) *D2File {
	if val, ok := (*d2files)[filename]; ok {
		return &val
	}

	d2file, err := ReadD2File(filename, dataDir)
	util.Check(err)

	(*d2files)[filename] = *d2file

	return d2file
}

func CheckD2FileErr(d2file *D2File, err error) {
	util.CheckError(fmt.Sprintf("Filename: %s", d2file.FileName), err)
}
