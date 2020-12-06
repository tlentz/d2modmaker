package generator

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/tlentz/d2modmaker/internal/d2fs/txts/propscores"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/propscores/propscorespartype"
	"github.com/tlentz/d2modmaker/internal/d2mod/d2items"
	"github.com/tlentz/d2modmaker/internal/util"
	"github.com/tlentz/d2modmaker/internal/weightrand"
)

const (
	// PropVariance controls the range between min & max that a randomly rolled prop/affix will have
	PropVariance = 0.2
)

// RollAffix Randomly roll a new Prop for a given item
// This assumes the targetPropScore has already been weighted by the SetBonusMultipler
func RollAffix(g *Generator, item *d2items.Item, colIdx int, targetPropScore int, w *weightrand.Weights) *d2items.Affix {
	line := rollPropScoreLine(g, item, colIdx, targetPropScore, w)
	newa := d2items.NewAffixFromLine(line, colIdx, g.IFI.FI.FileNumber)

	switch newa.Line.PropParType {
	case propscorespartype.R, propscorespartype.Rp, propscorespartype.Rt, propscorespartype.Smm, propscorespartype.C:
		//fmt.Printf("RollAffix: %d %d-%d -> ", targetPropScore, newa.P.Val.Min, newa.P.Val.Max)
		newa.P.Val.Min, newa.P.Val.Max = rollRange(newa.P.Val.Min, newa.P.Val.Max, newa.Line.ScoreMin, newa.Line.ScoreMax, item.Lvl, targetPropScore, true)
		newa.P.Min = strconv.Itoa(newa.P.Val.Min)
		newa.P.Max = strconv.Itoa(newa.P.Val.Max)
		//fmt.Printf("%s-%s\n", newa.P.Min, newa.P.Max)
	case propscorespartype.Req:
		newa.P.Val.Min, newa.P.Val.Max = rollRange(newa.P.Val.Min, newa.P.Val.Max, newa.Line.ScoreMin, newa.Line.ScoreMax, item.Lvl, targetPropScore, false)
		newa.P.Min = strconv.Itoa(newa.P.Val.Min)
		newa.P.Max = strconv.Itoa(newa.P.Val.Max)

	case propscorespartype.Lvl: // (pts or %)/lvl prop min & max are empty/ignored
		newa.P.Val.Par = rollMax(newa.P.Val.Par, newa.Line, item.Lvl, targetPropScore)
		newa.P.Par = strconv.Itoa(newa.P.Val.Par)
	case propscorespartype.Scl: // Skill, %chance, Level
		// OBC: I'm too lazy to roll both chance and level as that makes the targetscore calculation
		// a lot more complex, so roll either Chance or Level.
		if rand.Intn(4) >= 2 {
			// Roll Chance
			newa.P.Val.Min = rollMax(newa.P.Val.Min, newa.Line, item.Lvl, targetPropScore)
			newa.P.Min = strconv.Itoa(newa.P.Val.Min)
		} else {
			// Roll Level
			newa.P.Val.Max = rollMax(newa.P.Val.Max, newa.Line, item.Lvl, targetPropScore)
			newa.P.Max = strconv.Itoa(newa.P.Val.Max)
		}
	case propscorespartype.Sch: // Skill, #charges, Level
		// Don't touch # charges... roll for Level
		newa.P.Val.Max = rollMax(newa.P.Val.Max, newa.Line, item.Lvl, targetPropScore)
		newa.P.Max = strconv.Itoa(newa.P.Val.Max)
	case propscorespartype.S:
		// Nothing to adjust here
	default:
		log.Fatalf("calcPropScore: Unhandled Prop type %s:%d\n", newa.P.Name, newa.Line.PropParType)

	}
	return newa
}

