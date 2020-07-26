package monsterdensity

import (
	"strconv"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/levels"

	"github.com/tlentz/d2modmaker/internal/util"
)

func Scale(d2files d2fs.Files, scaleFactor float64) {
	f := d2files.Get(levels.FileName)
	maxScale := 30.0
	mult := util.MinFloat(scaleFactor, maxScale)
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
