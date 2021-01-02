package generator

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/propscores"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/runes"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/setItems"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/sets"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/uniqueItems"
	"github.com/tlentz/d2modmaker/internal/d2mod/config"
	"github.com/tlentz/d2modmaker/internal/d2mod/d2items"
	"github.com/tlentz/d2modmaker/internal/d2mod/enhancedsets"
	"github.com/tlentz/d2modmaker/internal/d2mod/scorer/scorerstatistics"
	"github.com/tlentz/d2modmaker/internal/util"
)

// Generator Object for generating Props
type Generator struct {
	//s   *scorer.Scorer
	d2files    *d2fs.Files
	IFI        *d2fs.ItemFileInfo // Initialized inside genFile, points at the IFI for the file currently being generated
	opts       config.GeneratorOptions
	Statistics *scorerstatistics.ScorerStatistics
	//RowToLine  []propscores.Line // Map from a PropScores.txt row index to a propscore.Line
	TypeTree      *d2items.TypeTree
	psi           *propscores.Maps
	numAffixRolls int
	rng           *rand.Rand
	SetToUnique   map[string]d2items.Item
}

// NewGenerator Initialize a Generator from Scorer statistics
func NewGenerator(d2files *d2fs.Files, opts *config.GeneratorOptions, tt *d2items.TypeTree, psi *propscores.Maps, stats *scorerstatistics.ScorerStatistics) *Generator {
	opts.MinProps = util.MaxInt(1, opts.MinProps)
	opts.MaxProps = util.MinInt(20, opts.MaxProps)
	opts.NumClones = util.MaxInt(0, opts.NumClones)
	opts.PropScoreMultiplier = util.MinFloat(10, opts.PropScoreMultiplier)
	if !opts.UseSeed {
		opts.Seed = time.Now().UnixNano()
	}
	if !opts.UseSetsSeed {
		opts.SetsSeed = time.Now().UnixNano()
	}

	g := Generator{
		d2files:    d2files,
		opts:       *opts,
		TypeTree:   tt,
		psi:        propscores.NewPropScoresIndex(d2files),
		Statistics: stats,
	}
	g.Statistics.SetupProbabilityWeights()

	g.rng = rand.New(rand.NewSource(opts.Seed))
	g.SetToUnique = make(map[string]d2items.Item)

	return &g
}

// Run Generate new props using the scores calculated by the Scorer
//  This routine and its children require the statistics gathered by propscorer to function
// Statistics are in Scorer:scoreLineWeights
func (g *Generator) Run() {
	genFile(g, &uniqueItems.IFI) // Beware that the Unique Items must be generated before Set Items due to EnhancedSets
	genFile(g, &setItems.IFI)
	oldRng := g.rng
	g.rng = rand.New(rand.NewSource(g.opts.SetsSeed))
	genFile(g, &sets.IFI)
	if g.opts.EnhancedSets {
		enhancedsets.BlankFullSetBonuses(g.d2files)
		enhancedsets.SetAddFunc(g.d2files, 2)
	}
	g.rng = nil
	g.rng = oldRng
	genFile(g, &runes.IFI)
	fmt.Printf("%d affixes rolled\n", g.numAffixRolls)
}
