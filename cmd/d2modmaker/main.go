package main

import (
	"github.com/tlentz/d2modmaker/internal/d2file"
	"github.com/tlentz/d2modmaker/internal/util"
)

const dataDir = "../../assets/113c-data/"
const outDir = "../../dist/"
const cfgPath = "../../cfg.json"

func main() {
	var cfg = ReadCfg(cfgPath)
	util.PP(cfg)

	d2file, err := d2file.ReadD2File("TreasureClassEx.txt", dataDir)
	util.Check(err)
	util.PP(d2file)
}
