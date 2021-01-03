package scorer

import (
	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/propscores"
	"github.com/tlentz/d2modmaker/internal/d2mod/config"
	"github.com/tlentz/d2modmaker/internal/d2mod/d2items"
	"github.com/tlentz/d2modmaker/internal/d2mod/scorer/scorerstatistics"
)

func buff2HanderScore(opts config.GeneratorOptions, s *scorerstatistics.ScorerStatistics, tt *d2items.TypeTree, item d2items.Item) {

	if d2items.CheckTwoHander(tt, item) {
		s.ItemScores[item.Name] = (s.ItemScores[item.Name] * 125) / 100 // buff by 25%
		//fmt.Printf("2handbuff: %s\n", item.Name)
	}
	if d2items.CheckBow(tt, item) {
		s.ItemScores[item.Name] = (s.ItemScores[item.Name] * 115) / 100 // buff by 15%
		//fmt.Printf("2handbuff: %s\n", item.Name)
	}

}
func buff2HanderWeights(s *scorerstatistics.ScorerStatistics, d2files d2fs.Files) {
	f := d2files.Get(propscores.FileName)
	for rowIdx := range f.Rows {
		if f.Rows[rowIdx][propscores.SourceItem] == "2handers" {
			for tsIdx := range s.TypeStatistics {
				if s.TypeStatistics[tsIdx].ItemType == "2h" {
					s.TypeStatistics[tsIdx].NumLines[rowIdx] += 100
				}

			}
		}
	}
}
