package qol

import (
	"github.com/tlentz/d2modmaker/internal/d2file"
	"github.com/tlentz/d2modmaker/internal/d2file/txts/charStats"
)

func StartWithCube(d2files d2file.D2Files) {
	f := d2file.GetOrCreateFile(d2files, charStats.FileName)
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
