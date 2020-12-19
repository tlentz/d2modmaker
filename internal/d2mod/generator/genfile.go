package generator

import (
	"fmt"
	"log"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/filenumbers"
	"github.com/tlentz/d2modmaker/internal/d2mod/d2items"
	"github.com/tlentz/d2modmaker/internal/util"
)

func genFile(g *Generator, ifi *d2fs.ItemFileInfo) {
	fmt.Println(ifi.FI.FileName)
	g.IFI = ifi
	numClones := 0
	if g.IFI.FI.FileNumber == filenumbers.UniqueItems {
		// Unfortunately cloning doesn't really work for any of the other types
		numClones = util.MaxInt(0, g.opts.NumClones)
	}

	f := g.d2files.Get(g.IFI.FI.FileName)
	pg := d2items.NewPropGetter(g.d2files, g.IFI, g.psi, *g.TypeTree)
	if (g.IFI.NumProps == 0) || (g.IFI.FirstProp == 0) {
		log.Fatalf("genfile: NumProps == %d, FirstProp = %d", g.IFI.NumProps, g.IFI.FirstProp)
	}
	numRows := len(f.Rows)
	for i := numClones; i >= 0; i-- {

		//fmt.Printf("genFile: %s %d\n", fileName, i)
		for rowIdx := 0; rowIdx < numRows; rowIdx++ {
			row := f.Rows[rowIdx]
			newRow := row
			item := d2items.NewItem(*pg, rowIdx, row)
			isClone := false
			if item != nil {
				if item.Lvl != 0 {
					if (g.IFI.HasEnabledColumn && (row[2] == "1")) || !g.IFI.HasEnabledColumn { // skip quest items
						isClone = true
						newi := GenItem(g, item)
						if newi.Affixes[0].ColIdx == 0 {
							log.Panicf("genFile: 0 affixes: %s|%s|%s|%s", newi.Affixes[0].P.Name, newi.Affixes[0].P.Par, newi.Affixes[0].P.Min, newi.Affixes[0].P.Max)
						}
						newRow = newi.ToRow(*pg, row)
					}
				}
			}
			if i != 0 {
				if isClone {
					f.Rows = append(f.Rows, newRow)
				}

			} else {
				f.Rows[rowIdx] = newRow
			}
		}

	}
}
