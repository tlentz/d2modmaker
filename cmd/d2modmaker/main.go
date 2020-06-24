package main

import (
	"fmt"
	"strconv"

	"github.com/tlentz/d2modmaker/internal/d2file"
	levels "github.com/tlentz/d2modmaker/internal/levelsTxt"
	misc "github.com/tlentz/d2modmaker/internal/miscTxt"
	skills "github.com/tlentz/d2modmaker/internal/skillsTxt"
	tc "github.com/tlentz/d2modmaker/internal/treasureclassextxt"
	"github.com/tlentz/d2modmaker/internal/util"
)

// dir constants
const (
	dataDir = "../../assets/113c-data/"
	outDir  = "../../dist/"
	cfgPath = "../../cfg.json"
)

func main() {
	d2files := d2file.D2Files{}
	f := d2file.GetOrCreateFile(dataDir, &d2files, "ItemRatio.txt")
	for i := range f.Headers {
		fmt.Println(f.Headers[i], " = ", i)
	}
}

func makeMod() {
	var cfg = ReadCfg(cfgPath)
	util.PP(cfg)

	var d2files = d2file.D2Files{}

	if cfg.IncreaseStackSizes {
		increaseStackSizes(&d2files)
	}
	if cfg.IncreaseMonsterDensity > 1.0 {
		increaseMonsterDensity(&d2files, cfg.IncreaseMonsterDensity)
	}
	if cfg.EnableTownTeleport {
		enableTownTeleport(&d2files)
	}
	if cfg.NoDropZero {
		noDropZero(&d2files)
	}
	if cfg.QuestDrops {
		questDrops(&d2files)
	}

	d2file.WriteFiles(&d2files, outDir)
}

func questDrops(d2files *d2file.D2Files) {
	f := d2file.GetOrCreateFile(dataDir, d2files, tc.FileName)
	diffOffsets := []int{0, 1, 2} // norm, nm, hell
	bossQOffset := 3
	for idx, row := range f.Rows {
		switch row[tc.TreasureClass] {
		case tc.Andariel, tc.Duriel, tc.DurielBase, tc.Mephisto, tc.Diablo, tc.Baal:
			{
				for _, offset := range diffOffsets {
					tmp := make([]string, cap(row))
					// copy quest drop row to tmp
					copy(tmp, f.Rows[idx+bossQOffset+offset])
					// copy all tmp values except 1st index to original row
					copy(f.Rows[idx+offset][tc.TreasureClass+1:], tmp[tc.TreasureClass+1:])
				}
			}
		}
	}
}

func noDropZero(d2files *d2file.D2Files) {
	f := d2file.GetOrCreateFile(dataDir, d2files, tc.FileName)
	for idx, row := range f.Rows {
		if row[tc.NoDrop] != "" {
			f.Rows[idx][tc.NoDrop] = "0"
		}
	}
}

func increaseStackSizes(d2files *d2file.D2Files) {
	f := d2file.GetOrCreateFile(dataDir, d2files, misc.FileName)
	for idx, row := range f.Rows {
		if row[misc.Name] == misc.TownPortalBook || row[misc.Name] == misc.IdentifyBook {
			f.Rows[idx][misc.MaxStack] = "100"
		}
	}
}

func enableTownTeleport(d2files *d2file.D2Files) {
	f := d2file.GetOrCreateFile(dataDir, d2files, skills.FileName)
	for idx, row := range f.Rows {
		if row[skills.Skill] == skills.Teleport {
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
