package perfectprops

import (
	"fmt"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/propscores"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/runes"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/setItems"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/sets"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/uniqueItems"
)

type propName2PropParType map[string]string

// Run Set props to be their perfect values, i.e. no min = max
func Run(d2files *d2fs.Files) {
	pmap := make(propName2PropParType)
	pmap.load(d2files)

	convertFile(d2files, pmap, &uniqueItems.IFI)
	convertFile(d2files, pmap, &setItems.IFI)
	convertFile(d2files, pmap, &sets.IFI)
	convertFile(d2files, pmap, &runes.IFI)
	fmt.Printf(">%s\n", pmap["death-skill"])
}
func convertFile(d2files *d2fs.Files, pmap propName2PropParType, ifi *d2fs.ItemFileInfo) {
	conversionCounter := 0
	file := d2files.Get(ifi.FI.FileName)
	for rowIdx := range file.Rows {
		if file.Rows[rowIdx][1] == "" {
			continue
		}
		for colIdx := ifi.FirstProp; colIdx < (ifi.FirstProp + ((ifi.NumProps - 1) * 4)); colIdx += 4 {
			propName := file.Rows[rowIdx][colIdx]
			parType := pmap[propName]
			switch {
			case (parType == "r") || (parType == "rp") || (parType == "rt") || (parType == "smm"):
				if file.Rows[rowIdx][colIdx+2] == file.Rows[rowIdx][colIdx+3] {
					conversionCounter++
				}
				file.Rows[rowIdx][colIdx+2] = file.Rows[rowIdx][colIdx+3] // Set prop Min = Max

			default:
				// all other, and unknown prop types aren't touched
			}

		}
	}
	fmt.Printf("%s skills perfected: %d\n", ifi.FI.FileName, conversionCounter)
}

func (pmap *propName2PropParType) load(d2files *d2fs.Files) {
	propScoresFile := d2files.GetWithPath(propscores.Path, propscores.FileName)
	for rowIdx := range propScoresFile.Rows {
		propName := propScoresFile.Rows[rowIdx][propscores.Prop]
		propParType := propScoresFile.Rows[rowIdx][propscores.PropParType]
		(*pmap)[propName] = propParType
	}

}
