package generator

import (
	"fmt"

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

// Generator Object for generating Props
type Generator struct {
	//s   *scorer.Scorer
	d2files    *d2fs.Files
	IFI        *d2fs.ItemFileInfo // Initialized inside genFile
	opts       config.RandomOptions
	Statistics *scorerstatistics.ScorerStatistics
	//RowToLine  []propscores.Line // Map from a PropScores.txt row index to a propscore.Line
	TypeTree      *d2items.TypeTree
	psi           *propscores.Maps
	numAffixRolls int
}

// NewGenerator Initialize a Generator from Scorer statistics
func NewGenerator(d2files *d2fs.Files, opts config.RandomOptions, tt *d2items.TypeTree, psi *propscores.Maps, stats *scorerstatistics.ScorerStatistics) *Generator {
	g := Generator{
		d2files:    d2files,
		opts:       opts,
		TypeTree:   tt,
		psi:        propscores.NewPropScoresIndex(d2files),
		Statistics: stats,
	}
	g.Statistics.SetupProbabilityWeights()
	if g.Statistics.TypeStatistics[0].Weights == nil {
		panic(1)
	}
	return &g
}

// Run Generate new props using the scores calculated by the Scorer
//  This routine and its children require the statistics gathered by propscorer to function
// Statistics are in Scorer:scoreLineWeights
func (g *Generator) Run() {
	fmt.Println("===============================")
	fmt.Println("| GenFiles                    |")
	fmt.Println("===============================")
	genFile(g, &uniqueItems.IFI)
	genFile(g, &setItems.IFI)
	genFile(g, &sets.IFI)
	genFile(g, &runes.IFI)
	fmt.Printf("%d affixes rolled\n", g.numAffixRolls)
}
