package main

import (
	"fmt"

	"github.com/tlentz/d2modmaker/internal/d2mod"
)

var (
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

	d2mod.MakeFromCfgPath(outDir, cfgPath)
}
