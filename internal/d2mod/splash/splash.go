package splash

import (
	"io"
	"os"
	"path"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/assets"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/itemStatCost"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/magicSuffix"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/missiles"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/properties"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/skills"
	"github.com/tlentz/d2modmaker/internal/util"
)

const (
	patchstring = "patchstring.tbl"
	splashDir   = "/splash/"
)

func Jewels(outDir string, d2files d2fs.Files) {
	mergeSplashFile(missiles.FileName, d2files)
	mergeSplashFile(skills.FileName, d2files)
	mergeSplashFile(itemStatCost.FileName, d2files)
	mergeSplashFile(properties.FileName, d2files)
	mergeSplashFile(magicSuffix.FileName, d2files)
	copyPatchString(outDir)
}

func copyPatchString(outDir string) {
	from, err := assets.Assets.Open(splashDir + patchstring)
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

func mergeSplashFile(fileName string, d2files d2fs.Files) {
	splashFile := d2fs.ReadAsset(fileName, splashDir)

	modFile := d2files.Get(fileName)
	d2fs.MergeRows(modFile, *splashFile)
}
