package splash

import (
	"io"
	"os"

	"github.com/tlentz/d2modmaker/internal/assets"
	"github.com/tlentz/d2modmaker/internal/d2file"
	"github.com/tlentz/d2modmaker/internal/itemStatCostTxt"
	"github.com/tlentz/d2modmaker/internal/magicSuffixTxt"
	"github.com/tlentz/d2modmaker/internal/missilesTxt"
	"github.com/tlentz/d2modmaker/internal/propertiesTxt"
	"github.com/tlentz/d2modmaker/internal/skillsTxt"
	"github.com/tlentz/d2modmaker/internal/util"
)

const (
	patchstring = "patchstring.tbl"
	splashDir   = "/splash/"
)

func SplashJewels(outDir string, d2files d2file.D2Files) {
	mergeSplashFile(missilesTxt.FileName, d2files)
	mergeSplashFile(skillsTxt.FileName, d2files)
	mergeSplashFile(itemStatCostTxt.FileName, d2files)
	mergeSplashFile(propertiesTxt.FileName, d2files)
	mergeSplashFile(magicSuffixTxt.FileName, d2files)
	copyPatchString(outDir)
}

func copyPatchString(outDir string) {
	from, err := assets.Assets.Open(splashDir + patchstring)
	util.Check(err)
	defer from.Close()

	patchstringDir := outDir + assets.PatchStringDest
	err2 := os.MkdirAll(patchstringDir, 0755)
	util.Check(err2)

	to, err := os.OpenFile(patchstringDir+patchstring, os.O_RDWR|os.O_CREATE, 0666)
	util.Check(err)
	defer to.Close()

	_, err = io.Copy(to, from)
	util.Check(err)
}

func mergeSplashFile(fileName string, d2files d2file.D2Files) {
	splashFile, err := d2file.ReadD2File(assets.AssetFS, fileName, splashDir)
	util.Check(err)

	modFile := d2file.GetOrCreateFile(d2files, fileName)
	d2file.MergeRows(modFile, *splashFile)
}
