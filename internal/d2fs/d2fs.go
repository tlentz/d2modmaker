package d2fs

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/tlentz/d2modmaker/internal/d2fs/assets"
	"github.com/tlentz/d2modmaker/internal/util"
)

// File is a struct holding all the d2 file data
type File struct {
	FileName string     `json:"filename,omitempty"`
	Headers  []string   `json:"headers,omitempty"`
	Rows     [][]string `json:"records,omitempty"`
}

// Files is a map[string]D2File
type Files struct {
	cache     map[string]*File
	sourceDir string
	outDir    string
}

func NewFiles(sourceDir string, outDir string) Files {
	files := Files{sourceDir: sourceDir, outDir: outDir}
	files.cache = make(map[string]*File)

	//os.RemoveAll(path.Join(files.outDir, "/data/"))	// obc:  This is c4 approach, use sword instead
	err := os.MkdirAll(path.Join(files.outDir, assets.DataGlobalExcel), 0755)
	util.Check(err)

	removefilenames, err := filepath.Glob(path.Join(files.outDir, assets.DataGlobalExcel+"*.txt"))
	util.Check(err)
	removefilenames2, err := filepath.Glob(path.Join(files.outDir, assets.PatchStringDest+"patchstring.tbl"))
	util.Check(err)
	removefilenames = append(removefilenames, removefilenames2...)

	for _, f := range removefilenames {
		if err := os.Remove(f); err != nil {
			util.Check(err)
		}
	}

	return files
}

// ReadD2File reads a given d2 file
func (d2files *Files) Read(filename string) *File {
	if d2files.sourceDir == "" {
		// open csvfile
		csvfile, err := assets.Assets.Open(path.Join(assets.DataDir, filename))
		//TODO: figure out how to move this out of if/else
		checkError(filename, err)
		defer csvfile.Close()
		return importCsv(csvfile, filename)
	} else {
		csvfile, err := os.Open(path.Join(d2files.sourceDir, filename))
		checkError(filename, err)
		defer csvfile.Close()
		return importCsv(csvfile, filename)
	}
}

func ReadAsset(filename string, filePath string) *File {
	// open csvfile
	csvfile, err := assets.Assets.Open(path.Join(filePath, filename))
	checkError(filename, err)
	defer csvfile.Close()

	return importCsv(csvfile, filename)
}

func importCsv(csvfile io.Reader, filename string) *File {
	// create new D2File pointer with fname
	d2file := &File{FileName: filename}

	// set up reader
	reader := csv.NewReader(csvfile)
	reader.Comma = '\t'

	raw, err := reader.ReadAll()
	checkError(filename, err)

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
	checkError(filename, err)

	// set the headers/records on D2File
	d2file.Headers = headers
	d2file.Rows = rows
	return d2file
}

// WriteFiles writes all d2 files
func (d2files *Files) Write() {
	for _, d2file := range d2files.cache {
		// fmt.Println("writing " + path.Join(d2files.outDir, d2file.FileName))
		d2file.write(d2files.outDir)
	}
}

func (d2file *File) write(outDir string) {
	file, err := os.Create(path.Join(outDir, assets.DataGlobalExcel, d2file.FileName))
	checkError(d2file.FileName, err)
	defer file.Close()

	w := csv.NewWriter(file)
	w.Comma = '\t'
	w.UseCRLF = true
	w.Write(d2file.Headers)
	e := w.WriteAll(d2file.Rows)
	checkError(d2file.FileName, e)
}

// GetOrCreateFile returns the D2File at the given key otherwise creates it
func (d2files *Files) Get(filename string) *File {
	if val, ok := d2files.cache[filename]; ok {
		return val
	}

	file := d2files.Read(filename)

	d2files.cache[filename] = file

	return file
}

// checkError checks for an error and logs it
func checkError(filename string, err error) {
	util.CheckError(fmt.Sprintf("Filename: %s", filename), err)
}

func MergeRows(f1 *File, f2 File) {
	f1.Rows = append(f1.Rows, f2.Rows...)
}

//func printFile() {
//	d2files := d2file.D2Files{}
//	f := d2file.GetOrCreateFile(d2files, magicSuffix.FileName)
//	for i := range f.Headers {
//		fmt.Println(f.Headers[i], " = ", i)
//	}
//	panic("")
//}
func DebugDumpFiles(f Files, filename string) {
	fmt.Printf("%s\n", f.cache[filename])
}
