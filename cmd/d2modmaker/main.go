package main

import (
	"github.com/tlentz/d2modmaker/internal/d2file"
	levels "github.com/tlentz/d2modmaker/internal/levelsTxt"
	skills "github.com/tlentz/d2modmaker/internal/skillsTxt"
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

	var d2files = d2file.D2Files{}

	if cfg.EnableTownTeleport {
		enableTownTeleport(&d2files)
	}

	d2file.WriteFiles(&d2files, outDir)
}

func enableTownTeleport(d2files *d2file.D2Files) {
	f := d2file.GetOrCreateFile(dataDir, d2files, skills.FileName)
	for idx, row := range f.Rows {
		if row[skills.Skill] == "Teleport" {
			f.Rows[idx][skills.InTown] = "1"
		}
	}
}

func increaseMonsterDensity(d2files *d2file.D2Files, m int) {
	f := d2file.GetOrCreateFile(dataDir, d2files, levels.FileName)
	maxM := 30
	multiplier := util.Min(maxM, m)
	maxDensity := 10000

	for idx, row := range f.Rows {

	}
	// var monDen = "MonDen"
	// var monUMin = "MonUMin"
	// var monUMax = "MonUMax"
	// var maxDensity = 10000
	// var maxM = 30
	// multiplier := Min(maxM, m)
	// var nmOffset = 1
	// var hellOffset = 2

	// difficulty offsets norm, nm, hell
	// var difficultyOffsets = []int{0, 1, 2}}

	// for idx, row := range f.Rows {
	// 	for _, offset := range offsets {

	// 		// MonDen
	// 		monDenIdx := levels.MonDen + offset
	// 		oldMonDen, err := strconv.Atoi(f.Rows[idx][monDen])
	// 		if err == nil {
	// 			newMonDen := util.Min(maxDensity, oldMonDen*multiplier)
	// 			f.Rows[idx][monDenIdx] = strconv.Itoa(newMonDen)
	// 		}

	// 		// MonUMin

	// 		//
	// 	}
	// }

	// var difficulties = []string{"", "(N)", "(H)"}
	// for idx, _ := range d2file.Records {
	// 	for _, diff := range difficulties {

	// 		// MonDen
	// 		oldMonDen, err := strconv.Atoi(d2file.Records[idx][monDen+diff])
	// 		if err == nil {
	// 			newMonDen := Min(maxDensity, oldMonDen*multiplier)
	// 			d2file.Records[idx][monDen+diff] = strconv.Itoa(newMonDen)
	// 		}

	// 		// MonUMin
	// 		oldMonUMin, err := strconv.Atoi(d2file.Records[idx][monUMin+diff])
	// 		if err == nil {
	// 			newMonUMin := oldMonUMin * multiplier
	// 			d2file.Records[idx][monUMin+diff] = strconv.Itoa(newMonUMin)
	// 		}

	// 		// MonUMax
	// 		oldMonUMax, err := strconv.Atoi(d2file.Records[idx][monUMax+diff])
	// 		if err == nil {
	// 			newMonUMax := oldMonUMax * multiplier
	// 			d2file.Records[idx][monUMax+diff] = strconv.Itoa(newMonUMax)
	// 		}
	// 	}
	// }
}
