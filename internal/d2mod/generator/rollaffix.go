package generator

import (
	"log"

	"github.com/tlentz/d2modmaker/internal/d2mod/d2items"
	"github.com/tlentz/d2modmaker/internal/weightrand"
)

// RollAffix Randomly roll a new Prop for a given item
func RollAffix(g *Generator, item d2items.Item, sbm float32, colIdx int, targetPropScore int, w *weightrand.Weights) *d2items.Affix {
	for rollcounter := 0; rollcounter < 500; rollcounter++ {
		scoreFileRowIndex := w.Generate()
		l := g.psi.RowLines[scoreFileRowIndex]
		//log.Printf("RollAffix: %s: ScoreMax=%d T:%d TargetScore=%d\n", l.Prop.Name, l.ScoreMax, l.PropParType, targetPropScore)
		if l.ScoreMax < (targetPropScore / 8) {
			//log.Print("<Reroll S>")
			continue
		}
		if !d2items.CheckIETypes(g.TypeTree, item.Types[0], l.Itypes, l.Etypes) {
			//log.Print("<Reroll T>")
			continue
		}
		if item.Lvl < l.MinLvl {
			//log.Print("<Reroll L>")
			continue
		}
		newa := d2items.NewAffixFromLine(l, colIdx, sbm)
		return newa
	}
	log.Fatal("RollProp: Too many rolls (500).. couldn't find a a valid PropScore.txt line")
	panic(1)
}
