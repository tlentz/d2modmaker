package townskills

import (
	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/missiles"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/skills"
)

func Enable(d2files d2fs.Files) {
	skillstxt := d2files.Get(skills.FileName)
	for i := range skillstxt.Rows {
		skillstxt.Rows[i][skills.InTown] = "1"
	}
	missilestxt := d2files.Get(missiles.FileName)
	for i := range missilestxt.Rows {
		missilestxt.Rows[i][missiles.Town] = "1"
	}
}
