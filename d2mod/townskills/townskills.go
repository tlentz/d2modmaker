package townskills

import (
	"github.com/tlentz/d2modmaker/internal/d2file"
	"github.com/tlentz/d2modmaker/internal/d2file/txts/missiles"
	"github.com/tlentz/d2modmaker/internal/d2file/txts/skills"
)

func Enable(d2files d2file.D2Files) {
	skillstxt := d2file.GetOrCreateFile(d2files, skills.FileName)
	for i := range skillstxt.Rows {
		skillstxt.Rows[i][skills.InTown] = "1"
	}
	missilestxt := d2file.GetOrCreateFile(d2files, missiles.FileName)
	for i := range missilestxt.Rows {
		missilestxt.Rows[i][missiles.Town] = "1"
	}
}
