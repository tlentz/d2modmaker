package treasure

import (
	"fmt"
	"strconv"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/treasureClassEx"
)

// BiggerGoldPiles Update TreasureClassEx.txt with fewer, larger gold piles.
func BiggerGoldPiles(d2files d2fs.Files) {
	f := d2files.Get(treasureClassEx.FileName)
	numChanges := 0
	for idx := range f.Rows {
		for colIdx := treasureClassEx.Item1; colIdx < treasureClassEx.Item10; colIdx += 2 {
			itm := f.Rows[idx][colIdx]
			if f.Rows[idx][treasureClassEx.TreasureClass] == "Gold" {
				// Don't mess with the first row, this is for player dropped gold
				continue
			}
			if len(itm) < 3 {
				continue
			}
			if itm[0:3] == "gld" {
				if len(itm) > 4 {
					goldMult, _ := strconv.Atoi(itm[8:])
					goldMult *= 10
					itm = "gld,mul=" + strconv.Itoa(goldMult)
				} else {
					itm = "gld,mul=1000"
				}
				f.Rows[idx][colIdx] = itm
				lowerProbs(f, idx, colIdx+1, 10)
				numChanges++
				break // Better not be more than one or the adjust code would screw it up royally...
			}
		}
	}
	fmt.Printf("BiggerGoldPiles: %d changes\n", numChanges)
}
