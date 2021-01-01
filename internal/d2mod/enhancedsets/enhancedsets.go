package enhancedsets

import (
	"fmt"
	"log"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/setItems"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/sets"
)

// SetAddFunc  Since the Sets.txt regenerating would alter set bonuses on existing items,
// all of the props must exist in SetItems.txt.  (Sets props have been blanked)
// AddFunc == "" would then not allow for the set piece to have any set bonuses, so
// force AddFunc to use mode 2, where the AProp* props are treated as partial set bonuses
func SetAddFunc(d2files *d2fs.Files, newAddFunc int) {
	if (newAddFunc > 2) || (newAddFunc < 0) {
		log.Panicf("SetAddFunc called with newAddFunc out of range")
	}
	f := d2files.Get(setItems.FI.FileName)
	for _, row := range f.Rows {
		if row[1] != "" {
			row[setItems.AddFunc] = fmt.Sprintf("%d", newAddFunc)
		}
	}
}

// BlankFullSetBonuses Blanks all properties pointed to by the scrambler structure.
func BlankFullSetBonuses(d2files *d2fs.Files) {
	f := d2files.Get(sets.FI.FileName)
	for _, row := range f.Rows {
		for propIndex := 0; propIndex < sets.IFI.NumProps; propIndex++ {
			i := sets.IFI.FirstProp + propIndex*4
			row[i] = ""
			row[i+1] = ""
			row[i+2] = ""
			row[i+3] = ""
		}
	}
}
