package d2items

import (
	"log"

	"github.com/tlentz/d2modmaker/internal/d2fs/txts/propscores"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/propscores/propscorespartype"
	"github.com/tlentz/d2modmaker/internal/d2mod/prop"
)

// checkPropScore Determine if a given prop matches a given line from PropScores.txt
// Checks (based on PropParType) Par, Min < Avg(Prop & Max) < Max and Itype/Etype restrictions
// Checks (based on PropParType) Par, Min < Avg(Prop & Max) < Max and Itype/Etype restrictions
// Checks (based on PropParType) Par, Min < Avg(Prop & Max) < Max and Itype/Etype restrictions
// Checks (based on PropParType) Par, Min < Avg(Prop & Max) < Max and Itype/Etype restrictions
func checkPropScore(tt *TypeTree, p prop.Prop, item Item, line *propscores.Line) bool {
	//fmt.Printf("checkPropScore: %s|%s>%s|%s\n", p.Name, p.Par, line.Prop.Name, line.Prop.Par)

	if item.Lvl < line.MinLvl {
		return false
	}
	switch line.PropParType {
	case propscorespartype.R, propscorespartype.Rp, propscorespartype.C, propscorespartype.Req:
		// LMin < Avg(Min,Max) < LMax
		avg := (p.Val.Min + p.Val.Max) / 2

		if (avg < line.Prop.Val.Min) || (avg > line.Prop.Val.Max) {
			return false
		}

	case propscorespartype.Rt:
		//
		if p.Par != line.Prop.Par {
			return false
		}
		avg := (p.Val.Min + p.Val.Max) / 2

		if (avg < line.Prop.Val.Min) || (avg > line.Prop.Val.Max) {
			return false
		}

	case propscorespartype.Lvl:
		// Par = N.  Min & Max are don't care

	case propscorespartype.S:
		// Par = LPar, Min = LMin, Max = LMax
		if p.Par != line.Prop.Par {
			return false
		}
		if p.Min != line.Prop.Min {
			return false
		}
		if p.Max != line.Prop.Max {
			return false
		}

	case propscorespartype.Scl, propscorespartype.Sch:
		// Par = LPar
		if p.Par != line.Prop.Par {
			return false
		}

	case propscorespartype.Smm:
		// Par = LPar, LMin < Avg(Min,Max) < LMax
		if p.Par != line.Prop.Par {
			return false
		}
		avg := (p.Val.Min + p.Val.Max) / 2
		//fmt.Printf("$%d", avg)
		if (avg < line.Prop.Val.Min) || (avg > line.Prop.Val.Max) {
			//fmt.Printf("%d %d %d", avg, line.Prop.Val.Min, line.Prop.Val.Max)
			return false
		}

	default:
		log.Fatalf("PropScores.txt has a line with and unknown PropParType")
	}
	//log.Println(item.Name)
	//log.Println(item.Types)
	return CheckIETypes(tt, item.Types, line.Itypes, line.Etypes)
}
