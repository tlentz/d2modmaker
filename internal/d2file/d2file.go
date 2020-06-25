package d2file

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/tlentz/d2modmaker/internal/util"
)

// D2File is a struct holding all the d2 file data
type D2File struct {
	FileName string     `json:"filename,omitempty"`
	Headers  []string   `json:"headers,omitempty"`
	Rows     [][]string `json:"records,omitempty"`
}

// D2Files is a map[string]D2File
type D2Files = map[string]D2File

// ReadD2File reads a given d2 file
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

	raw, err := reader.ReadAll()
	CheckD2FileErr(d2file, err)

	rows := make([][]string, 0)
	headers := make([]string, 0)

	for i, line := range raw {
		if i == 0 {
			for j := range line {
				headers = append(headers, line[j])
			}
		} else {
			rows = append(rows, line)
		}
	}
	CheckD2FileErr(d2file, err)

	// set the headers/records on D2File
	d2file.Headers = headers
	d2file.Rows = rows

	return d2file, nil
}

// WriteD2File writes the given d2File
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

// WriteFiles writes all d2 files
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

// GetOrCreateFile returns the D2File at the given key otherwise creates it
func GetOrCreateFile(dataDir string, d2files *map[string]D2File, filename string) *D2File {
	if val, ok := (*d2files)[filename]; ok {
		return &val
	}

	d2file, err := ReadD2File(filename, dataDir)
	util.Check(err)

	(*d2files)[filename] = *d2file

	return d2file
}

// CheckD2FileErr checks for an error and logs it
func CheckD2FileErr(d2file *D2File, err error) {
	util.CheckError(fmt.Sprintf("Filename: %s", d2file.FileName), err)
}
