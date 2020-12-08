package scorer

import (
	"log"

	"github.com/tlentz/d2modmaker/internal/d2fs/txts/propscores"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/propscores/propscorespartype"
	"github.com/tlentz/d2modmaker/internal/d2mod/d2items"
	"github.com/tlentz/d2modmaker/internal/d2mod/prop"
	"github.com/tlentz/d2modmaker/internal/d2mod/scorer/scorerstatistics"
	"github.com/tlentz/d2modmaker/internal/util"
)

// scoreProp returns
func scoreProp(ss *scorerstatistics.ScorerStatistics, tt *d2items.TypeTree, item d2items.Item, p prop.Prop, l *propscores.Line) int {
	//fmt.Printf(" %s/%s/%s/%s\n", p.Name, p.Par, p.Min, p.Max)

	if p.Name[0] == '*' { // Blizz comments out with *
		return 0
	}
	score := calcPropScore(p, l)
	if l.LvlScale {
		score = util.Round32(float32(score) * float32(item.Lvl) / 50.0)
	}
	return score
	/*
			for _, line := range scorelines[p.Name] {
				if checkPropScore(tt, p, item, line) {
					//s.Group[p] = line.Group
					//s.SynergyGroup[p] = line.SynergyGroup
					return calcPropScore(p, line), &line
				}

			}
		for _, l := range scorelines[p.Name] {
			//fmt.Printf("L%d - %s|%s|%s\n", ln, l.Prop.Name, l.Prop.Par)
			d2items.PrintProp(l.Prop)
		}

		log.Fatalf("ScoreProp: Couldn't Score %s<%d>: %s|%s|%s|%s\n", item.Name, item.Lvl, p.Name, p.Par, p.Min, p.Max)
		return 0, nil
	*/
}

func calcPropScore(p prop.Prop, line *propscores.Line) int {
	score := 0
	//log.Printf("CalcPropScore %s (/lvl)->%d, %d|%d|%d|%d|%d|%d\n", line.Prop.Name, score, p.Val.Par, p.Val.Par, 0, line.Prop.Val.Par, line.ScoreMin, line.ScoreMax)
	switch line.PropParType {
	case propscorespartype.R, propscorespartype.Req, propscorespartype.Rp, propscorespartype.Rt, propscorespartype.Smm, propscorespartype.C:
		score = util.Interpolate(p.Val.Min, p.Val.Max, line.Prop.Val.Min, line.Prop.Val.Max, line.ScoreMin, line.ScoreMax)
		// if line.Prop.Name == "ac%" {
		// 	log.Printf("%s:%d/%d/%d/%d %d \n %+v", line.Prop.Name, p.Val.Min, p.Val.Max, line.ScoreMin, line.ScoreMax, score, line)
		// }
	case propscorespartype.Lvl: // (pts or %)/lvl prop min & max are empty/ignored
		if line.Prop.Val.Par == 0 {
			log.Fatalf("calcPropScore: PropsScore.txt prop %s has 0 for Par", line.Prop.Name)
		}
		score = util.Interpolate(p.Val.Par, p.Val.Par, 0, line.Prop.Val.Par, line.ScoreMin, line.ScoreMax)
	case propscorespartype.Scl: // Skill, %chance, Level
		if p.Par != line.Prop.Par {
			log.Fatalf("calcPropScore: SCL par mismatch %s %s <> %s\n", p.Name, p.Par, line.Prop.Par)
		}
		v := p.Val.Min + (p.Val.Max * 100)
		lmax := line.Prop.Val.Min + (line.Prop.Val.Max * 100)
		score = util.Interpolate(v, v, 0, lmax, line.ScoreMin, line.ScoreMax)
	case propscorespartype.Sch: // Skill, #charges, Level
		score = util.Interpolate(p.Val.Max, p.Val.Max, 0, line.Prop.Val.Max, line.ScoreMin, line.ScoreMax)
	case propscorespartype.S:
		score = line.ScoreMax
	default:
		log.Fatalf("calcPropScore: Can't score type %s:%d\n", p.Name, line.PropParType)

	}
	return score
}
