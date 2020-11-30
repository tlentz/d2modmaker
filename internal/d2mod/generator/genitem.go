package generator

import (
	"log"
	"math/rand"

	"github.com/tlentz/d2modmaker/internal/d2mod/d2items"
	"github.com/tlentz/d2modmaker/internal/d2mod/scorer"
	"github.com/tlentz/d2modmaker/internal/util"
)

// GenItem returns a new item with random affixes based on supplied item
func GenItem(g *Generator, item *d2items.Item) *d2items.Item {
	if g.Statistics == nil {
		log.Fatalln("GenItems: generator statistics was not initialized")
	}
	its := g.Statistics.GetItemTypeStatistics(g.TypeTree, item.FileNumber, item.Types[0]) // FIXME: This is broken for runewords that work in both weapons & armor...
	if len(its.NumLines) == 0 {
		log.Panic("item statistics structure empty")
	}
	its.Weights.Generate()
	targetScore := util.Round32(float32(g.Statistics.ItemScores[item.Name]) * float32(g.opts.PropScoreMultiplier))
	targetPropCount := 0
	if g.opts.BalancedPropCount {
		targetPropCount = len(item.Affixes)
	} else {
		if g.IFI.NumProps == 0 {
			log.Panic("g.IFI.NumProps == 0")
		}
		minProps := util.MinInt(g.opts.MinProps, g.IFI.NumProps)
		minProps = util.MaxInt(minProps, 1) // Because everyone wants some props
		maxProps := util.MinInt(g.opts.MaxProps, g.IFI.NumProps)

		targetPropCount = minProps + rand.Intn(maxProps-minProps+1) // beware how Intn behaves...
		targetPropCount = util.MinInt(targetPropCount, g.IFI.NumProps)
	}
	newi := item.CloneWithoutAffixes()
	newi.Affixes = []d2items.Affix{}
	rollCount := 0
	itemScore := 0
	minItemScore := util.MaxInt(10, util.Round32(float32(targetScore)*0.6))
	maxItemScore := util.MaxInt(30, util.Round32(float32(targetScore)*1.4))
	for (itemScore < minItemScore) || (itemScore > maxItemScore) || (len(newi.Affixes) < targetPropCount) || (len(newi.Affixes) > targetPropCount+3) {
		// Roll/Reroll entire item
		newi.Affixes = []d2items.Affix{}
		itemScore = 0
		//log.Printf("genItem: # Target Prop Count %d, Scores(min/tgt/max): %d/%d/%d", targetPropCount, minItemScore, targetScore, maxItemScore)
		if targetPropCount == 0 {
			log.Printf("genItem: # Item Affixes=%d, Opts.MinProps=%d, Opts.MaxProps=%d, File NumProps=%d", len(item.Affixes), g.opts.MinProps, g.opts.MaxProps, g.IFI.NumProps)
			log.Panic("genItem: 0 targetPropCount")
		}
		// TODO: Add better partial set bonus support
		//for (itemScore < int((float32(targetScore) * 0.8))) && (len(newi.Affixes) < targetPropCount) && (len(newi.Affixes) < g.IFI.NumProps) {
		for len(newi.Affixes) < targetPropCount {

			targetPropScore := util.MaxInt(0, targetScore-itemScore)
			if len(newi.Affixes) < (targetPropCount - 1) {
				targetPropScore = util.MaxInt(0, int(float32((targetScore-itemScore))*0.3))
			}
			//log.Printf("genItem: #Affixes: %d TargetScore: %d TargetPropScore: %d", len(newi.Affixes), targetScore, targetPropScore)

			newColIdx := len(newi.Affixes)*4 + g.IFI.FirstProp
			sbm := d2items.CalcSetBonusMultiplier(g.IFI.FI.FileNumber, newColIdx)
			newAffix := RollAffix(g, *item, sbm, newColIdx, targetPropScore, its.Weights)
			newi.Affixes = append(newi.Affixes, *newAffix)
			oldScore := itemScore
			itemScore = scorer.ScoreItem(g.Statistics, g.TypeTree, *newi)
			reroll := false
			if ((itemScore * 2) > ((targetScore + 10) * 3)) && ((itemScore - oldScore) > 10) {
				// Over budget and new item score is > 10:  reroll
				//log.Printf("<Reroll Sc %d %d>", itemScore*2, ((targetScore + 10) * 4))
				reroll = true
			}
			if itemScore > maxItemScore {
				reroll = true
			}
			for i := 0; i < len(newi.Affixes)-1; i++ {
				if newi.Affixes[i].Line.Group == newAffix.Line.Group {
					//log.Printf("<Reroll Grp>")
					reroll = true
				}
			}
			if reroll {
				// Overshot, remove last prop so we can re-roll
				newi.Affixes = newi.Affixes[:len(newi.Affixes)-1]
				itemScore = oldScore
			}
			rollCount++
			if rollCount > 400 {
				log.Fatalf("genItem: > 400 rolls for %s", item.Name)
				return newi
			}
			/*
				// Debugging
				lastAffixName := ""
				if len(newi.Affixes) > 0 {
					lastAffixName = newi.Affixes[len(newi.Affixes)-1].P.Name
				}
				log.Printf("genItem: Count(itm/max) %d/%d, Scores(min/itm/tgt/max): %d/%d/%d/%d, reroll:%t, last aff=%s", len(newi.Affixes), targetPropCount, minItemScore, itemScore, targetScore, maxItemScore, reroll, lastAffixName)
			*/
		}
	}
	// TODO: Write these statistics out to a file (Target score, Generated item score, # props Vanilla vs generated)
	//fmt.Printf("Item\t%s\t%d\t%d\t%d\n", newi.Name, newi.FileNumber, targetScore, itemScore)
	return newi
}
