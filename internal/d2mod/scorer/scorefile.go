package scorer

import (
	"fmt"

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
		s.Statistics.ItemScores[items[idx].Name] = ScoreItem(s.Statistics, s.TypeTree, items[idx])
	}
	return
}
