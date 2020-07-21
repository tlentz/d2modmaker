package splash

import (
	"github.com/tlentz/d2modmaker/internal/d2file"
	"github.com/tlentz/d2modmaker/internal/itemStatCostTxt"
	"github.com/tlentz/d2modmaker/internal/magicSuffixTxt"
	"github.com/tlentz/d2modmaker/internal/missilesTxt"
	"github.com/tlentz/d2modmaker/internal/propertiesTxt"
	"github.com/tlentz/d2modmaker/internal/skillsTxt"
	"github.com/tlentz/d2modmaker/internal/util"
)

func SplashJewels(d2files d2file.D2Files) {
	mergeSplashFile(missilesTxt.FileName, d2files)
	mergeSplashFile(skillsTxt.FileName, d2files)
	mergeSplashFile(itemStatCostTxt.FileName, d2files)
	mergeSplashFile(propertiesTxt.FileName, d2files)
	mergeSplashFile(magicSuffixTxt.FileName, d2files)
}

func writePatchStrings() {

}

func mergeSplashFile(fileName string, d2files d2file.D2Files) {
	splashDir := "/splash/"
	splashFile, err := d2file.ReadD2File(fileName, splashDir)
	util.Check(err)

	modFile := d2file.GetOrCreateFile(d2files, fileName)
	d2file.MergeRows(modFile, *splashFile)
}
