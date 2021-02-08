package d2fs

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/tlentz/d2modmaker/internal/d2fs/assets"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/propscorestxt"
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

// FileInfo 1 declared in each file in d2fs\txts\
type FileInfo struct {
	FileName   string
	FileNumber int // from package filenumbers
	NumColumns int
}

// ItemFileInfo 1 declared in each file in d2fs\txtx\ that contains Items
type ItemFileInfo struct {
	FI               FileInfo
	ItemName         int // Was going to call this Name but got confused about whether is Filename or Item name
	Code             int // Column Index for Item Code
	Lvl              int // Column Index for Item Level
	FirstProp        int // Column Index for first Prop
	NumProps         int
	HasEnabledColumn bool
}

// NewFiles Create a new Files from configured directories
func NewFiles(sourceDir string, outDir string) Files {

	if (sourceDir == outDir) && (sourceDir != "") {
		log.Fatalf("Error: Source Directory == Output Directory.  Either set the Source Directory to blank (to use the assets/ dir), or set it to a mods data/ directory.  OutDir is deleted and re-created each run")
	}
	// If the Source Directory is not specified, point at the assets\113c-data\ directory
	if sourceDir == "" {
		sourceDir = path.Join(assets.AssetDir, assets.DataDir)
	}

	files := Files{sourceDir: sourceDir, outDir: outDir}
	files.cache = make(map[string]*File)
	//os.RemoveAll(path.Join(files.outDir, "/data/"))	// obc:  This is the old c4 approach, use sword instead
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

// Read reads a given d2 file from the specified Directory
func (d2files *Files) Read(filepath string, filename string) *File {
	filePathName := path.Join(filepath, filename)
	csvfile, err := os.Open(filePathName)
	checkError(filename, err)
	defer csvfile.Close()
	return importCsv(csvfile, filename)
}

// ReadAsset Reads in a tsv file from the assets/ directory but doesn't cache it in Files
func ReadAsset(filePath string, filename string) *File {
	filePathName := path.Join(assets.AssetDir, filePath, filename)
	csvfile, err := os.Open(filePathName)
	checkError(filePathName, err)
	defer csvfile.Close()
	return importCsv(csvfile, filename)
}

// List list all filenames in a directory
func (d2files *Files) List(filePath string) []os.FileInfo {
	//filenames := make([]string, 0)
	// []os.FileInfo
	var err error
	absDir, err := filepath.Abs(filePath)
	util.CheckError(filePath, err)

	//fmt.Printf("d2fs.List: Listing files in %s\n", absDir)
	dir, err := os.Open(filePath)
	util.CheckError(filePath, err)
	finfo, err := dir.Readdir(-1)
	util.CheckError(filePath, err)
	fmt.Printf("Found %d files in %s\n", len(finfo), absDir)
	return finfo
}

// importCsv This is actually tab separated value, i.e. tsv
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
		//fmt.Println("writing " + path.Join(d2files.outDir, d2file.FileName))
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

// Get returns the D2File at the given key otherwise creates it.
// Using assets.DataDir as path
func (d2files *Files) Get(filename string) *File {
	if val, ok := d2files.cache[filename]; ok {
		return val
	}
	dataDir := path.Join(d2files.sourceDir, assets.GlobalExcelDir)
	if d2files.sourceDir == "" {
		dataDir = path.Join(assets.AssetDir, assets.DataDir, assets.GlobalExcelDir)
	}
	file := d2files.Read(dataDir, filename)

	d2files.cache[filename] = file

	return file
}

// GetAsset read a tab delimited text file from assets/pathname & filename
// pathname is assumed to be relative to the assets/ dir.
// The Get routines add the filename to the cache, which is
// written out by Write.
func (d2files *Files) GetAsset(pathname string, filename string) *File {
	if val, ok := d2files.cache[filename]; ok {
		return val
	}
	file := d2files.Read(path.Join(assets.AssetDir, pathname), filename)

	d2files.cache[filename] = file

	return file
}

// checkError checks for an error and logs it
func checkError(filename string, err error) {
	util.CheckError(fmt.Sprintf("Filename: %s", filename), err)
}

// AppendRows concatenate all rows from f2 into f1
func AppendRows(f1 *File, f2 File) {
	f1.Rows = append(f1.Rows, f2.Rows...)
}

// MergeRows Merges f2 into f1
// f1.Rows[][0] is assumed to be unique
func MergeRows(f1 *File, f2 File) {
	keys := make(map[string]int, 0)
	for rowIdx := range f1.Rows {

		keys[genUniqueKey(*f1, rowIdx)] = rowIdx
	}
	for f2RowIdx := range f2.Rows {
		f1RowIdx, ok := keys[genUniqueKey(f2, f2RowIdx)]
		if ok {
			f1.Rows[f1RowIdx] = f2.Rows[f2RowIdx]
			//fmt.Printf("Merging...%s\n", f2.Rows[f2RowIdx][0])
		} else {
			f1.Rows = append(f1.Rows, f2.Rows[f2RowIdx])
		}
	}
}
func genUniqueKey(f File, rowIndex int) string {
	switch {
	case f.FileName == "PropScores.txt":
		return f.Rows[rowIndex][propscorestxt.Prop] + "/" + f.Rows[rowIndex][propscorestxt.Par] + "/" + f.Rows[rowIndex][propscorestxt.Min] + "/" + f.Rows[rowIndex][propscorestxt.Max]
	default:
		return f.Rows[rowIndex][0] // Default to first column is ID
	}
}

// This function will generate a list of constants from a header.
// Useful when creating a new internal\txt\blah\blah.go file
//func printFileHeader() {
//	d2files := d2file.D2Files{}
//	f := d2file.GetOrCreateFile(d2files, magicSuffix.FileName)
//	for i := range f.Headers {
//		fmt.Println(f.Headers[i], " = ", i)
//	}
//	panic("")
//}

// DebugDumpVFSFileNames Dump all of Files to console
func DebugDumpVFSFileNames(f Files, filename string) {
	fmt.Printf("%s\n", f.cache[filename])
}
