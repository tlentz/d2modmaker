package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/tlentz/d2modmaker/internal/d2file"
)

// D2File is a struct holding all the d2 file data
type D2File struct {
	FileName string              `json:"filename,omitempty"`
	Headers  []string            `json:"headers,omitempty"`
	Records  []map[string]string `json:"records,omitempty"`
}

// D2Files is a map[string]D2File
type D2Files = map[string]d2file.D2File

// ReadD2File reads a given d2 file
func ReadD2File(fname string) (*D2File, error) {
	// create new D2File pointer with fname
	d2file := &D2File{FileName: fname}

	var filePath = dataDir + d2file.FileName

	// open csvfile
	csvfile, err := os.Open(filePath)
	CheckD2FileErr(d2file, err)

	defer csvfile.Close()

	// set up reader
	reader := csv.NewReader(csvfile)
	reader.Comma = '\t'

	rawCsvData, err := reader.ReadAll()
	CheckD2FileErr(d2file, err)

	headers := []string{} // holds first row (ordered headers)
	records := []map[string]string{}
	for lineNum, record := range rawCsvData {
		// for first row, build header slice
		if lineNum == 0 {
			for i := 0; i < len(record); i++ {
				headers = append(headers, strings.TrimSpace(record[i]))
			}
		} else {
			// for each cell, map[string] k=header v=value
			line := map[string]string{}
			for i := 0; i < len(record); i++ {
				line[headers[i]] = record[i]
			}
			records = append(records, line)
		}
	}

	// set the headers/records on D2File
	d2file.Headers = headers
	d2file.Records = records

	return d2file, nil
}

// WriteD2File writes the given d2File
func WriteD2File(d2file *D2File) {
	// create file at filePath
	var filePath = outDir + d2file.FileName
	file, err := os.Create(filePath)
	CheckError("Cannot create file", err)
	defer file.Close()

	var records = [][]string{d2file.Headers}

	for _, record := range d2file.Records {
		line := []string{}
		for _, header := range d2file.Headers {
			itm, ok := record[header]
			if ok {
				line = append(line, itm)
			}
		}
		records = append(records, line)
	}

	// create writer
	writer := csv.NewWriter(file)
	writer.Comma = '\t'
	writer.UseCRLF = true

	// write all records (including headers)
	writer.WriteAll(records)

	if err := writer.Error(); err != nil {
		CheckError("error writing csv: ", err)
	}
}

// CheckD2FileErr checks for an error and logs it
func CheckD2FileErr(d2file *D2File, err error) {
	CheckError(fmt.Sprintf("Filename: %s", d2file.FileName), err)
}

// GetOrCreateFile returns the D2File at the given key otherwise creates it
func GetOrCreateFile(d2files *D2Files, filename string) *D2File {
	if val, ok := (*d2files)[filename]; ok {
		return &val
	}

	d2file, err := ReadD2File(filename)
	Check(err)

	(*d2files)[filename] = *d2file

	return d2file
}

// WriteFiles writes all d2 files
func WriteFiles(d2files *D2Files) {

	fmt.Println("removing " + outDir)
	os.RemoveAll(outDir)

	fmt.Println("creating " + outDir)
	err := os.Mkdir(outDir, 0755)
	Check(err)

	for _, file := range *d2files {
		fmt.Println("writing " + outDir + file.FileName)
		WriteD2File(&file)
	}

	fmt.Println("Done")
}
