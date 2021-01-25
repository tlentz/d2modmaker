package elementalskills

import (
	"io"
	"os"
	"path"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/assets"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/itemStatCost"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/properties"
	"github.com/tlentz/d2modmaker/internal/d2mod/prop"
	"github.com/tlentz/d2modmaker/internal/util"
)

const (
	patchstring        = "patchstring.tbl"
	elementalAssetsDir = "/elementalskills/"
)

// Run Enable/Disable ElementalSkills
func Run(outDir string, d2files d2fs.Files, enabled bool) {

	copyPatchString(outDir)
	d2fs.MergeRows(d2files.Get(itemStatCost.FileName), *d2fs.ReadAsset(elementalAssetsDir, itemStatCost.FileName))
	d2fs.MergeRows(d2files.Get(properties.FileName), *d2fs.ReadAsset(elementalAssetsDir, properties.FileName))

}

func copyPatchString(outDir string) {
	from, err := assets.Assets.Open(elementalAssetsDir + patchstring)
	util.Check(err)
	defer from.Close()

	patchstringDir := path.Join(outDir, assets.PatchStringDest)
	err2 := os.MkdirAll(patchstringDir, 0755)
	util.Check(err2)

	to, err := os.OpenFile(path.Join(patchstringDir, patchstring), os.O_RDWR|os.O_CREATE, 0666)
	util.Check(err)
	defer to.Close()

	_, err = io.Copy(to, from)
	util.Check(err)
}

// Props Add elemental skill props to a list of props
func Props() prop.Props {
	props := make(prop.Props, 4)
	props[0] = prop.NewProp("coldskill", "", "1", "4", 0)
	props[1] = prop.NewProp("poisonskill", "", "1", "4", 0)
	props[2] = prop.NewProp("lightningskill", "", "1", "4", 0)
	props[3] = prop.NewProp("magicskill", "", "1", "4", 0)

	return props
}
