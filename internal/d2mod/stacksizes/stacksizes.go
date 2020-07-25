package stacksizes

import (
	"github.com/tlentz/d2modmaker/internal/d2file"
	"github.com/tlentz/d2modmaker/internal/d2file/txts/misc"
)

func Increase(d2files d2file.D2Files) {
	f := d2file.GetOrCreateFile(d2files, misc.FileName)
	for idx, row := range f.Rows {
		if row[misc.Name] == misc.TownPortalBook || row[misc.Name] == misc.IdentifyBook || row[misc.Name] == misc.SkeletonKey {
			f.Rows[idx][misc.MaxStack] = "100"
		}
		if row[misc.Name] == misc.Arrows || row[misc.Name] == misc.Bolts {
			f.Rows[idx][misc.MaxStack] = "511"
		}
	}
}
