package scorer

import (
	"github.com/tlentz/d2modmaker/internal/d2mod/d2items"
	"github.com/tlentz/d2modmaker/internal/d2mod/scorer/scorerstatistics"
)

// scoreAffix Score an Affix. Sets the affixes RawScore field, and returns the raw score
func scoreAffix(ss *scorerstatistics.ScorerStatistics, tt *d2items.TypeTree, item *d2items.Item, aff *d2items.Affix) int {
	score := scoreProp(ss, tt, item, aff.P, aff.Line)
	if ss != nil {
		ss.AddScoreLineWeight(tt, item, aff.Line)
	}
	aff.RawScore = score

	// Don't adjust score here because SynergyGroup multiplier isn't known yet
	//aff.AdjustedScore = util.Round32(aff.SetBonusMultiplier * float32(score))
	return aff.RawScore
}
