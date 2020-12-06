package scorer

import (
	"fmt"
	"log"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2mod/d2items"
)

func scoreFile(s *Scorer, ifi *d2fs.ItemFileInfo) {
	s.IFI = ifi
	pgetter := d2items.NewPropGetter(s.d2files, s.opts, ifi, s.PSI, *s.TypeTree)
	_, items := pgetter.GetProps()
	fmt.Printf("Scoring %s:%d\n", s.IFI.FI.FileName, len(items))
	for idx := range items {
		s.items[items[idx].Name] = &items[idx]
		s.Statistics.ItemScores[items[idx].Name] = ScoreItem(s.Statistics, s.TypeTree, &items[idx])
		if items[idx].Score == 0 && items[idx].Lvl > 50 {
			log.Panicf("%+v", items[idx])
		}
		WriteItemScore(s.d2files, s.IFI, &items[idx], true)
	}
	return
}
