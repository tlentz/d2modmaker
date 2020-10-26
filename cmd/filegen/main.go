package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

var fs http.FileSystem = http.Dir("../../assets")

func main() {
	err := vfsgen.Generate(fs, vfsgen.Options{})
	if err != nil {
		log.Fatalln(err)
	}
}
