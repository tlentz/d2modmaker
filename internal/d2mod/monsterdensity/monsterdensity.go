package monsterdensity

import (
	"strconv"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/levels"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/monstats"

	"github.com/tlentz/d2modmaker/internal/util"
)

const (
	maxDensity = 10000 // Levels.txt:MonDen* columns
	maxMonU    = 255   // Levels.txt:MonUMin & MonuMax limit
	maxGrp     = 99    // MonStats.txt:MinGrp & MaxGrp limit
)

// Scale Adjust monster density in levels file by scaleFactor
func Scale(d2files d2fs.Files, scaleFactor float64) {
	maxScale := 45.0

	monUMult := util.MinFloat(scaleFactor, maxScale)

	grpMult := 1.0
	maxDensityMult := monUMult
	if monUMult >= 4.5 {
		// As the vanilla highest vanilla value/cap for Grp is 10/99 (99/10 = 9.9) and the highest vanilla value/cap for MonDen is 2200/10000 (10000/2200 = 4.5)
		// this gives a maximum density while still having the same scale/progression of density increase of 9.9*4.5 = 45.
		// Increased slider max to 45, then split the the value back down into 2 separate multipliers, the Grp multiplier and the MonDen multiplier, of which
		// the Grp multiplier caps out at 4.5 and the MonDen multiplier caps out at 10.
		// SuperUniques highest vanilla value/cap is 15/255.  Since (255/15 = 17) 17 is < 45, past 17x the density of superunique packs for the
		// most dense area will stop increasing.  The cap of 99 may be purely artificial, however in 1.06 & 1.07 increasing the density past 99 would cause
		// d2 to crash.
		grpMult = monUMult / 4.5
		maxDensityMult = monUMult / grpMult // 10 * 4.5 is about 45
	}

	//maxDensity := 10000

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
						newVal := util.MinInt(maxMonU, increaseNumByMult(oldVal, monUMult))
						f.Rows[rowIdx][colIdx] = strconv.Itoa(newVal)
					}
				}
			}
		}
	}
	if grpMult > 1 {
		f := d2files.Get(monstats.FileName)
		for rowIdx := range f.Rows {
			// Don't mess with unkillable, Boss, Primeevil, NPC or mobs that are aligned with the player.
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

			colIdx := monstats.MinGrp
			oldVal, err := strconv.Atoi(f.Rows[rowIdx][colIdx])
			if err == nil {
				newVal := util.MinInt(maxGrp, increaseNumByMult(oldVal, grpMult))
				f.Rows[rowIdx][colIdx] = strconv.Itoa(newVal)
			}

			colIdx = monstats.MaxGrp
			oldVal, err = strconv.Atoi(f.Rows[rowIdx][colIdx])
			if err == nil {
				newVal := util.MinInt(maxGrp, increaseNumByMult(oldVal, grpMult))
				f.Rows[rowIdx][colIdx] = strconv.Itoa(newVal)
			}

		}
	}
}
