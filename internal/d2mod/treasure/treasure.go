package treasure

import (
	"math"
	"strconv"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/itemRatio"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/treasureClassEx"
)

func SetNoDropZero(d2files d2fs.Files) {
	f := d2files.Get(treasureClassEx.FileName)
	for idx, row := range f.Rows {
		if row[treasureClassEx.NoDrop] != "" {
			f.Rows[idx][treasureClassEx.NoDrop] = "0"
		}
	}
}

func EnableQuestDrops(d2files d2fs.Files) {
	f := d2files.Get(treasureClassEx.FileName)
	diffOffsets := []int{0, 1, 2} // norm, nm, hell
	bossQOffset := 3
	for idx, row := range f.Rows {
		switch row[treasureClassEx.TreasureClass] {
		case treasureClassEx.Andariel, treasureClassEx.Duriel, treasureClassEx.DurielBase, treasureClassEx.Mephisto, treasureClassEx.Diablo, treasureClassEx.Baal:
			{
				for _, offset := range diffOffsets {
					tmp := make([]string, len(row))
					// copy quest drop row to tmp
					copy(tmp, f.Rows[idx+bossQOffset+offset])
					// copy all tmp values except 1st index to original row
					copy(f.Rows[idx+offset][treasureClassEx.TreasureClass+1:], tmp[treasureClassEx.TreasureClass+1:])
				}
			}
		}
	}
}

func ScaleUniqueDropRate(d2files d2fs.Files, d float64) {
	f := d2files.Get(itemRatio.FileName)

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
		dM := 1 + (d-1)/10
		return one(int(float64(n) / dM))
	}

	for i := range f.Rows {

		// Uniques
		oldUnique, err1 := strconv.Atoi(f.Rows[i][itemRatio.Unique])
		oldUniqueMin, err2 := strconv.Atoi(f.Rows[i][itemRatio.UniqueMin])
		if err1 == nil && err2 == nil {
			newUnique := divU(oldUnique)
			newUniqueMin := divM(oldUniqueMin)
			f.Rows[i][itemRatio.Unique] = strconv.Itoa(newUnique)
			f.Rows[i][itemRatio.UniqueMin] = strconv.Itoa(newUniqueMin)
		}

		// Sets
		oldSet, err3 := strconv.Atoi(f.Rows[i][itemRatio.Set])
		oldSetMin, err4 := strconv.Atoi(f.Rows[i][itemRatio.SetMin])
		if err3 == nil && err4 == nil {
			newSet := divU(oldSet)
			newSetMin := divM(oldSetMin)
			f.Rows[i][itemRatio.Set] = strconv.Itoa(newSet)
			f.Rows[i][itemRatio.SetMin] = strconv.Itoa(newSetMin)
		}
	}
}

func ScaleRuneDropRate(d2files d2fs.Files, rateScale float64) {
	f := d2files.Get(treasureClassEx.FileName)

	// Clip rateScale to valid 1-100 range
	if rateScale > 100.0 {
		rateScale = 100.0
	}
	if rateScale < 1.0 {
		rateScale = 1.0
	}
	// Convert rateScale to 0.0-1.0 range
	rateScale = (rateScale - 1.0) / 99.0

	// Invert rate scale to get the drop rate multiplier
	rateMult := 1.0 - rateScale

	targetProb12 := 1.0

	origProb1 := 3.0
	baseProb1 := math.Log2(targetProb12)
	rangeProb1 := math.Log2(origProb1) - math.Log2(targetProb12)
	newProb1 := math.Pow(2, baseProb1+rangeProb1*rateMult)

	origProb2 := 2.0
	baseProb2 := math.Log2(targetProb12)
	rangeProb2 := math.Log2(origProb2) - math.Log2(targetProb12)
	newProb2 := math.Pow(2, baseProb2+rangeProb2*rateMult)

	for idx, row := range f.Rows {
		treasureClass := row[treasureClassEx.TreasureClass]
		if len(treasureClass) >= 5 && treasureClass[:5] == "Runes" {
			runeTc, err := strconv.Atoi(row[treasureClassEx.TreasureClass][6:])

			if err == nil {
				targetProb3 := float64((runeTc - 1) * 2)
				var newProb3 float64 = 0
				if targetProb3 != 0 {
					var origProb3 int
					if runeTc != 17 {
						origProb3, _ = strconv.Atoi(row[treasureClassEx.Prob3])
					} else {
						// Runes 17 doesn't have a second rune chance, and the next rune TC slot is in Prob2
						origProb3, _ = strconv.Atoi(row[treasureClassEx.Prob2])
					}
					baseProb3 := math.Log2(targetProb3)
					rangeProb3 := math.Log2(float64(origProb3)) - math.Log2(targetProb3)
					newProb3 = math.Pow(2, baseProb3+rangeProb3*rateMult)
				}

				if runeTc != 17 {
					f.Rows[idx][treasureClassEx.Prob1] = strconv.Itoa(int(newProb1 + 0.5))
					f.Rows[idx][treasureClassEx.Prob2] = strconv.Itoa(int(newProb2 + 0.5))
					if runeTc != 1 {
						// Runes 1 does not have a next rune TC
						f.Rows[idx][treasureClassEx.Prob3] = strconv.Itoa(int(newProb3 + 0.5))
					}
				} else {
					// Runes 17 doesn't have a second rune chance, and the next rune TC slot is in Prob2
					f.Rows[idx][treasureClassEx.Prob2] = strconv.Itoa(int(newProb3 + 0.5))
				}
			}
		}
	}
}
