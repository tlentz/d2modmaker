package main

import (
	"github.com/tlentz/d2modmaker/internal/d2file"
	skillstxt "github.com/tlentz/d2modmaker/internal/skillsTxt"
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

	var d2files = D2Files{}

	if cfg.EnableTownTeleport {
		enableTownTeleport(d2files)
	}

	d2file.WriteFiles(&d2files, outDir)
}

func enableTownTeleport(d2files *d2file.D2Files) {
	f := d2file.GetOrCreateFile(dataDir, &d2file, skillstxt.FileName)
	for idx, row := range f.Records {
		if row[skillstxt.Skill] == "Teleport" {
			f.Records[idx][skillstxt.InTown] = "1"
		}
	}
}
