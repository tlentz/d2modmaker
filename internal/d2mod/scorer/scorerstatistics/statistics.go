package scorerstatistics

import (
	"github.com/tlentz/d2modmaker/internal/d2fs/filenumbers"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/propscores"
	"github.com/tlentz/d2modmaker/internal/d2mod/d2items"
	"github.com/tlentz/d2modmaker/internal/weightrand"
)

// ItemScoreMap  Map to quickly get a Score based on its Score.Item.Name
type ItemScoreMap map[string]int

// ItemTypeStatistics contains information from the Scorer about what it saw when scoring items for an item type.  Generator uses this to calculate probabilities.
type ItemTypeStatistics struct {
	ItemTypeIndex int                 // from the above constants, weaponWeightIndex, etc
	ItemType      string              //
	NumLines      map[int]int         // Records # of times a PropScore.txt line was seen that applied to an item, by row index
	Weights       *weightrand.Weights //
}

// ScorerStatistics Results from the Scorer
type ScorerStatistics struct {
	TypeStatistics [5]ItemTypeStatistics
	ItemScores     ItemScoreMap //
}

// NewScorerStatistics .
func NewScorerStatistics() *ScorerStatistics {
	ss := ScorerStatistics{}
	ss.ItemScores = ItemScoreMap{}
	for idx := range &ss.TypeStatistics {
		ss.TypeStatistics[idx].NumLines = map[int]int{}
	}
	return &ss
}

// ScoreLines Maps Prop name to list of Lines from propScores.txt
//var ScoreLines propscores.ScoreMap

// AddScoreLineWeight  Increases the weight for a specific prop on an item
func (ss *ScorerStatistics) AddScoreLineWeight(tt *d2items.TypeTree, item d2items.Item, line *propscores.Line) {

	its := ss.GetItemTypeStatistics(tt, item.FileNumber, item.Types[0])
	//fmt.Printf("Len NumLines = %d\n", len(its.NumLines))
	its.NumLines[line.RowIndex]++
}

// GetItemTypeStatistics Return the statistics for a given item type
func (ss *ScorerStatistics) GetItemTypeStatistics(tt *d2items.TypeTree, filenumber int, itemType string) *ItemTypeStatistics {
	if filenumber == filenumbers.Sets {
		//fmt.Println("GetWeights:set")
		return &ss.TypeStatistics[setWeightIndex]
	}
	if d2items.CheckTypeTree(tt, itemType, "armo") {
		//fmt.Println("GetWeights:armo")
		return &ss.TypeStatistics[armorWeightIndex]
	}
	if d2items.CheckTypeTree(tt, itemType, "weap") {
		//fmt.Println("GetWeights:weap")
		return &ss.TypeStatistics[weaponWeightIndex]
	}
	if d2items.CheckTypeTree(tt, itemType, "rin") || d2items.CheckTypeTree(tt, itemType, "amu") {
		//fmt.Println("GetWeights:rin")
		return &ss.TypeStatistics[jewelryWeightIndex]
	}
	// Quest items i.e. vip
	//fmt.Println("GetWeights:Unknown")
	return &ss.TypeStatistics[unknownWeightIndex]
}

/*
func (ss *ScorerStatistics) SetupProbabilityWeights() {
	for idx, w := range ss.TypeStatistics {
		weightedrarity := make([]int, len(s.Scorefile.Rows))
		total := 0
		for i, lenscores := 0, len(s.Scorefile.Rows); i < lenscores; i++ {
			// TODO: Change this *5 to opts.something.  Formula isn't good either, should
			//  calculate totals & normalize
			weightedrarity[i] = 1 + w.numLines[i]*5
			total += weightedrarity[i]
		}
		//w.weights = weightrand.NewWeights(weightedrarity)
		//w.weights.Generate()
		// Above 2 don't work, below 2 do work and I don't know why....!
		s.ScoreLineWeights[idx].weights = weightrand.NewWeights(weightedrarity)
		s.ScoreLineWeights[idx].weights.Generate()
	}
	for _, w := range s.scoreLineWeights {
		w.weights.Generate()
	}

}
*/
func (ss *ScorerStatistics) SetupProbabilityWeights() {
	// Klunky because indexing an array doesn't increase its size, i.e. just easier to use a map.
	// weightrand wants an array, so turn NumLines (map) into an array
	for ssidx := range ss.TypeStatistics {
		// find max key in NumLines
		maxidx := 0
		for rowIdx := range ss.TypeStatistics[ssidx].NumLines {
			if rowIdx > maxidx {
				maxidx = rowIdx
			}
		}
		//buckets = make([]int, maxidx)
		buckets := make([]int, maxidx+1)
		for rowIdx, count := range ss.TypeStatistics[ssidx].NumLines {
			buckets[rowIdx] += count
		}
		ss.TypeStatistics[ssidx].Weights = weightrand.NewWeights(buckets)
	}
}