// rollRange:  Calculate a min & max value +/- PropVariance centered around targetPropScore.
func rollRange(min int, max int, scoreMin int, scoreMax int, itemLvl int, targetPropScore int, applyVariance bool) (int, int) {
	if min > max {
		log.Fatalf("rollRange: min > max: (%d %d)", min, max)
	}
	newAvg := util.Interpolate(targetPropScore, targetPropScore, scoreMin, scoreMax, min, max)
	//fmt.Printf("rollRange: %d %d-%d -> %d %d\n", targetPropScore, line.ScoreMin, line.ScoreMax, min, max)
	variance := 0
	if applyVariance {
		variance = util.Round32(float32((max - min)) * (PropVariance / 2.0))
	}
	newMax := newAvg + variance
	newMin := newAvg - variance
	roundTo := 1
	largest := util.MaxInt(util.AbsInt(max), util.AbsInt(min))
	if largest >= 30 {
		roundTo = 5
	}
	if largest >= 70 {
		roundTo = 10
	}
	newMin = util.Round32(float32(newMin)/float32(roundTo)) * roundTo
	newMax = util.Round32(float32(newMax)/float32(roundTo)) * roundTo

	newMin = util.MaxInt(newMin, min)
	newMin = util.MinInt(newMin, max)
	newMax = util.MinInt(newMax, max)
	newMax = util.MaxInt(newMax, min)

	return newMin, newMax
}

// rollMax:  Roll for a maximum value within 20% (PropVariance = 0.2) of the targetPropScore
func rollMax(max int, line *propscores.Line, itemLvl int, targetPropScore int) int {
	if max < 0 {
		log.Fatalf("rollMax: upper limit cannot be negative: Prop:%s: %d", line.Prop.Name, max)
	}
	if targetPropScore > util.Round32(float32(line.ScoreMax)*(1.0-PropVariance)) {
		return max
	}
	newMax := util.Interpolate(targetPropScore, targetPropScore, line.ScoreMin, line.ScoreMax, 0, max)
	targetvariance := float32(newMax) * PropVariance
	newMax += util.Round32((rand.Float32() * targetvariance * 2) - targetvariance)
	if newMax > max {
		newMax = line.Prop.Val.Max
	}
	if newMax < 0 {
		// 0% str per level would have no effect at all  If score ends up being too high then the affix will ge rerolled
		newMax = 1
	}
	return newMax
}

// rollPropScoreLine Find a line in PropScores.txt where Line.ScoreMin < targetLineScore < Line.ScoreMax
func rollPropScoreLine(g *Generator, item *d2items.Item, colIdx int, targetLineScore int, w *weightrand.Weights) *propscores.Line {
	var closestLine *propscores.Line
	var closestLineDelta = 999999999 // couldn't get a portable answer for maximum value of int.
	rollcounter := 0
	for rollcounter = 0; rollcounter < 100; rollcounter++ {
		scoreFileRowIndex := w.Generate()
		line := g.psi.RowLines[scoreFileRowIndex]
		//log.Printf("rollPropScoreLine: %s: ScoreMax=%d T:%d TargetScore=%d\n", line.Prop.Name, line.ScoreMax, line.PropParType, targetPropScore)
		if !d2items.CheckIETypes(g.TypeTree, item.Types, line.Itypes, line.Etypes) {
			//log.Print("<Reroll T>")
			continue
		}
		if item.Lvl < line.MinLvl {
			//log.Print("<Reroll L>")
			continue
		}
		if checkGroups(line.Group, item) {
			line = nil
			continue
		}

		if closestLine == nil {
			closestLine = line
		}
		// var lineDelta int
		//fmt.Printf("%s %d <%d %d %d>\n", item.Name, closestLineDelta, line.ScoreMin, targetLineScore, line.ScoreMax)
		// Adjust for cases where Min < Max
		scoreMin := line.ScoreMin
		scoreMax := line.ScoreMax
		if line.ScoreMax < line.ScoreMin {
			scoreMax = line.ScoreMin
			scoreMin = line.ScoreMax
		}
		switch {
		case targetLineScore >= scoreMax:
			lineDelta := targetLineScore - scoreMax
			if lineDelta < closestLineDelta {
				closestLine = line
				closestLineDelta = lineDelta
			}
		case targetLineScore <= scoreMin:
			lineDelta := scoreMin - targetLineScore
			if lineDelta < closestLineDelta {
				closestLine = line
				closestLineDelta = lineDelta
			}
		default:
			closestLine = line
			closestLineDelta = 0
			break
		}
	}
	if closestLineDelta > 0 {
		//log.Println(item.Affixes)
		//log.Printf("rollPropScoreLine: Couldn't hit target: %s: %d %d", item.Name, closestLineDelta, targetLineScore)

	}
	g.numAffixRolls += rollcounter
	return closestLine
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
