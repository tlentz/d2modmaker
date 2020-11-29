package scorer

import (
	"github.com/tlentz/d2modmaker/internal/d2mod/d2items"
	"github.com/tlentz/d2modmaker/internal/d2mod/scorer/scorerstatistics"
	"github.com/tlentz/d2modmaker/internal/util"
)

// scoreAffix Applies set bonuses to what scoreProp returns
func scoreAffix(ss *scorerstatistics.ScorerStatistics, tt *d2items.TypeTree, item d2items.Item, aff d2items.Affix) int {
	score := scoreProp(ss, tt, item, aff.P, aff.Line)
	if ss != nil {
		ss.AddScoreLineWeight(tt, item, aff.Line)
	}
	return util.Round32(aff.SetBonusMultiplier * float32(score))
}
