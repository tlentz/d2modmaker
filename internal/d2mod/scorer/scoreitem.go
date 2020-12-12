package scorer

import (
	"fmt"
	"log"

	"github.com/tlentz/d2modmaker/internal/d2mod/d2items"
	"github.com/tlentz/d2modmaker/internal/d2mod/scorer/scorerstatistics"
	"github.com/tlentz/d2modmaker/internal/util"
)

// ScoreItem .
func ScoreItem(ss *scorerstatistics.ScorerStatistics, tt *d2items.TypeTree /* scorelines *propscores.ScoreMap,*/, item *d2items.Item) int {

	//var sgroups []string = make([]string, len(item.Affixes))
	if len(item.Affixes) == 0 {
		log.Panic("No Affixes")
	}
	sgroups := make([]string, len(item.Affixes)) //  synergy groups to apply syn.  individually
	for idx := range item.Affixes {
		//sbm := CalcSetBonusMultiplier(item.FileNumber, item, pidx)

		//pscore, synergygroup := scoreProp(ss, tt, scorelines, item, item.Affixes[pidx])
		if item.Affixes[idx].Line == nil {
			log.Fatalf("ScoreItem: Affix with no line encountered %s", item.Name)
		}

		sgroups[idx] = item.Affixes[idx].Line.SynergyGroup
		scoreAffix(ss, tt, item, &item.Affixes[idx])
		//fmt.Printf("~%s %+v %d~\n", item.Affixes[idx].P.Name, item.Affixes[idx], affScore)

		if item.Affixes[idx].SetBonusMultiplier == 0 {
			log.Panic("SetBonumMultiplier == 0")
		}
	}

	// Compute & add Synergy bonuses
	//NextProp:
	for idx := range item.Affixes {
		item.Affixes[idx].SynergyMultiplier = 1
		if sgroups[idx] != "" {
			for oidx := range item.Affixes {
				if (idx != oidx) && (sgroups[idx] == sgroups[oidx]) {
					//fmt.Printf("SynBonus:%s:%s<%s>%s %d", item.Name, p.Name, sgroups[pidx], op.Name, scores[pidx])
					//scores[idx] = util.Round32(float32(scores[idx]) * synergyBonus)
					item.Affixes[idx].SynergyMultiplier = item.Affixes[idx].SynergyMultiplier * synergyBonus
					//fmt.Printf("->%d\n", scores[pidx])
					//continue NextProp	// per macohan, 10% per other prop in same synergygroup. uncomment for flat 10%
					//fmt.Printf("%d", item.Affixes[idx].RawScore)
				}
			}
		}
	}
	// Check if is a 2hander, and apply 2hander nerf if it is
	// pole staf bow xbow abow aspe spea
	//fmt.Printf("2h check: %s\n", item.Name)
	// Disabling 2hander math
	// if d2items.CheckTwoHander(tt, *item) {
	// 	//fmt.Printf("2hander: %s\n", item.Name)
	// 	//score = int(math.Round(float64(score) * twoHandNerf))
	// 	for idx := range item.Affixes {
	// 		item.Affixes[idx].ScoreMult = item.Affixes[idx].ScoreMult * twoHandNerf
	// 	}
	// }

	//log.Printf("Scores: %v", scores)
	score := 0
	for idx := range item.Affixes {
		if item.Affixes[idx].SetBonusMultiplier == 0 || item.Affixes[idx].SynergyMultiplier == 0 {
			fmt.Printf("%v\n", item)
			log.Panicf("ScoreItem: a score multiplier is not set")
		}
		item.Affixes[idx].AdjustedScore = util.Round32(float32(item.Affixes[idx].RawScore) * item.Affixes[idx].SetBonusMultiplier * item.Affixes[idx].SynergyMultiplier)
		score += item.Affixes[idx].AdjustedScore
	}
	//log.Printf("ScoreItem:\t%s\t%d\n", item.Name, score)
	item.Score = score
	return score
}
