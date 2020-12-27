package cuberecipes

import (
	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/cubeMain"
)

// SafeUnsocket Recipe to remove runes/gems from items safely (Quiver + item=>Item + runes/gems)
func SafeUnsocket(d2files d2fs.Files) {
	cubeF := d2files.Get(cubeMain.FileName)
	tmp := make([]string, len(cubeF.Headers))

	tmp[cubeMain.Description] = "SafeUnsocket"
	tmp[cubeMain.Enabled] = "1"
	tmp[cubeMain.Version] = "100"
	tmp[cubeMain.NumInputs] = "2"
	tmp[cubeMain.Input1] = "any,sock"
	tmp[cubeMain.Input2] = "aqv,qty=1"
	tmp[cubeMain.Output] = "useitem,rem"
	tmp[cubeMain.Eol] = "0"

	cubeF.Rows = append(cubeF.Rows, tmp)

}
