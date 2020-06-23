package main

import (
	"strconv"

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

	if cfg.IncreaseMonsterDensity > 1.0 {
		increaseMonsterDensity(&d2files, cfg.IncreaseMonsterDensity)
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

func increaseMonsterDensity(d2files *d2file.D2Files, m float64) {
	f := d2file.GetOrCreateFile(dataDir, d2files, levels.FileName)
	maxM := 30.0
	mult := util.MinFloat(m, maxM)
	maxDensity := 10000

	increaseNumByMult := func(n int) int {
		return int(float64(n) * mult)
	}

	diffOffsets := []int{0, 1, 2} // norm, nm, hell
	minMaxOffsets := []int{0, 1}  // min, max

	for rowIdx := range f.Rows {

		for _, diffOffset := range diffOffsets {

			// MonDen
			colIdx := levels.MonDen + diffOffset
			oldVal, err := strconv.Atoi(f.Rows[rowIdx][colIdx])
			if err == nil {
				newVal := util.MinInt(maxDensity, increaseNumByMult(oldVal))
				f.Rows[rowIdx][colIdx] = strconv.Itoa(newVal)
			}

			// MonUMin / MonUMax
			for _, minMaxOffset := range minMaxOffsets {
				colIdx := levels.MonUMin + minMaxOffset + diffOffset
				oldVal, err := strconv.Atoi(f.Rows[rowIdx][colIdx])
				if err == nil {
					newVal := increaseNumByMult(oldVal)
					f.Rows[rowIdx][colIdx] = strconv.Itoa(newVal)
				}
			}
		}
	}
}
