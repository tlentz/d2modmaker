package main

import (
	"fmt"
	"strconv"

	charStats "github.com/tlentz/d2modmaker/internal/charStatsTxt"
	"github.com/tlentz/d2modmaker/internal/d2file"
	itmRatio "github.com/tlentz/d2modmaker/internal/itemRatioTxt"
	levels "github.com/tlentz/d2modmaker/internal/levelsTxt"
	misc "github.com/tlentz/d2modmaker/internal/miscTxt"
	missiles "github.com/tlentz/d2modmaker/internal/missilesTxt"
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
	makeMod()
	// printFile()
}

func printFile() {
	d2files := d2file.D2Files{}
	f := d2file.GetOrCreateFile(dataDir, &d2files, charStats.FileName)
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
	if cfg.EnableTownSkills {
		enableTownSkills(&d2files)
	}
	if cfg.NoDropZero {
		noDropZero(&d2files)
	}
	if cfg.QuestDrops {
		questDrops(&d2files)
	}
	if cfg.UniqueItemDropRate > 1.0 {
		uniqueItemDropRate(&d2files, cfg.UniqueItemDropRate)
	}
	if cfg.StartWithCube {
		startWithCube(&d2files)
	}
	if cfg.RandomOptions.Randomize {
		Randomize(&cfg, &d2files)
	}

	d2file.WriteFiles(&d2files, outDir)
}

func startWithCube(d2files *d2file.D2Files) {
	f := d2file.GetOrCreateFile(dataDir, d2files, charStats.FileName)
	itemOffset := charStats.Item1
	countOffset := 2
	for idx, row := range f.Rows {
		for i := itemOffset; i < len(row)-countOffset; i += 3 {
			if (row[i] == "0" || row[i] == "") && (row[i+countOffset] == "0" || row[i+countOffset] == "") {
				f.Rows[idx][i] = "box"
				f.Rows[idx][i+countOffset] = "1"
				break // we added a cube, we are done with this row
			}
		}
	}
}

func uniqueItemDropRate(d2files *d2file.D2Files, d float64) {
	f := d2file.GetOrCreateFile(dataDir, d2files, itmRatio.FileName)

	one := func(n int) int {
		if n < 1 {
			return 1
		}
		return n
	}

	divU := func(n int) int {
		return one(int(float64(n) / d))
	}

	divM := func(n int) int {
		return one(int(float64(n) / d * 10))
	}

	for i := range f.Rows {

		// Uniques
		oldUnique, err1 := strconv.Atoi(f.Rows[i][itmRatio.Unique])
		oldUniqueMin, err2 := strconv.Atoi(f.Rows[i][itmRatio.UniqueMin])
		if err1 == nil && err2 == nil {
			newUnique := divU(oldUnique)
			newUniqueMin := divM(oldUniqueMin)
			f.Rows[i][itmRatio.Unique] = strconv.Itoa(newUnique)
			f.Rows[i][itmRatio.UniqueMin] = strconv.Itoa(newUniqueMin)
		}

		// Sets
		oldSet, err3 := strconv.Atoi(f.Rows[i][itmRatio.Set])
		oldSetMin, err4 := strconv.Atoi(f.Rows[i][itmRatio.SetMin])
		if err3 == nil && err4 == nil {
			newSet := divU(oldSet)
			newSetMin := divM(oldSetMin)
			f.Rows[i][itmRatio.Set] = strconv.Itoa(newSet)
			f.Rows[i][itmRatio.SetMin] = strconv.Itoa(newSetMin)
		}
	}
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

func enableTownSkills(d2files *d2file.D2Files) {
	sktxt := d2file.GetOrCreateFile(dataDir, d2files, skills.FileName)
	for i := range sktxt.Rows {
		sktxt.Rows[i][skills.InTown] = "1"
	}
	missilestxt := d2file.GetOrCreateFile(dataDir, d2files, missiles.FileName)
	for i := range missilestxt.Rows {
		missilestxt.Rows[i][missiles.Town] = "1"
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
