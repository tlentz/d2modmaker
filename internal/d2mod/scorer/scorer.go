package scorer

import (
	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/propscores"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/runes"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/setItems"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/sets"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/uniqueItems"
	"github.com/tlentz/d2modmaker/internal/d2mod/config"
	"github.com/tlentz/d2modmaker/internal/d2mod/d2items"
	"github.com/tlentz/d2modmaker/internal/d2mod/scorer/scorerstatistics"
)

const (
	synergyBonus = 1.1 // 10% increase.  Applies n time to every prop that has n siblings in same SynergyGroup
	twoHandNerf  = 0.7 // 30% decrease.  2handers (not 1 or 2handers) scores reduced due to hogging the shield slot.
)

// Scorer Accumulates setup and scoring information for scored files
type Scorer struct {
	d2files   *d2fs.Files          //
	opts      config.RandomOptions //
	Scorefile *d2fs.File           // PropScores.txt
	IFI       *d2fs.ItemFileInfo   //
	//TypeMap          map[string][]string     // Maps from a type to its parents.  From Weapons, Armor & ItemTypes
	Group        map[d2items.Prop]string  // Maps from Prop to Group
	SynergyGroup map[d2items.Prop]string  // Maps from a Prop to SynergyGroup
	items        map[string]*d2items.Item //
	//ScoreLineWeights []weights               // Keeps track of probabilities & counts for each type of item
	Statistics *scorerstatistics.ScorerStatistics
	TypeTree   *d2items.TypeTree
	PSI        *propscores.Maps
}

func newScorer(d2files *d2fs.Files, opts config.RandomOptions) *Scorer {
	s := Scorer{
		d2files: d2files,
		opts:    opts,
		//TypeMap:          map[string][]string{},
		Group:        map[d2items.Prop]string{},
		SynergyGroup: map[d2items.Prop]string{},
		items:        map[string]*d2items.Item{},
		Statistics:   scorerstatistics.NewScorerStatistics(),
		TypeTree:     d2items.NewTypeTree(d2files),
		PSI:          propscores.NewPropScoresIndex(d2files),
	}

	return &s
}

// Run Scores all Items files
func Run(d2files *d2fs.Files, opts config.RandomOptions) *Scorer {
	s := newScorer(d2files, opts)
	scoreFile(s, &uniqueItems.IFI)
	scoreFile(s, &setItems.IFI)
	scoreFile(s, &runes.IFI)
	scoreFile(s, &sets.IFI)
	debugshow(s)
	return s
}

func debugshow(s *Scorer) {
	//fmt.Printf("==========Debugging============\n")
	/*
		for propname, lines := range ScoreLines {
			fmt.Printf("%s\t%d\n", propname, len(lines))
		}
	*/
	/*
		for _, lines := range ScoreLines {
			for _, l := range lines {
				fmt.Printf("%s %d %d\n", l.Prop.Name, l.Prop.Val.Min, l.Prop.Val.Max)
			}
		}
	*/
	/*
		fmt.Printf("===============================\n")
			fmt.Printf("*****\nDumping typemap\n*****\n")
			for ctype, ptype := range s.TypeMap {
				fmt.Printf("%s\t%s\n", ctype, ptype)
			}
	*/
	/*
		fmt.Printf("===============================\n")
			fmt.Printf("Dumping Item Scores\n")
			for key, el := range s.Statistics.ItemScores {
				fmt.Printf("%s\t%d\n", key, el)
			}
	*/
	//fmt.Printf("===============================\n")
}
