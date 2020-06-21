package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type D2File struct {
	FileName string              `json:"filename,omitempty"`
	Headers  []string            `json:"headers,omitempty"`
	Records  []map[string]string `json:"records,omitempty"`
}

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

func GetItemFromRecords(d2file *D2File, key string, name string) (*int, error) {
	for i, record := range d2file.Records {
		itm, ok := record[key]
		if ok && itm == name {
			return &i, nil
		}
	}
	return nil, fmt.Errorf("Cannot find %s : %s", key, name)
}

func CheckD2FileErr(d2file *D2File, err error) {
	CheckError(fmt.Sprintf("Filename: %s", d2file.FileName), err)
}

func AddFileIfNotExists(d2files map[string]D2File, filename string) map[string]D2File {
	if _, ok := d2files[filename]; ok {
		return d2files
	}

	d2file, err := ReadD2File(filename)
	Check(err)

	d2files[filename] = *d2file
	return d2files
}

func GetOrCreateFile(d2files *map[string]D2File, filename string) *D2File {
	if val, ok := (*d2files)[filename]; ok {
		return &val
	}

	d2file, err := ReadD2File(filename)
	Check(err)

	(*d2files)[filename] = *d2file

	return d2file
}

func WriteFiles(d2files *map[string]D2File) {

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
