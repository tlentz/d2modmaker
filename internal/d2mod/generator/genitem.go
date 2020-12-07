package generator

import (
	"log"
	"math/rand"

	"github.com/tlentz/d2modmaker/internal/d2mod/d2items"
	"github.com/tlentz/d2modmaker/internal/d2mod/scorer"
	"github.com/tlentz/d2modmaker/internal/util"
)

// GenItem returns a new item with random affixes based on supplied item
func GenItem(g *Generator, oldItem *d2items.Item) *d2items.Item {
	if g.Statistics == nil {
		log.Fatalln("GenItems: generator statistics was not initialized")
	}
	its := g.Statistics.GetItemTypeStatistics(g.TypeTree, oldItem.FileNumber, oldItem.Types[0]) // FIXME: This is broken for runewords that work in both weapons & armor...
	if len(its.NumLines) == 0 {
		log.Panic("item statistics structure empty")
	}
	its.Weights.Generate()
	targetScore := util.Round32(float32(g.Statistics.ItemScores[oldItem.Name]) * float32(g.opts.PropScoreMultiplier))
	targetPropCount := 0
	maxProps := 0
	if g.opts.BalancedPropCount {
		targetPropCount = len(oldItem.Affixes)
		maxProps = util.MinInt(targetPropCount+4, g.IFI.NumProps) // With BalancedPropCount allow 4 extra props worth of breathing room to hit the score target
	} else {
		if g.IFI.NumProps == 0 {
			log.Panic("g.IFI.NumProps == 0")
		}
		minProps := util.MinInt(g.opts.MinProps, g.IFI.NumProps)
		minProps = util.MaxInt(minProps, 1) // Because everyone wants some props
		maxProps = util.MinInt(g.opts.MaxProps, g.IFI.NumProps)

		targetPropCount = minProps + rand.Intn(maxProps-minProps+1) // beware how Intn behaves...
		targetPropCount = util.MinInt(targetPropCount, g.IFI.NumProps)
	}
	newi := oldItem.CloneWithoutAffixes()
	newi.Affixes = []d2items.Affix{}
	var rollCount int
	itemScore := 0
	minItemScore := util.Round32(float32(targetScore) - util.Absf32((float32(targetScore) * 0.4)))
	maxItemScore := util.Round32(float32(targetScore) + util.Absf32((float32(targetScore) * 0.4)))

	// Loop while checking score
	//for ((itemScore < minItemScore) || (itemScore > maxItemScore)) && (rollCount < 200) {

	for rollCount = 0; rollCount < 10; rollCount++ {
		// Roll/Reroll entire item
		newi.Affixes = []d2items.Affix{}
		itemScore = 0
		//log.Printf("genItem: # Target Prop Count %d, Scores(min/tgt/max): %d/%d/%d %s", targetPropCount, minItemScore, targetScore, maxItemScore, item.Name)
		if targetPropCount == 0 {
			log.Printf("genItem: # Item Affixes=%d, Opts.MinProps=%d, Opts.MaxProps=%d, File NumProps=%d", len(newi.Affixes), g.opts.MinProps, g.opts.MaxProps, g.IFI.NumProps)
			log.Panic("genItem: 0 targetPropCount")
		}
		// TODO: Add better partial set bonus support
		// Roll each prop, checking that the item doesn't
		affixRollCount := 0
		for len(newi.Affixes) < maxProps { // len(newi.Affixes) < maxProps {

			targetPropScore := targetScore - itemScore
			if len(newi.Affixes) < (targetPropCount - 1) {
				targetPropScore = calcTargetPropScore(len(newi.Affixes), targetPropCount, maxProps, itemScore, targetScore)

				numAffixesLeft := util.MaxInt(0, targetPropCount-len(newi.Affixes))
				if numAffixesLeft > 0 {
					targetPropScore = util.MaxInt(0, int(float32((targetScore-itemScore))/float32(numAffixesLeft)))
				} else {
					targetPropScore = targetScore - itemScore
				}
			}
			newColIdx := len(newi.Affixes)*4 + g.IFI.FirstProp
			sbm := d2items.CalcSetBonusMultiplier(g.IFI.FI.FileNumber, newColIdx)
			targetPropScore = util.Round32(float32(targetPropScore) / sbm)
			//log.Printf("genItem: #Affixes: %d TargetScore: %d TargetPropScore: %d itemScore:%d", len(newi.Affixes), targetScore, targetPropScore, itemScore)

			newAffix := RollAffix(g, newi, newColIdx, targetPropScore, its.Weights)
			if newAffix.ScoreMult == 0 {
				log.Panicf("GenItem: RollAffix returned affix with ScoreMult == 0")
			}
			newi.Affixes = append(newi.Affixes, *newAffix)
			itemScore = scorer.ScoreItem(g.Statistics, g.TypeTree, newi) // Doing full item score because of SynergyGroup calculation
			// if itemScore == 0 {

			// 	log.Panicf("%+v\n", newi)
			// }
			affixRollCount++
			if (affixRollCount > 200) && (len(newi.Affixes) >= targetPropCount) {
				//log.Fatalf("GenItem: > 100 rolls for %s", newi.Name)
				break
			}
			if ((itemScore > minItemScore) && (itemScore < maxItemScore)) && (len(newi.Affixes) >= targetPropCount) {
				break
			}
		}
		if ((itemScore > minItemScore) && (itemScore < maxItemScore)) && (len(newi.Affixes) >= targetPropCount) {
			break
		}
	}
	// TODO: Write these statistics out to a file (Target score, Generated item score, # props Vanilla vs generated)
	//fmt.Printf("Item\t%s\t%d\t%d\t%d\n", newi.Name, newi.FileNumber, targetScore, itemScore)
	/*
		// Debugging
		lastAffixName := ""
		if len(newi.Affixes) > 0 {
			lastAffixName = newi.Affixes[len(newi.Affixes)-1].P.Name
		}
		log.Printf("genItem: Count(tgt/itm/max) %d/%d/%d, Scores(min/itm/tgt/max): %d/%d/%d/%d,  %s last aff=%s", len(newi.Affixes), targetPropCount, maxProps, minItemScore, itemScore, targetScore, maxItemScore, item.Name, lastAffixName)
	*/
	if (itemScore < minItemScore) || (itemScore > maxItemScore) {
		log.Printf("genItem: Missed target: #  Prop Count (tgt/itm/max) %2d/%2d/%d, Scores(Vanilla <min itm max>): %5d <%5d %5d %5d> -- %s", targetPropCount, len(newi.Affixes), maxProps, g.Statistics.ItemScores[oldItem.Name], minItemScore, itemScore, maxItemScore, newi.Name)
	}
	if minItemScore == 0 && maxItemScore == 0 {
		log.Panicf("%+v", newi)
	}
	scorer.WriteItemScore(g.d2files, g.IFI, newi, false)
	return newi
}
func calcTargetPropScore(numItemAffixes int, targetPropCount int, maxPropCount int, itemScore int, targetScore int) int {
	if numItemAffixes-1 >= targetPropCount {
		return targetScore - itemScore
	}
	numAffixesLeft := targetPropCount - numItemAffixes
	negDev := float32(0.0)
	posDev := float32(0)
	switch numAffixesLeft {
	case 1:
		log.Panic("shouldn't get here")
	case 2:
		negDev = 0.4
		posDev = 0.75
	case 3:
		negDev = 0.3
		posDev = 0.75
	case 4:
		negDev = -0.1
		posDev = 0.70
	default:
		negDev = -0.3
		posDev = 0.6
	}
	minScore := util.Round32(negDev * float32((targetScore - itemScore)))
	maxScore := util.Round32(posDev * float32((targetScore - itemScore)))

	return minScore + util.Round32(rand.Float32()*float32(maxScore-minScore))

}
func checkDupeGroups(item *d2items.Item) bool {
	for idxa, a := range item.Affixes {
		for idxb, b := range item.Affixes {
			if (a.Line.Group == b.Line.Group) && (a != b) {
				log.Printf("Duplicate group: %s %d == %d\n%v\n%v\n", a.Line.Group, idxa, idxb, a, b)
				return true
			}
		}
	}
	return false
}
