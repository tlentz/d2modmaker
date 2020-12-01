package scorer

import (
	"log"
	"math"

	"github.com/tlentz/d2modmaker/internal/d2mod/d2items"
	"github.com/tlentz/d2modmaker/internal/d2mod/scorer/scorerstatistics"
	"github.com/tlentz/d2modmaker/internal/util"
)

// ScoreItem .
func ScoreItem(ss *scorerstatistics.ScorerStatistics, tt *d2items.TypeTree /* scorelines *propscores.ScoreMap,*/, item d2items.Item) int {

	var scores []int = make([]int, len(item.Affixes)) // scores by affix index, needed for synergy calc
	//var sgroups []string = make([]string, len(item.Affixes)) //  synergy groups to apply syn.  individually
	sgroups := make([]string, len(item.Affixes))
	if len(item.Affixes) == 0 {
		log.Panic("No Affixes")
	}
	for idx, aff := range item.Affixes {
		//sbm := CalcSetBonusMultiplier(item.FileNumber, item, pidx)

		//pscore, synergygroup := scoreProp(ss, tt, scorelines, item, item.Affixes[pidx])
		if aff.Line == nil {
			log.Fatalf("ScoreItem: Affix with no line encountered %s", item.Name)
		}

		sgroups[idx] = aff.Line.SynergyGroup
		affixRawScore := scoreAffix(ss, tt, item, aff)
		if aff.SetBonusMultiplier == 0 {
			log.Panic("SetBonumMultiplier == 0")
		}
		scores = append(scores, util.Round32(float32(affixRawScore)*aff.SetBonusMultiplier))
	}

	// Compute & add Synergy bonuses
	//NextProp:
	for idx, aff := range item.Affixes {

		if sgroups[idx] != "" {
			for oidx, oaff := range item.Affixes {
				if (&aff != &oaff) && (sgroups[idx] == sgroups[oidx]) {
					//fmt.Printf("SynBonus:%s:%s<%s>%s %d", item.Name, p.Name, sgroups[pidx], op.Name, scores[pidx])
					scores[idx] = util.Round32(float32(scores[idx]) * synergyBonus)
					//fmt.Printf("->%d\n", scores[pidx])
					//continue NextProp	// per macohan, 10% per other prop in same synergygroup. uncomment for flat 10%
				}
			}
		}
	}
	//log.Printf("Scores: %v", scores)
	score := 0
	for _, s := range scores {
		score += s
	}

	// Check if is a 2hander, and apply 2hander nerf if it is
	// pole staf bow xbow abow aspe spea
	//fmt.Printf("2h check: %s\n", item.Name)
	if d2items.CheckTwoHander(tt, item) {
		//fmt.Printf("2hander: %s\n", item.Name)
		score = int(math.Round(float64(score) * twoHandNerf))
	}
	//log.Printf("ScoreItem:\t%s\t%d\n", item.Name, score)
	return score
}
