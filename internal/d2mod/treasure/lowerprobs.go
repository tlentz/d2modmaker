package treasure

import (
	"strconv"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/treasureClassEx"
)

// lowerProbs Lowers the probability of scaleColIdx being rolled by lowerByFactor for row rowIndex in file f
// Does this by increasing the Prob(n) columns of every row except ScaleColIdx and increasing
// the NoDrop column additionally by (lowerByFactor-1)
// Used by biggergoldpiles to make fewer gold piles.
func lowerProbs(f *d2fs.File, rowIndex int, scaleColIdx int, lowerByFactor int) {
	sumWeights := 0
	for colIdx := treasureClassEx.Prob1; colIdx < treasureClassEx.Prob10; colIdx += 2 {

		sumWeights += first(strconv.Atoi(f.Rows[rowIndex][colIdx]))
	}
	noDrop, _ := strconv.Atoi(f.Rows[rowIndex][treasureClassEx.NoDrop])
	sumWeights += noDrop
	origWeight, _ := strconv.Atoi(f.Rows[rowIndex][scaleColIdx])

	for colIdx := treasureClassEx.Prob1; colIdx < treasureClassEx.Prob10; colIdx += 2 {
		if scaleColIdx == colIdx {
			continue // Adjust everything except scaleColIdx
		}
		f.Rows[rowIndex][colIdx] = strconv.Itoa(first(strconv.Atoi(f.Rows[rowIndex][colIdx])) * lowerByFactor)
	}
	f.Rows[rowIndex][scaleColIdx] = strconv.Itoa(origWeight)
	noDrop *= lowerByFactor
	noDrop += origWeight * (lowerByFactor - 1)
	f.Rows[rowIndex][treasureClassEx.NoDrop] = strconv.Itoa(noDrop)
}
func first(x int, _ error) int { return x }
