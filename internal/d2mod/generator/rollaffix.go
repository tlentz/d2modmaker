package generator

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/tlentz/d2modmaker/internal/d2fs/txts/propscorestxt/propscorespartype"
	"github.com/tlentz/d2modmaker/internal/d2mod/d2items"
	"github.com/tlentz/d2modmaker/internal/d2mod/propscores"
	"github.com/tlentz/d2modmaker/internal/util"
	"github.com/tlentz/d2modmaker/internal/weightrand"
)

const (
	// DeviationRatio specifies the + & - range difference between targetPropScore and the max targetPropScore that a randomly rolled prop/affix will have
	DeviationRatio = 0.1
)

// RollAffix Randomly roll a new Prop for a given item
// This assumes the targetPropScore has already been weighted by the SetBonusMultipler
func RollAffix(g *Generator, item *d2items.Item, colIdx int, propScoreMin int, targetPropScore int, propScoreMax int, w *weightrand.Weights) *d2items.Affix {
	line := rollPropScoreLine(g, item, colIdx, propScoreMin, targetPropScore, propScoreMax, w)
	newa := d2items.NewAffixFromLine(line, colIdx, g.IFI.FI.FileNumber)

	// if line.ScoreLimit > 0 {
	// 	//ScoreLimit specified verify this prop doesn't go over it.
	// 	vanillaScore := g.Statistics.ItemScores[item.Name]
	// 	if vanillaScore > 0 {
	// 		scoreCap := (vanillaScore * line.ScoreLimit) / 100
	// 		//log.Printf("RollAffix: found limit %s %s Van/Cap/Tgt %d/%d/%d ScoreLimit:%d", item.Name, line.Prop.Name, vanillaScore, scoreCap, targetPropScore, line.ScoreLimit)
	// 		if targetPropScore > scoreCap {
	// 			fmt.Printf("Capping %s:%s %d->%d\n", item.Name, line.Prop.Name, targetPropScore, scoreCap)
	// 		}

	// 		targetPropScore = util.MinInt(targetPropScore, scoreCap)
	// 		if targetPropScore < propScoreMin {
	// 			log.Panicf("RollAffix: Logic bug: targetPropScore (%d) < propScoreMin(%d) %s", targetPropScore, propScoreMin, line.Prop.Name)
	// 		}
	// 	}
	// }

	switch newa.Line.PropParType {
	case propscorespartype.R, propscorespartype.Rp, propscorespartype.Rt, propscorespartype.Smm, propscorespartype.C:
		//fmt.Printf("RollAffix: %d %d-%d -> ", targetPropScore, newa.P.Val.Min, newa.P.Val.Max)
		newa.P.Val.Min, newa.P.Val.Max = rollRange(g, newa.Line, item.Lvl, targetPropScore, true)
		newa.P.Min = strconv.Itoa(newa.P.Val.Min)
		newa.P.Max = strconv.Itoa(newa.P.Val.Max)
		if newa.P.Val.Min < newa.Line.Prop.Val.Min {
			log.Panicf("Min Out of range %s => %d < %d ", newa.Line.Prop.Name, newa.P.Val.Min, newa.Line.Prop.Val.Min)
		}
		if newa.P.Val.Max > newa.Line.Prop.Val.Max {
			log.Panicf("Min Out of range %s => %d < %d ", newa.Line.Prop.Name, newa.P.Val.Max, newa.Line.Prop.Val.Max)
		}
		//fmt.Printf("%s-%s\n", newa.P.Min, newa.P.Max)
	case propscorespartype.Req:
		newa.P.Val.Min, newa.P.Val.Max = rollRange(g, newa.Line, item.Lvl, targetPropScore, false)
		newa.P.Min = strconv.Itoa(newa.P.Val.Min)
		newa.P.Max = strconv.Itoa(newa.P.Val.Max)

	case propscorespartype.Lvl: // (pts or %)/lvl prop min & max are empty/ignored
		newa.P.Val.Par = rollMax(g.rng, newa.P.Val.Par, newa.Line, item.Lvl, targetPropScore)
		newa.P.Par = strconv.Itoa(newa.P.Val.Par)
	case propscorespartype.Scl: // Skill, %chance, Level
		// OBC: I'm too lazy to roll both chance and level as that makes the targetscore calculation
		// a lot more complex, so roll either Chance or Level.
		if g.rng.Intn(4) >= 2 {
			// Roll Chance
			newa.P.Val.Min = rollMax(g.rng, newa.P.Val.Min, newa.Line, item.Lvl, targetPropScore)
			newa.P.Min = strconv.Itoa(newa.P.Val.Min)
		} else {
			// Roll Level
			newa.P.Val.Max = rollMax(g.rng, newa.P.Val.Max, newa.Line, item.Lvl, targetPropScore)
			newa.P.Max = strconv.Itoa(newa.P.Val.Max)
		}
	case propscorespartype.Sch: // Skill, #charges, Level
		// Don't touch # charges... roll for Level
		newa.P.Val.Max = rollMax(g.rng, newa.P.Val.Max, newa.Line, item.Lvl, targetPropScore)
		newa.P.Max = strconv.Itoa(newa.P.Val.Max)
	case propscorespartype.S:
		// Nothing to adjust here
	default:
		log.Fatalf("calcPropScore: Unhandled Prop type %s:%d\n", newa.P.Name, newa.Line.PropParType)

	}
	// if line.Prop.Name == "dmg-fire" || line.Prop.Name == "dmg-cold" {
	// 	fmt.Printf("%+v\n", newa.P)
	// }
	return newa
}

