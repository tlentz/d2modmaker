package main

import (
	"github.com/tlentz/d2modmaker/internal/d2file"
	"github.com/tlentz/d2modmaker/internal/treasureclassextxt"
	"github.com/tlentz/d2modmaker/internal/util"
)

// dir constants
const (
	dataDir = "../../assets/113c-data/"
	outDir  = "../../dist/"
	cfgPath = "../../cfg.json"
)

func main() {
	var cfg = ReadCfg(cfgPath)
	util.PP(cfg)

	var d2files = map[string]d2file.D2File{}

	d2file.GetOrCreateFile(dataDir, &d2files, treasureclassextxt.FileName)

	d2file.WriteFiles(&d2files, outDir)
}
