package splash

import (
	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/itemStatCost"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/magicSuffix"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/missiles"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/properties"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/propscores"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/skills"
)

const (
	patchstring = "patchstring.tbl"
	splashDir   = "/splash/"
)

// Run implement MeleeSplash functionality
func Run(outDir string, d2files d2fs.Files) {
	mergeSplashFile(missiles.FileName, d2files)
	mergeSplashFile(skills.FileName, d2files)
	mergeSplashFile(itemStatCost.FileName, d2files)
	mergeSplashFile(properties.FileName, d2files)
	mergeSplashFile(magicSuffix.FileName, d2files)
	//copyPatchString(outDir)	// Disabled because ElementalSkills contains a more complete copy of patchstrings.tbl
	// TODO: implement code to generate patchstrings.tbl instead of just copying
}

/*
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
*/
func mergeSplashFile(fileName string, d2files d2fs.Files) {
	splashFile := d2fs.ReadAsset(splashDir, fileName)

	modFile := d2files.Get(fileName)
	d2fs.AppendRows(modFile, *splashFile)
}

// DisableMeleeSplash comment out meleesplash in PropScores.txt
func DisableMeleeSplash(d2files d2fs.Files) {
	propScoresFile := d2files.GetWithPath(propscores.Path, propscores.FileName)
	for rowIdx := range propScoresFile.Rows {
		if propScoresFile.Rows[rowIdx][propscores.Prop] == "meleesplash" {
			propScoresFile.Rows[rowIdx][propscores.Prop] = "*meleesplash"
		}
	}
}