// rollRange:  Calculate a range for a given targetScore and Line.  If applyDeviation a
func rollRange(g *Generator, line *propscores.Line /*min int, max int, scoreMin int, scoreMax int, */, itemLvl int, targetPropScore int, applyDeviation bool) (int, int) {
	scoreCap := calcScoreCap(g, line, itemLvl)
	if line.ScoreMin < 0 {
		if targetPropScore < scoreCap {
			targetPropScore = scoreCap
		}
	} else {
		if targetPropScore > scoreCap {
			targetPropScore = scoreCap
		}
	}
	newAvg := util.Interpolate(targetPropScore, targetPropScore, line.ScoreMin, line.ScoreMax, line.Prop.Val.Min, line.Prop.Val.Max)
	//newMin := util.Interpolate(targetPropScore, targetPropScore, scoreMin, scoreMax, min, max)
	//fmt.Printf("rollRange: %d %d-%d -> %d %d\n", targetPropScore, line.ScoreMin, line.ScoreMax, min, max)

	deviation := 0
	if applyDeviation {
		deviation = util.AbsInt(util.Round64((g.rng.NormFloat64()*0.10 + .05) * float64(newAvg)))
	}
	newMin := newAvg - deviation
	newMax := newAvg + deviation

	roundTo := 1
	largest := util.MaxInt(util.AbsInt(newMin), util.AbsInt(newMax))
	if largest >= 15 {
		roundTo = 5
	}
	if largest >= 50 {
		roundTo = 10
	}
	newMin = util.Round32(float32(newMin)/float32(roundTo)) * roundTo
	newMax = util.Round32(float32(newMax)/float32(roundTo)) * roundTo

	newMin = util.MaxInt(newMin, line.Prop.Val.Min)
	newMin = util.MinInt(newMin, line.Prop.Val.Max)
	newMax = util.MinInt(newMax, line.Prop.Val.Max)
	newMax = util.MaxInt(newMax, line.Prop.Val.Min)

	return newMin, newMax
}

// rollMax:  Roll for a maximum value within 20% (PropVariance = 0.2) of the targetPropScore
func rollMax(rng *rand.Rand, max int, line *propscores.Line, itemLvl int, targetPropScore int) int {
	if max < 0 {
		log.Fatalf("rollMax: upper limit cannot be negative: Prop:%s: %d", line.Prop.Name, max)
	}
	if targetPropScore > util.Round32(float32(line.ScoreMax)*(1.0-DeviationRatio)) {
		return max
	}
	newMax := util.Interpolate(targetPropScore, targetPropScore, line.ScoreMin, line.ScoreMax, 0, max)
	targetvariance := float32(newMax) * DeviationRatio
	newMax += util.Round32((rng.Float32() * targetvariance * 2) - targetvariance)
	if newMax > max {
		newMax = line.Prop.Val.Max
	}
	if newMax < 0 {
		// 0% str per level would have no effect at all  If score ends up being too high then the affix will ge rerolled
		newMax = 1
	}
	return newMax
}

