package stacksizes

import (
	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/misc"
)

func Increase(d2files d2fs.Files) {
	f := d2files.Get(misc.FileName)
	for idx, row := range f.Rows {
		if row[misc.Name] == misc.TownPortalBook || row[misc.Name] == misc.IdentifyBook || row[misc.Name] == misc.SkeletonKey {
			f.Rows[idx][misc.MaxStack] = "100"
		}
		if row[misc.Name] == misc.Arrows || row[misc.Name] == misc.Bolts {
			f.Rows[idx][misc.MaxStack] = "511"
		}
	}
}
