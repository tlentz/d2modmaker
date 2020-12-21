package treasure

import (
	"fmt"
	"strconv"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/treasureClassEx"
)

// NoFlawGems Set TreasureClassEx columns with gems up to no drop flawed or flawless
func NoFlawGems(d2files d2fs.Files) {
	f := d2files.Get(treasureClassEx.FileName)
	numChanges := 0
	for idx := range f.Rows {
		for colIdx := treasureClassEx.Item1; colIdx < treasureClassEx.Item10; colIdx += 2 {
			itm := f.Rows[idx][colIdx]
			lvl, _ := strconv.Atoi(f.Rows[idx][treasureClassEx.Level])
			if lvl < 38 {
				// Don't touch Normal
				continue
			}
			if itm == "Flawless Gem" {
				f.Rows[idx][colIdx] = "Perfect Gem"
				lowerProbs(f, idx, colIdx+1, 3)
				numChanges++
			}
		}
	}
	for idx := range f.Rows {
		for colIdx := treasureClassEx.Item1; colIdx < treasureClassEx.Item10; colIdx += 2 {
			itm := f.Rows[idx][colIdx]
			lvl, _ := strconv.Atoi(f.Rows[idx][treasureClassEx.Level])
			if lvl < 38 {
				// Don't touch Normal
				continue
			}
			if itm == "Flawed Gem" {
				//prob,_ := strconv.Atoi(f.Rows[idx][colIdx+1])
				f.Rows[idx][colIdx+1] = "1" // Just make them unlikely
				numChanges++
			}
		}
	}
	fmt.Printf("BiggerGoldPiles: %d changes\n", numChanges)

}
