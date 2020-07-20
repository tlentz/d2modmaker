package main

import (
	"fmt"
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
