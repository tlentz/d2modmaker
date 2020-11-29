package d2items

import (
	"github.com/tlentz/d2modmaker/internal/d2fs/filenumbers"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/setItems"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/sets"
)

// CalcSetBonusMultiplier Set Bonuses are valued lower than native item props. Returns multiplier to scale the
// Prop score based on how many items it takes to trigger the set bonus.
func CalcSetBonusMultiplier(filenumber int /* item Item,*/, colIndex int) float32 {
	sbn := 0
	if filenumber == filenumbers.Sets {
		switch colIndex {
		case sets.PCode2a, sets.PCode2b:
			sbn = 2
		case sets.PCode3a, sets.PCode3b:
			sbn = 3
		case sets.PCode4a, sets.PCode4b:
			sbn = 4
		case sets.PCode5a, sets.PCode5b:
			sbn = 5
		default:
			sbn = -1 // FIXME: Don't know # of items in full set, using -1 to indicate full set bonus
		}

	}
	if filenumber == filenumbers.SetItems {
		// It's too hard to track AddFunc.  Ignore since it's only on Civerbs
		//if (row[setItems.AddFunc] == "1") || (row[setItems.AddFunc] == "2") {
		switch colIndex {
		case setItems.AProp1a, setItems.AProp1b:
			sbn = 2
		case setItems.AProp2a, setItems.AProp2b:
			sbn = 3
		case setItems.AProp3a, setItems.AProp3b:
			sbn = 4
		case setItems.AProp4a, setItems.AProp4b:
			sbn = 5
		case setItems.AProp5a, setItems.AProp5b:
			sbn = 6
		}
		//}
	}

	// Apply pset/fset bonus nerf
	sbmult := float32(1.0)
	switch sbn {
	case 0:
		sbmult = 1
	case 2:
		sbmult = 0.5
	case 3:
		sbmult = 0.4
	case 4:
		sbmult = 0.30
	case 5:
		sbmult = 0.25
	case 6:
		sbmult = 0.20
	case -1: // full set bonus
		sbmult = 0.3 // hack since the # items in full set is not known
	}
	return sbmult
}
