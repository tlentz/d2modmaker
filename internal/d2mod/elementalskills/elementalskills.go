package elementalskills

import (
	"io"
	"os"
	"path"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/assets"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/itemStatCost"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/properties"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/propscores"
	"github.com/tlentz/d2modmaker/internal/d2mod/d2items"
	"github.com/tlentz/d2modmaker/internal/d2mod/prop"
	"github.com/tlentz/d2modmaker/internal/d2mod/scorer/scorerstatistics"
	"github.com/tlentz/d2modmaker/internal/util"
)

const (
	patchstring        = "patchstring.tbl"
	elementalAssetsDir = "/elementalskills/"
)

// Run Enable/Disable ElementalSkills
func Run(outDir string, d2files d2fs.Files, enabled bool) {

	copyPatchString(outDir)
	f := d2files.Get(properties.FileName)
	d2fs.MergeRows(d2files.Get(itemStatCost.FileName), *d2fs.ReadAsset(elementalAssetsDir, itemStatCost.FileName))
	d2fs.MergeRows(d2files.Get(properties.FileName), *d2fs.ReadAsset(elementalAssetsDir, properties.FileName))

	if enabled {
	} else {
		propScoresFile := d2files.GetWithPath(propscores.Path, propscores.FileName)
		for rowIdx := range propScoresFile.Rows {
			switch propScoresFile.Rows[rowIdx][propscores.Prop] {
			case
				"lightningskill",
				"magicskill",
				"coldskill",
				"poisonskill":
				propScoresFile.Rows[rowIdx][propscores.Prop] = "*" + propScoresFile.Rows[rowIdx][propscores.Prop]
			}
		}

	}
}

// SetProbability Increase probability of getting the other elementalskills to match fireskills
func SetProbability(d2files d2fs.Files, ss *scorerstatistics.ScorerStatistics, enabled bool) {
	if enabled {
		fireSkillsRows := make(map[int]bool, 0)
		elemSkillsRows := make(map[int]bool, 0)
		psf := d2files.Get(propscores.FileName)
		for rowIdx := range psf.Rows {
			switch psf.Rows[rowIdx][propscores.Prop] {
			case "fireskill":
				fireSkillsRows[rowIdx] = true
			case
				"lightningskill",
				"magicskill",
				"coldskill",
				"poisonskill":
				elemSkillsRows[rowIdx] = true
			}
		}
		// FireSkillsRows
		for t := range ss.TypeStatistics {
			fireSkillWeight := 0
			for propScoresRowIdx := range fireSkillsRows {
				fireSkillWeight += ss.TypeStatistics[t].NumLines[propScoresRowIdx]
			}
			fireSkillWeight = (fireSkillWeight*10 + 5) / 30 // Integer Ghetto rounding
			for propScoresRowIdx := range fireSkillsRows {
				ss.TypeStatistics[t].NumLines[propScoresRowIdx] = fireSkillWeight
			}
			for propScoresRowIdx := range elemSkillsRows {
				ss.TypeStatistics[t].NumLines[propScoresRowIdx] = fireSkillWeight
			}
		}
	} else {

	}
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
func Props() d2items.Props {
	props := make(d2items.Props, 4)
	props[0] = prop.NewProp("coldskill", "", "1", "4")
	props[1] = prop.NewProp("poisonskill", "", "1", "4")
	props[2] = prop.NewProp("lightningskill", "", "1", "4")
	props[3] = prop.NewProp("magicskill", "", "1", "4")

	return props
}
