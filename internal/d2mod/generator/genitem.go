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
	targetScore := g.Statistics.ItemScores[item.Name]
	targetPropCount := 0
	if g.opts.BalancedPropCount {
		targetPropCount = len(item.Affixes)
	} else {
		if g.IFI.NumProps == 0 {
			log.Panic("g.IFI.NumProps == 0")
		}
		minProps := util.MinInt(g.opts.MinProps, g.IFI.NumProps)
		maxProps := util.MinInt(g.opts.MaxProps, g.IFI.NumProps)

		targetPropCount = minProps + rand.Intn(maxProps-minProps+1) // beware how Intn behaves...
		targetPropCount = util.MinInt(targetPropCount, g.IFI.NumProps)
	}
	newi := item.CloneWithoutAffixes()
	newi.Affixes = []d2items.Affix{}
	itemScore := 0
	rollCount := 0
	//log.Printf("# Item Affixes=%d, Opts.MinProps=%d, Ops.MaxProps=%d, File NumProps=%d", len(item.Affixes), g.opts.MinProps, g.opts.MaxProps, g.IFI.NumProps)
	if targetPropCount == 0 {
		log.Printf("# Item Affixes=%d, Opts.MinProps=%d, Ops.MaxProps=%d, File NumProps=%d", len(item.Affixes), g.opts.MinProps, g.opts.MaxProps, g.IFI.NumProps)
		log.Panic("0 targetPropCount")
	}
	// TODO: Add better partial set bonus support
	for (itemScore < int((float32(targetScore) * 0.8))) && (len(newi.Affixes) < targetPropCount) && (len(newi.Affixes) < g.IFI.NumProps) {
		targetPropScore := util.MaxInt(0, int(float32((targetScore-itemScore))*0.6))
		//log.Printf("GenItem TargetScore: %d TargetPropScore: %d", targetScore, targetPropScore)
		if len(newi.Affixes) >= (targetPropCount - 1) {
			targetPropScore = util.MaxInt(0, targetScore-itemScore)
		}

		newColIdx := len(newi.Affixes)*4 + g.IFI.FirstProp
		sbm := d2items.CalcSetBonusMultiplier(g.IFI.FI.FileNumber, newColIdx)
		newAffix := RollAffix(g, *item, sbm, newColIdx, targetPropScore, its.Weights)
		newi.Affixes = append(newi.Affixes, *newAffix)
		oldScore := itemScore
		itemScore = scorer.ScoreItem(g.Statistics, g.TypeTree, *newi)
		reroll := false
		if (itemScore*2) > ((targetScore+10)*4) && ((itemScore - oldScore) > 10) {
			// Over budget and new item score is > 10:  reroll
			//log.Printf("<Reroll Sc %d %d>", itemScore*2, ((targetScore + 10) * 4))
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
		if rollCount > 200 {
			log.Fatal("> 200 rolls")
			return newi
		}
	}

	return newi
}
