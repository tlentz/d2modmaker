package propdebug

import (
	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/cubeMain"
)

const (
	patchstring    = "patchstring.tbl"
	debugAssetsDir = "/propdebug/"
)

// Run Enable/Disable ElementalSkills
func Run(d2files d2fs.Files) {
	d2fs.MergeRows(d2files.Get(cubeMain.FileName), *d2fs.ReadAsset(debugAssetsDir, cubeMain.FileName))
}