// rollPropScoreLine Find a line in PropScores.txt where Line.ScoreMin < targetPropScore < Line.ScoreMax
func rollPropScoreLine(g *Generator, item *d2items.Item, colIdx int, propScoreMin int, targetPropScore int, propScoreMax int, w *weightrand.Weights) *propscores.Line {

	// These tests don't work on sets
	// if targetPropScore < propScoreMin {
	// 	log.Panic("targetPropScore not > propScoreMin")
	// }
	// if targetPropScore > propScoreMax {
	// 	log.Panic("targetPropScore not < propScoreMax")
	// }

	var closestLine *propscores.Line
	var closestLineDelta = 999999999 // couldn't get a portable answer for maximum value of int.
	rollcounter := 0
	maxRolls := 200
	vanillaScore := g.Statistics.ItemScores[item.Name]
DoneRolling:
	for rollcounter = 0; rollcounter < maxRolls; rollcounter++ {

		scoreFileRowIndex := w.Generate(g.rng)
		line := g.psi.RowLines[scoreFileRowIndex]
		//log.Printf("rollPropScoreLine: %s: ScoreMax=%d T:%d Score Min/Tgt/Max:%d/%d/%d %t\n", line.Prop.Name, line.ScoreMax, line.PropParType, propScoreMin, targetPropScore, propScoreMax, line.LvlScale)
		if !d2items.CheckIETypes(g.TypeTree, item.Types, line.Itypes, line.Etypes) {
			continue
		}
		if item.Lvl < line.MinLvl {
			continue
		}
		if checkGroups(line.Group, item) {
			continue
		}
		if line.ScoreLimit == 0 {
			continue // alpha-13 Special case where we want to generate a score but not ever generate the prop
		}

		if closestLine == nil {
			closestLine = line
		}

		if line.Prop.Name[0] == '*' {
			// Skipping commented out props, i.e. meleesplash
			continue
		}
		// var lineDelta int
		//log.Printf("rollPropScoreLine: %s Delta: %d <%d %d %d> %t %d %d\n", item.Name, closestLineDelta, line.ScoreMin, targetPropScore, line.ScoreMax, line.LvlScale, propScoreMin, propScoreMax)
		// Adjust for cases where Min < Max
		scoreMin := line.ScoreMin
		scoreMax := line.ScoreMax
		if line.ScoreMax < line.ScoreMin {
			scoreMax = line.ScoreMin
			scoreMin = line.ScoreMax
		}

		if line.LvlScale {
			scoreMin = util.Round32(float32(scoreMin) * float32(item.Lvl) / 50.0)
			scoreMax = util.Round32(float32(scoreMax) * float32(item.Lvl) / 50.0)
		}
		// TODO: SynergyBonus calculation

		if line.ScoreLimit > 0 {
			//ScoreLimit specified verify this prop doesn't go over it.
			scoreCap := (vanillaScore * line.ScoreLimit) / 100
			if scoreCap < scoreMin {
				//log.Printf("rollPropScoreLine Skipping %s:%s Tgt/Cap/Min %d/%d/%d Line Score Min/Max (%d/%d)", item.Name, line.Prop.Name, targetPropScore, scoreCap, scoreMin, line.ScoreMin, line.ScoreMax)
				continue
			}
			// if line.Prop.Name == "str" {
			// 	log.Printf("rollPropScoreLine Limiting %s:%s Min/Tgt/Max %d/%d/%d to Min/Max/Cap (%d/%d/%d)", item.Name, line.Prop.Name, propScoreMin, targetPropScore, propScoreMax, scoreMin, scoreMax, scoreCap)
			// }
			scoreMax = util.MinInt(scoreMax, scoreCap)
		}
		scoreCap := calcScoreCap(g, line, item.Lvl)
		if scoreCap < scoreMax {
			scoreMax = scoreCap
		}

		switch {
		case propScoreMin > scoreMax:
			lineDelta := propScoreMin - scoreMax
			if lineDelta < closestLineDelta {
				closestLine = line
				closestLineDelta = lineDelta
			}
		case propScoreMax < scoreMin:
			lineDelta := scoreMin - propScoreMax
			if lineDelta < closestLineDelta {
				closestLine = line
				closestLineDelta = lineDelta
			}
		default:
			closestLine = line
			closestLineDelta = 0
			//log.Printf("rollPropScoreLine: Found on %d rolls: %s %d/%d/%d\n", rollcounter+1, line.Prop.Name, scoreMin, targetPropScore, scoreMax)
			break DoneRolling
		}
	}
	if closestLineDelta != 0 {
		//log.Println(item.Affixes)
		//log.Printf("rollPropScoreLine: Couldn't hit target: %s: %d %d", item.Name, closestLineDelta, targetPropScore)
	}
	// if rollcounter == MaxRolls {
	// 	if closestLineDelta > 100 {
	// 		log.Printf("Hit maxrolls trying to roll %d Closest:%d ", targetPropScore, closestLineDelta)
	// 		log.Printf("closestLine:%+v\n\n", closestLine)
	// 		log.Printf("Item:%+v", item)
	// 		panic(1)
	// 	}
	//}
	g.numAffixRolls += rollcounter
	return closestLine
}

