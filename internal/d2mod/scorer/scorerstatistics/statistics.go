package scorerstatistics

import (
	"log"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2mod/d2items"
	"github.com/tlentz/d2modmaker/internal/d2mod/propscores"
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
type pBucketMap map[string]string

// ScorerStatistics Results from the Scorer
type ScorerStatistics struct {
	TypeStatistics []ItemTypeStatistics

	pBucket        *pBucketMap // From Armor[Type] & Weapons[Type], : Maps membership of item into groups for doing weighted probability calc
	pBucketIndexes map[string]int

	ItemScores ItemScoreMap // Vanilla item scores, used as target scores

}

// NewScorerStatistics .
func NewScorerStatistics(d2files *d2fs.Files) *ScorerStatistics {
	ss := ScorerStatistics{
		pBucket:        newpBucketMap(d2files),
		pBucketIndexes: make(map[string]int),
		ItemScores:     ItemScoreMap{},
	}

	nextBucketIndex := 1
	for _, p := range *ss.pBucket {
		if ss.pBucketIndexes[p] == 0 {
			ss.pBucketIndexes[p] = nextBucketIndex
			nextBucketIndex++
		}
	}
	//fmt.Printf("%+v\n", ss.pBucket)
	ss.TypeStatistics = make([]ItemTypeStatistics, nextBucketIndex)
	for idx := range ss.TypeStatistics {
		ss.TypeStatistics[idx].NumLines = map[int]int{}
	}

	return &ss
}

// ScoreLines Maps Prop name to list of Lines from propScores.txt
//var ScoreLines propscores.ScoreMap

// AddScoreLineWeight  Increases the weight for a specific prop on an item
func (ss *ScorerStatistics) AddScoreLineWeight(tt *d2items.TypeTree, item *d2items.Item, line *propscores.Line) {

	its := ss.GetItemTypeStatistics(item)
	//fmt.Printf("Len NumLines = %d\n", len(its.NumLines))
	its.NumLines[line.RowIndex]++
}

// GetItemTypeStatistics Return the statistics for a given item type
func (ss *ScorerStatistics) GetItemTypeStatistics(item *d2items.Item) *ItemTypeStatistics {
	bucketName := (*ss.pBucket)[item.Code]
	if bucketName == "" {
		log.Panicf("Unknown item type %s, check PBucketList.txt\n%+v", item.Code, ss.pBucket)
	}
	bucketIndex := ss.pBucketIndexes[bucketName]
	return &ss.TypeStatistics[bucketIndex]
}

// GetBucketName Gets name of the bucket that is used for probability based weighting
func (ss *ScorerStatistics) GetBucketName(item *d2items.Item) string {
	return (*ss.pBucket)[item.Code]
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

// SetupProbabilityWeights Set up weighted proability array so for the weighted random roller.
// Each line in PropScores is counted as  (1 + 5 * # items of that type using this line) in
// 5 different buckets (armor, weapons, etc)
// The probability of any 1 line being picked for an item type at the start of the rolling process is then
// Count for that line / (sum of all counts for that type of item)
func (ss *ScorerStatistics) SetupProbabilityWeights() {
	// Klunky because indexing an array doesn't increase its size, i.e. just easier to use a map.
	// weightrand wants an array, so turn NumLines (map) into an array

	// find max key in NumLines
	maxidx := 0
	for ssidx := range ss.TypeStatistics {
		for rowIdx := range ss.TypeStatistics[ssidx].NumLines {
			if rowIdx > maxidx {
				maxidx = rowIdx
			}
		}
	}
	for ssidx := range ss.TypeStatistics {
		buckets := make([]int, maxidx+1)
		for rowIdx, count := range ss.TypeStatistics[ssidx].NumLines {
			buckets[rowIdx] += count * 5
		}
		for idx := range buckets {
			buckets[idx]++
		}
		ss.TypeStatistics[ssidx].Weights = weightrand.NewWeights(buckets)
	}
}
