package main

import (
	"fmt"

	"github.com/tlentz/d2modmaker/internal/d2file"
	"github.com/tlentz/d2modmaker/internal/superUniquesTxt"
)

var (
	dataDir string
	outDir  string
	cfgPath string
	mode    string
	version string
)

func main() {

	if mode == "production" {
		dataDir = "./113c-data/"
		outDir = "./data/global/excel/"
		cfgPath = "./cfg.json"
	} else {
		dataDir = "../../assets/113c-data/"
		outDir = "../../dist/"
		cfgPath = "../../cfg.json"
	}

	if version == "" {
		version = "[Dev Build]"
	}
	line := "==============================="
	fmt.Println(line)
	fmt.Println("", "D2 Mod Maker", version)
	fmt.Println(line)

	makeMod()
}

func printFile() {
	d2files := d2file.D2Files{}
	f := d2file.GetOrCreateFile(dataDir, d2files, superUniquesTxt.FileName)
	for i := range f.Headers {
		fmt.Println(f.Headers[i], " = ", i)
	}
}
