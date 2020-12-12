package generator

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/filenumbers"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/sets"
	"github.com/tlentz/d2modmaker/internal/d2mod/d2items"
	"github.com/tlentz/d2modmaker/internal/d2mod/scorer"
	"github.com/tlentz/d2modmaker/internal/util"
)

// GenItem returns a new item with random affixes based on supplied item
func GenItem(g *Generator, oldItem *d2items.Item) *d2items.Item {
	if g.Statistics == nil {
		log.Fatalln("GenItems: generator statistics was not initialized")
	}
	its := g.Statistics.GetItemTypeStatistics(oldItem) // FIXME: This is broken for runewords that work in both weapons & armor...
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

	for rollCount = 0; rollCount < 10; rollCount++ {
		// Roll/Reroll entire item
		newi.Affixes = []d2items.Affix{}
		itemScore = 0
		//log.Printf("genItem: # Target Prop Count %d, Scores(min/tgt/max): %d/%d/%d %s", targetPropCount, minItemScore, targetScore, maxItemScore, oldItem.Name)
		if targetPropCount == 0 {
			log.Printf("genItem: # Item Affixes=%d, Opts.MinProps=%d, Opts.MaxProps=%d, File NumProps=%d", len(newi.Affixes), g.opts.MinProps, g.opts.MaxProps, g.IFI.NumProps)
			log.Panic("genItem: 0 targetPropCount")
		}
		// TODO: Add better partial set bonus support
		// Roll each prop, checking that the item score doesn't go out of bounds (between minItemScore & maxItemScore) and < maxProps
		affixRollCount := 0
		for len(newi.Affixes) < maxProps { // len(newi.Affixes) < maxProps {

			minPropScore, targetPropScore, maxPropScore := calcTargetPropScore(len(newi.Affixes), targetPropCount, maxProps, itemScore, targetScore)
			if targetPropScore < minPropScore || targetPropScore > maxPropScore {
				log.Panicf("targetPropScore not between min & max")
			}
			newColIdx := getNewColIdx(*g.IFI, oldItem, newi)
			if newColIdx <= 0 {
				log.Panicf("%d colidx %+v", newColIdx, newi)
			}
			sbm := d2items.CalcSetBonusMultiplier(g.IFI.FI.FileNumber, newColIdx)
			targetPropScore = util.Round32(float32(targetPropScore) / sbm)
			//log.Printf("genItem: #Affixes: %d TargetScore: %d TargetPropScore: %d itemScore:%d", len(newi.Affixes), targetScore, targetPropScore, itemScore)

			newAffix := RollAffix(g, newi, newColIdx, minPropScore, targetPropScore, maxPropScore, its.Weights)
			if newAffix.SetBonusMultiplier == 0 {
				log.Panicf("GenItem: RollAffix returned affix with ScoreMult == 0")
			}
			newi.Affixes = append(newi.Affixes, *newAffix)
			//oldItemScore := itemScore
			itemScore = scorer.ScoreItem(g.Statistics, g.TypeTree, newi) // Doing full item score because of SynergyGroup calculation
			//affixScore := itemScore - oldItemScore

			//log.Printf("genItem: itemScore = %d %+v", itemScore, newi)
			// This check fails for LvlScale
			// if affixScore < minPropScore || affixScore > maxPropScore {
			// 	log.Printf("Problem with RollAffix: (Actual) min/tgt/max:(%d) %d %d %d", affixScore, minPropScore, targetPropScore, maxPropScore)
			// 	log.Printf("%+v", newAffix.P)
			// 	log.Printf("%+v", newAffix.Line)
			// 	log.Printf("%+v", newi)
			// 	panic(1)
			// }

			// if itemScore == 0 {
			// 	log.Panicf("%+v\n", newi)
			// }

			affixRollCount++
			if (affixRollCount > 200) && (len(newi.Affixes) >= targetPropCount) {
				log.Fatalf("GenItem: > 200 rolls for %s", newi.Name)
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
	scorer.WriteItemScore(g.Statistics, g.d2files, g.IFI, newi, false)
	return newi
}
func calcTargetPropScore(numItemAffixes int, targetPropCount int, maxPropCount int, itemScore int, targetScore int) (int, int, int) {
	numAffixesLeft := targetPropCount - numItemAffixes
	if numAffixesLeft <= 1 {
		delta := util.AbsInt((targetScore * 40) / 100)
		//fmt.Printf("#Props:%d, %d %d %d\n", numAffixesLeft, itemScore-delta, targetScore-itemScore, itemScore+delta)
		return targetScore - itemScore - delta, targetScore - itemScore, targetScore - itemScore + delta
	}
	negDev := float32(0.0)
	posDev := float32(0)
	switch numAffixesLeft {
	case 1:
		log.Panicf("shouldn't get here: %d %d", numItemAffixes, targetPropCount)
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
	if targetScore < itemScore {
		negDev, posDev = posDev, negDev
	}
	minScore := util.Round32(negDev * float32((targetScore - itemScore)))
	maxScore := util.Round32(posDev * float32((targetScore - itemScore)))
	targetPropScore := minScore + util.Round32(rand.Float32()*float32(maxScore-minScore))
	if targetPropScore > maxScore || targetPropScore < minScore {
		log.Panicf("targetPropScore out of bounds")
	}
	if itemScore > targetScore && (maxScore > 0) {
		if minScore > 0 {
			log.Panic("minScore should be negative")
		}
		maxScore = 0
		targetPropScore = minScore
	}
	if itemScore > 3*targetScore && (targetScore > 100) {
		fmt.Printf("Score Item/Tgt/#p - Min/Tgt/Max: %d/%d/%d - %d/%d/%d\n", itemScore, targetScore, numAffixesLeft, minScore, targetPropScore, maxScore)
		panic(1)
	}
	return minScore, targetPropScore, maxScore

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
func getNewColIdx(ifi d2fs.ItemFileInfo, oldItem *d2items.Item, newItem *d2items.Item) int {
	newItemAffixIndex := len(newItem.Affixes)
	newColIdx := -1
	switch {
	case oldItem.FileNumber == filenumbers.Sets:
		if len(newItem.Affixes) < len(oldItem.Affixes) {
			newColIdx = oldItem.Affixes[newItemAffixIndex].ColIdx
		} else {
			// search for a gap in full set bonuses first
			newColIdx = findUnusedColIdx(newItem, sets.FCode1, sets.FCode8)
			if newColIdx < 0 {
				// search for gap in partial set bonus area
				newColIdx = findUnusedColIdx(newItem, sets.FCode1, sets.FCode8)
				if newColIdx < 0 {
					// assertion: Shouldn't hit here unless # affix math is buggered
					log.Panicln("Logic error in getNewColIdx")
				}
			}
		}
	case oldItem.FileNumber == filenumbers.SetItems:
		if len(newItem.Affixes) < len(oldItem.Affixes) {
			newColIdx = oldItem.Affixes[newItemAffixIndex].ColIdx
		} else {
			newColIdx = findUnusedColIdx(newItem, ifi.FirstProp, (ifi.NumProps-1)*4+ifi.FirstProp)
			if newColIdx < 0 {
				log.Panicf("Logic error or ran out of affixes")
			}
		}
	default:
		newColIdx = len(newItem.Affixes)*4 + ifi.FirstProp
	}
	return newColIdx
}
func findUnusedColIdx(newItem *d2items.Item, startCol int, endCol int) int {
	for i := startCol; i <= endCol; i += 4 {
		for idx := range newItem.Affixes {
			if newItem.Affixes[idx].ColIdx == i {
				return i
			}
		}
	}
	return -1
}
