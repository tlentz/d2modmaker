package monsterdensity

import (
	"strconv"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/levels"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/monstats"

	"github.com/tlentz/d2modmaker/internal/util"
)

// Scale Adjust monster density in levels file by scaleFactor
func Scale(d2files d2fs.Files, scaleFactor float64) {
	maxScale := 45.0

	monUMult := util.MinFloat(scaleFactor, maxScale)

	grpMult := 1.0
	maxDensityMult := monUMult
	if monUMult >= 4.5 {
		maxDensityMult = 4.5
		grpMult = monUMult / 4.5
		maxDensityMult = monUMult / grpMult // 6.6 * 4.5 is about 30
	}
	maxDensity := 10000

	increaseNumByMult := func(n int, m float64) int {
		return int(float64(n) * m)
	}

	diffOffsets := []int{0, 1, 2} // norm, nm, hell
	minMaxOffsets := []int{0, 1}  // min, max
	{
		f := d2files.Get(levels.FileName)
		for rowIdx := range f.Rows {

			for _, diffOffset := range diffOffsets {

				// MonDen
				colIdx := levels.MonDen + diffOffset
				oldVal, err := strconv.Atoi(f.Rows[rowIdx][colIdx])
				if err == nil {
					newVal := util.MinInt(maxDensity, increaseNumByMult(oldVal, maxDensityMult))
					f.Rows[rowIdx][colIdx] = strconv.Itoa(newVal)
				}

				// MonUMin / MonUMax
				for _, minMaxOffset := range minMaxOffsets {
					colIdx := levels.MonUMin + minMaxOffset + diffOffset*2
					oldVal, err := strconv.Atoi(f.Rows[rowIdx][colIdx])
					if err == nil {
						newVal := util.MinInt(255, increaseNumByMult(oldVal, monUMult))
						f.Rows[rowIdx][colIdx] = strconv.Itoa(newVal)
					}
				}
			}
		}
	}
	if grpMult > 1 {
		f := d2files.Get(monstats.FileName)
		for rowIdx := range f.Rows {
			if f.Rows[rowIdx][monstats.Killable] != "1" {
				continue
			}
			if f.Rows[rowIdx][monstats.Boss] == "1" {
				continue
			}
			if f.Rows[rowIdx][monstats.Primeevil] == "1" {
				continue
			}
			if f.Rows[rowIdx][monstats.Npc] == "1" {
				continue
			}
			if f.Rows[rowIdx][monstats.Align] != "" {
				continue
			}
			{
				colIdx := monstats.MinGrp
				oldVal, err := strconv.Atoi(f.Rows[rowIdx][colIdx])
				if err == nil {
					newVal := util.MinInt(99, increaseNumByMult(oldVal, grpMult))
					f.Rows[rowIdx][colIdx] = strconv.Itoa(newVal)
				}
			}
			{
				colIdx := monstats.MaxGrp
				oldVal, err := strconv.Atoi(f.Rows[rowIdx][colIdx])
				if err == nil {
					newVal := util.MinInt(99, increaseNumByMult(oldVal, grpMult))
					f.Rows[rowIdx][colIdx] = strconv.Itoa(newVal)
				}
			}
		}
	}
}
