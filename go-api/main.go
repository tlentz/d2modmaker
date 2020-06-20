package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type D2File struct {
	FileName string              `json:"filename,omitempty"`
	Headers  []string            `json:"headers,omitempty"`
	Records  []map[string]string `json:"records,omitempty"`
}

func readD2File(fname string) (*D2File, error) {
	// create new D2File pointer with fname
	d2file := &D2File{FileName: fname}

	var filePath = "../assets/test/" + d2file.FileName

	// open csvfile
	csvfile, err := os.Open(filePath)
	if err != nil {
		return nil, encodeError(d2file, err.Error())
	}

	defer csvfile.Close()

	// set up reader
	reader := csv.NewReader(csvfile)
	reader.Comma = '\t'

	rawCsvData, err := reader.ReadAll()
	if err != nil {
		return nil, encodeError(d2file, err.Error())
	}

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

func writeD2File(d2file *D2File) {
	// create file at filePath
	var filePath = "../assets/test/" + d2file.FileName
	file, err := os.Create(filePath)
	checkError("Cannot create file", err)
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
		checkError("error writing csv: ", err)
	}
}

func encodeError(d2file *D2File, msg string) error {
	return fmt.Errorf("Filename: %s, error that occurred: %s", d2file.FileName, msg)
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func main() {
	var fileName = "UniqueItems.txt"
	d2file, err := readD2File(fileName)
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}
	writeD2File(d2file)
}
