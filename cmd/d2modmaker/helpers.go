package main

import (
	"fmt"

	"github.com/tlentz/d2modmaker/internal/d2file"
	"github.com/tlentz/d2modmaker/internal/superUniquesTxt"
)

func printFile() {
	d2files := d2file.D2Files{}
	f := d2file.GetOrCreateFile(dataDir, d2files, superUniquesTxt.FileName)
	for i := range f.Headers {
		fmt.Println(f.Headers[i], " = ", i)
	}
}
