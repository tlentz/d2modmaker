package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

var fs http.FileSystem = http.Dir("../../assets")

func main() {
	var vfsoptions = vfsgen.Options{
		Filename:     "../../internal/d2fs/assets/assets_vfsdata.go",
		PackageName:  "assets",
		VariableName: "Assets",
	}
	err := vfsgen.Generate(fs, vfsoptions)
	if err != nil {
		log.Fatalln(err)
	}
}
