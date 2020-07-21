package main

import (
	"fmt"

	"github.com/tlentz/d2modmaker/internal/d2file"
	"github.com/tlentz/d2modmaker/internal/magicSuffixTxt"
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
		outDir = "./"
		cfgPath = "./cfg.json"
	} else {
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

	// printFile()
	makeMod()
}

func printFile() {
	d2files := d2file.D2Files{}
	f := d2file.GetOrCreateFile(d2files, magicSuffixTxt.FileName)
	for i := range f.Headers {
		fmt.Println(f.Headers[i], " = ", i)
	}
	panic("")
}