// calcScoreCap Returns the highest (or lowest if negative) score allowed by the line for an item.
func calcScoreCap(g *Generator, line *propscores.Line, itemLvl int) int {
	// level based limit is always from absolute valued smallest to absolute valued largest
	scoreMin := line.ScoreMin
	scoreMax := line.ScoreMax
	if line.ScoreMax < line.ScoreMin {
		scoreMax = line.ScoreMin
		scoreMin = line.ScoreMax
	}
	if (line.ScoreMax < 0) || (line.ScoreMin < 0) {
		scoreMax, scoreMin = scoreMin, scoreMax
	}
	scoreCap := scoreMax

	if g.opts.PropScoreMultiplier < 4 {
		// alpha-17, add level based prop scaling
		lMax := line.MaxLvl
		if lMax == 0 {
			lMax = util.MaxInt(70, line.MinLvl)
		}
		// Scale item level up based on PropScaleMultiplier
		scaledItemLevel := util.MinInt(util.Round64(float64(itemLvl)*(1.2+0.1*(g.opts.PropScoreMultiplier-1))), lMax)
		if scaledItemLevel < line.MinLvl || scaledItemLevel > lMax {
			log.Panicf("calcScoreCap: scaledItemLevel not within limits %+v", line)
		}
		scoreCap = util.Interpolate(scaledItemLevel, scaledItemLevel, line.MinLvl, lMax, scoreMin, scoreMax)
		if scoreCap > scoreMax && (scoreMax > 0) {
			log.Panicf("scoreCap %d > scoreMax %d", scoreCap, scoreMax) // assertion
		}
		// if scoreMin < 0 && (scaledItemLevel < 70) {
		// 	log.Printf("calcScoreCap: %d (%d %d %d)", item.Lvl, scoreMin, scoreCap, scoreMax)
		// }
		// if (line.Prop.Name == "att") && (item.Lvl < 20) {
		// 	log.Printf("calcScoreCap Limiting %s:%s Min/Tgt/Max %d/%d/%d to Min/Max/Cap (%d/%d/%d)", item.Name, line.Prop.Name, propScoreMin, targetPropScore, propScoreMax, scoreMin, scoreMax, scoreCap)
		// }

	}
	return scoreCap

}

// checkGroups Returns true if any Affix on item is a member of the same group
func checkGroups(group string, item *d2items.Item) bool {
	if group == "" {
		log.Fatalf("Empty PropScores.txt Group")
	}
	for _, a := range item.Affixes {
		if a.Line.Group == "" {
			log.Fatalf("Empty PropScores.txt Group")
		}
		if a.Line.Group == group {
			return true
		}
	}
	return false
}
