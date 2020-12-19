package scorer

import (
	"log"
	"strconv"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/itemscores"
	"github.com/tlentz/d2modmaker/internal/d2mod/d2items"
	"github.com/tlentz/d2modmaker/internal/d2mod/scorer/scorerstatistics"
)

/*
func genItemScores(g *Generator, ifi *d2fs.ItemFileInfo) {
	//itemScoreFile := d2fs.ReadAsset(itemscores.Path, itemscores.FileName)
	itemScoreFile := g.d2files.GetWithPath(itemscores.Path, itemscores.FileName)
	itemFile := g.d2files.Get(ifi.FI.FileName)
	pg := d2items.NewPropGetter(g.d2files, g.opts, ifi, g.psi, *g.TypeTree)
	for rowIdx, row := range itemFile.Rows {
		item := d2items.NewItem(*pg, rowIdx, row)
		if item == nil {
			continue
		}
		if item.Lvl == 0 {
			continue
		}
		newRow := make([]string, itemscores.NumColumns)
		newRow[itemscores.File] = ifi.FI.FileName[:len(ifi.FI.FileName)-4] // Chop off the .txt extension
		newRow[itemscores.Item] = row[ifi.ItemName]
		if ifi.Lvl > 0 {
			newRow[itemscores.Lvl] = item.Lvl
		} else {
			newRow[itemscores.Lvl] = ""
		}
		newRow[itemscores.VanillaScore] = strconv.Itoa(g.Statistics.ItemScores[row[ifi.ItemName]])
		s := scorer.ScoreItem(g.Statistics, g.TypeTree, *item)
		newRow[itemscores.ItemScore] = strconv.Itoa(s)
		itemScoreFile.Rows = append(itemScoreFile.Rows, newRow)
	}
	// path.Join(d2files.outDir, d2file.FileName)
	//d2fs.DebugDumpFiles(*g.d2files, itemscores.FileName)
}
*/ // WriteItemScore write Item out to ItemScores.txt

// WriteItemScore write Item out to ItemScores.txt
func WriteItemScore(ss *scorerstatistics.ScorerStatistics, d2files *d2fs.Files, ifi *d2fs.ItemFileInfo, item *d2items.Item, VanillaFlag bool) {
	itemScoreFile := d2files.GetWithPath(itemscores.Path, itemscores.FileName)
	if item == nil {
		return
	}
	if item.Lvl == 0 {
		return
	}
	newRow := make([]string, itemscores.NumColumns)
	newRow[itemscores.File] = ifi.FI.FileName[:len(ifi.FI.FileName)-4] // Chop off the .txt extension
	newRow[itemscores.Item] = item.Name
	if ifi.Lvl > 0 {
		newRow[itemscores.Lvl] = strconv.Itoa(item.Lvl)
	} else {
		newRow[itemscores.Lvl] = ""
	}
	newRow[itemscores.Pbucket] = ss.GetBucketName(item)
	if VanillaFlag {
		newRow[itemscores.VanillaFlag] = "Y"
	}
	if item.Score == 0 && item.Lvl > 50 {
		log.Panicf("%+v", item)
	}
	newRow[itemscores.ItemScore] = strconv.Itoa(item.Score)
	for idx, aff := range item.Affixes {
		newRow[itemscores.Prop1+(idx)*6] = aff.P.Name
		newRow[itemscores.Prop1+(idx)*6+1] = aff.P.Par
		newRow[itemscores.Prop1+(idx)*6+2] = aff.P.Min
		newRow[itemscores.Prop1+(idx)*6+3] = aff.P.Max
		newRow[itemscores.Prop1+(idx)*6+4] = strconv.Itoa(aff.AdjustedScore)
		newRow[itemscores.Prop1+(idx)*6+5] = strconv.FormatFloat(float64(aff.SetBonusMultiplier*aff.SynergyMultiplier), 'f', 2, 64)
	}
	newRow[itemscores.FI.NumColumns-1] = "eol"
	itemScoreFile.Rows = append(itemScoreFile.Rows, newRow)
}
