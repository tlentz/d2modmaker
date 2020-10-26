package qol

import (
	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/charStats"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/uniqueItems"
)

func StartWithCube(d2files d2fs.Files) {
	f := d2files.Get(charStats.FileName)
	itemOffset := charStats.Item1
	countOffset := 2
	for idx, row := range f.Rows {
		for i := itemOffset; i < len(row)-countOffset; i += 3 {
			if (row[i] == "0" || row[i] == "") && (row[i+countOffset] == "0" || row[i+countOffset] == "") {
				f.Rows[idx][i] = "box"
				f.Rows[idx][i+countOffset] = "1"
				break // we added a cube, we are done with this row
			}
		}
	}
}

func RemoveUniqCharmLimit(d2files d2fs.Files) {

	uniqtxt := d2files.Get(uniqueItems.FileName)
	for i := range uniqtxt.Rows {
		name := uniqtxt.Rows[i][uniqueItems.Index]
		if name == uniqueItems.Anni || name == uniqueItems.Torch || name == uniqueItems.Gheed {
			uniqtxt.Rows[i][uniqueItems.Carry1] = ""
		}
	}

}
