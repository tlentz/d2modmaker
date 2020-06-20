package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type D2File struct {
	FileName string              `json:"filename,omitempty"`
	Headers  []string            `json:"headers,omitempty"`
	Records  []map[string]string `json:"records,omitempty"`
}

const dataDir = "../assets/113c-data/"
const outDir = "./dist/"

func readD2File(fname string) (*D2File, error) {
	// create new D2File pointer with fname
	d2file := &D2File{FileName: fname}

	var filePath = dataDir + d2file.FileName

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
	var filePath = outDir + d2file.FileName
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

////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////

func getItemFromRecords(d2file *D2File, key string, name string) (*int, error) {
	for i, record := range d2file.Records {
		itm, ok := record[key]
		if ok && itm == name {
			return &i, nil
		}
	}
	return nil, fmt.Errorf("Cannot find %s : %s", key, name)
}

func pp(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

func increaseStackSizes(d2file *D2File) (*D2File, error) {

	// Sets TP Book size to 100
	tpBookIdx, err := getItemFromRecords(d2file, "name", "Town Portal Book")
	if err != nil {
		encodeError(d2file, err.Error())
	}
	d2file.Records[*tpBookIdx]["maxstack"] = "100"

	// Sets Id Book size to 100
	idBookIdx, err := getItemFromRecords(d2file, "name", "Identify Book")
	if err != nil {
		encodeError(d2file, err.Error())
	}
	d2file.Records[*idBookIdx]["maxstack"] = "100"

	return d2file, nil
}

func addFileIfNotExists(d2files map[string]D2File, filename string) map[string]D2File {
	if _, ok := d2files[filename]; ok {
		return d2files
	}

	pp("hi2")
	d2file, err := readD2File(filename)
	if err != nil {
		pe(err)
	}

	d2files[filename] = *d2file
	return d2files
}

func getOrCreateFile(d2files *map[string]D2File, filename string) *D2File {
	if val, ok := (*d2files)[filename]; ok {
		return &val
	}

	d2file, err := readD2File(filename)
	if err != nil {
		pe(err)
	}

	(*d2files)[filename] = *d2file

	return d2file
}

type ModConfig struct {
	IncreaseStackSizes bool `json:"IncreaseStackSizes"`
}

func pe(e error) {
	panic(fmt.Sprintf("An error encountered :: ", e))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readCfg() ModConfig {
	file, _ := ioutil.ReadFile("cfg.json")
	data := ModConfig{}
	_ = json.Unmarshal([]byte(file), &data)
	return data
}

func writeFiles(d2files *map[string]D2File) {

	fmt.Println("removing " + outDir)
	os.RemoveAll(outDir)

	fmt.Println("creating " + outDir)
	err := os.Mkdir(outDir, 0755)
	check(err)

	for _, file := range *d2files {
		fmt.Println("writing " + outDir + file.FileName)
		writeD2File(&file)
	}

	fmt.Println("Done")
}

func main() {
	var cfg = readCfg()
	var d2files = map[string]D2File{}

	d2file := getOrCreateFile(&d2files, "Misc.txt")
	if cfg.IncreaseStackSizes {
		increaseStackSizes(d2file)
	}

	writeFiles(&d2files)

	// pp(d2files["Misc.txt"].Records[10]["maxstack"])
}
