package randomizer

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/tlentz/d2modmaker/internal/d2mod/config"
	"github.com/tlentz/d2modmaker/internal/d2mod/d2items"
	"github.com/tlentz/d2modmaker/internal/d2mod/elementalskills"
	"github.com/tlentz/d2modmaker/internal/d2mod/enhancedsets"
	"github.com/tlentz/d2modmaker/internal/d2mod/prop"
	"github.com/tlentz/d2modmaker/internal/d2mod/propscores"
	"github.com/tlentz/d2modmaker/internal/d2mod/runewordlevels"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/gems"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/runes"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/setItems"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/sets"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/uniqueItems"

	"fmt"

	"github.com/tlentz/d2modmaker/internal/util"
)

// Prop == d2items.Prop
type Prop = prop.Prop

// Props == d2items.Prop
type Props = prop.Props

// Item == d2items.Item
type Item = d2items.Item

// Items == d2items.Items
type Items = d2items.Items

func newScrambler(cfg *config.RandomOptions, d2files *d2fs.Files) (s *scrambler) {
	psi := propscores.NewPropScoresIndex(d2files)
	tt := d2items.NewTypeTree(d2files)
	props, items := getAllProps(d2files, psi, *tt)
	if cfg.ElementalSkills {
		props = append(props, elementalskills.Props()...)
	}

	if !cfg.UseSeed {
		cfg.Seed = time.Now().UnixNano()
		//fmt.Printf("New Seed = %d", cfg.Seed)
	}
	cfg.MinProps = util.MaxInt(1, cfg.MinProps)
	cfg.MaxProps = util.MinInt(20, cfg.MaxProps)

	snew := scrambler{
		opts:    *cfg,
		d2files: d2files,
		props:   props,
		items:   items,
		rng:     rand.New(rand.NewSource(cfg.Seed)),
	}
	return &snew
}

// Run Randomize randomizes all items based on the RandomOptions
func Run(cfg *config.RandomOptions, d2files *d2fs.Files) {
	s := newScrambler(cfg, d2files)

	s.rng.Seed(s.opts.Seed)
	randomizeUniqueProps(*s)
	randomizeSetItemsProps(*s)
	randomizeRWProps(*s)
	randomizeSetProps(*s)
}

// Returns all props bucketized
func getAllProps(d2files *d2fs.Files, psi *propscores.Maps, tt d2items.TypeTree) (Props, Items) {
	props := Props{}
	items := Items{}
	var p *d2items.PropGetter

	p = d2items.NewPropGetter(d2files, &uniqueItems.IFI, psi, tt)
	uniqueItemProps, uniqueItems := p.GetProps()
	props = append(props, uniqueItemProps...)
	items = append(items, uniqueItems...)

	p = d2items.NewPropGetter(d2files, &sets.IFI, psi, tt)
	setProps, _ := p.GetProps()
	d2items.AppendProps(props, setProps)
	//items = append(items, setBonuses...) //These aren't really items

	p = d2items.NewPropGetter(d2files, &setItems.IFI, psi, tt)
	setItemProps, setItems := p.GetProps()
	props = append(props, setItemProps...)
	items = append(items, setItems...)

	p = d2items.NewPropGetter(d2files, &runes.IFI, psi, tt)
	runeWordProps, runewords := p.GetProps()
	props = append(props, runeWordProps...)
	items = append(items, runewords...)

	//	gemProps, gems = getAllGemsProps(d2files);
	//	props = append(props, gemProps...)
	//	items = append(items, gems...)

	return props, items
}

// Randomize Unique Props
func randomizeUniqueProps(s scrambler) {
	s.fileName = uniqueItems.FileName
	s.propOffset = uniqueItems.Prop1
	s.itemMaxProps = uniqueItems.MaxNumProps
	s.minMaxProps = getMinMaxProps(s.opts, uniqueItems.MaxNumProps)
	s.lvl = uniqueItems.Lvl

	f := s.d2files.Get(s.fileName)
	cloneTable(f, s.opts.NumClones)

	scramble(s)
}

// Randomize Set Props
func randomizeSetProps(s scrambler) {
	if s.opts.EnhancedSets {
		s.fileName = sets.FileName
		s.propOffset = sets.PCode2a
		s.itemMaxProps = sets.MaxNumProps
		s.minMaxProps = getMinMaxProps(s.opts, sets.MaxNumProps)
		s.lvl = sets.Level
		scramble(s)
	} else {
		enhancedsets.BlankFullSetBonuses(s.d2files)
	}

}

// Randomize Set Items Props
func randomizeSetItemsProps(s scrambler) {
	s.fileName = setItems.FileName
	s.propOffset = setItems.Prop1
	s.itemMaxProps = 9 // OBC: TODO: Fix this hardcoding by changing setItems.go
	s.minMaxProps = getMinMaxProps(s.opts, 9)
	s.lvl = setItems.Lvl

	// OBC: The problem with cloneTable for sets is that if you
	// ever reduce NumClones, existing items will disappear.
	//f := s.d2files.Get(s.fileName)
	//cloneTable(f, s.opts.NumClones)

	scramble(s)

	// Above scrambled props 1-9, now populate the partial set bonuses.
	s.propOffset = setItems.AProp1a
	s.itemMaxProps = 10
	s.minMaxProps = getMinMaxProps(s.opts, 10) // (AProp1-AProp5) * 2 (a & b)
	scramble(s)                                // OBC:  It would be nice if this call to scramble would always generate 10 props even if balancedpropcount is on.

	if s.opts.EnhancedSets {
		// Add Func (f.Rows[][16]) controls how the AProp* columns show up
		// If Add Func == 2, then for each additional piece worn, a pair of props (a & b)
		// will be added as Green partial set bonuses
		// If Add Func == "", then all of the props in Prop* and AProp* show up at once.
		enhancedsets.SetAddFunc(s.d2files, 2)
	}
}

// Randomize RW Props
func randomizeRWProps(s scrambler) {
	s.fileName = runes.FileName
	s.propOffset = runes.T1Code1
	s.itemMaxProps = runes.MaxNumProps
	s.minMaxProps = getMinMaxProps(s.opts, runes.MaxNumProps)

	f := s.d2files.Get(runes.FileName)

	miscLevels := runewordlevels.GetMiscItemLevels(s.d2files)

	for idx, row := range f.Rows {
		level := runewordlevels.GetRunewordLevel(row, miscLevels)
		scrambleRow(s, f, idx, level)
	}

}

// Get Gem Props
func getAllGemsProps(d2files d2fs.Files) Props {
	f := d2files.Get(gems.FileName)
	propOffset := gems.WeaponMod1Code
	props := Props{}
	for _, row := range f.Rows {
		for i := propOffset; i < len(row)-3; i += 4 {
			props = append(props, Prop{
				Name: row[i],
				Par:  row[i+1],
				Min:  row[i+2],
				Max:  row[i+3],
				Lvl:  0,
			})
		}
	}
	return props
}

func getBalancedRandomProp(s scrambler, lvl int, props Props) Prop {
	prop := Prop{}
	numProps := len(props)

	for prop.Name == "" {
		prop = props[randInt(s.rng, 0, numProps)]
		//Check if this prop is balanced if using that feature
		//fmt.Printf("%s: %d|", prop.Name, prop.Lvl)
		if s.opts.IsBalanced && prop.Lvl-lvl > 10 {
			// Blank the prop name and pick again
			prop.Name = ""
		}
	}

	return prop
}

func getMaxInt(list []int) int {
	max := 0
	for i, e := range list {
		if i == 0 || e > max {
			max = e
		}
	}
	return max
}

func getAdjustNumProps(opts config.RandomOptions) bool {
	return opts.MinProps >= 0 || opts.MaxProps >= 0
}

func getMinMaxProps(opts config.RandomOptions, maxItemProps int) minMaxProps {
	min := util.MinInt(maxItemProps, util.MaxInt(0, opts.MinProps))
	max := maxItemProps
	if opts.MaxProps > 0 {
		max = util.MinInt(opts.MaxProps, maxItemProps)
	}
	return minMaxProps{
		minNumProps: min,
		maxNumProps: util.MaxInt(min, max),
	}
}

type scrambler struct {
	opts         config.RandomOptions
	d2files      *d2fs.Files
	props        Props
	items        Items
	fileName     string
	propOffset   int
	minMaxProps  minMaxProps
	itemMaxProps int
	lvl          int
	rng          *rand.Rand
}

type minMaxProps struct {
	minNumProps int
	maxNumProps int
}

// cloneTable: Creates a copy of the table N time, with several exceptions
// - Quest items should not be duped
// TODO: In the future if we have patchstring.tbl support we could use new names instead of copying the old names.
func cloneTable(f *d2fs.File, numClones int) {
	if numClones < 1 {
		numClones = 0
		return
	}
	if (len(f.Rows) * (numClones + 1)) > 4090 { // Limit for any file is 4095 rows
		numClones = int(4090/len(f.Rows)) - 1
		fmt.Printf("NumClones too large, clamped to %d\n", numClones)
	}
	// Deep copy the old row to the new clone row
	originalLength := len(f.Rows)
	newrows := make([][]string, originalLength*numClones, originalLength*numClones)
	f.Rows = append(f.Rows, newrows...)
	for i := 0; i < numClones; i++ {
		for j := 0; j < originalLength; j++ {
			var rowidx = (originalLength) + i*originalLength + j
			f.Rows[rowidx] = make([]string, len(f.Rows[j]))
			if f.Rows[j][6] != "0" {
				for colidx, col := range f.Rows[j] {
					f.Rows[rowidx][colidx] = col
				}
			} else {
				// Dupe quest items cause the quests to fail.
				// Quest items are level 0 items.  Don't clone them, make them be blank lines.
				//fmt.Printf("Quest Item:%s\n", f.Rows[j][0])
				f.Rows[rowidx][0] = ""                 // f.Rows[j][0]
				f.Rows[rowidx][len(f.Rows[j])-1] = "0" // keep 0 in *eol
			}
		}
	}

}

func scramble(s scrambler) {
	f := s.d2files.Get(s.fileName)
	for idx, row := range f.Rows {
		level := 0
		rowLvl, err := strconv.Atoi(row[s.lvl])
		if err == nil {
			level = rowLvl
		}
		scrambleRow(s, f, idx, level)
	}
}

func scrambleRow(s scrambler, f *d2fs.File, idx int, level int) {

	if f.Rows[idx][1] == "" {
		// Don't run scrambler on row dividers
		return
	}
	if level == 0 {
		// Don't run scrambler on Quest items.  It may raise level requirements
		return
	}

	//Choose a random number of props between min and max
	numProps := randInt(s.rng, s.minMaxProps.minNumProps, s.minMaxProps.maxNumProps+1)

	//If using balanced prop counts, override the random count
	if s.opts.BalancedPropCount {
		item := s.items[randInt(s.rng, 0, len(s.items))]
		if s.opts.IsBalanced {
			for item.Lvl-s.lvl > 10 {
				item = s.items[randInt(s.rng, 0, len(s.items))]
			}
		}
		numProps = util.MinInt(len(item.Affixes), s.minMaxProps.maxNumProps)
	}

	//map to track duplicate properties
	propList := make(map[string]bool)

	// fill in the props
	for currentNumProps := 0; currentNumProps < s.itemMaxProps; currentNumProps++ {
		prop := Prop{Name: "", Par: "", Min: "", Max: ""}
		if currentNumProps < numProps {
			prop = getBalancedRandomProp(s, level, s.props)

			propIDString := prop.GetID()
			for propList[propIDString] {
				prop = getBalancedRandomProp(s, s.lvl, s.props)
				propIDString = prop.GetID() // FIXME: TODO: This doesn't prevent duplicate props i.e. ac/lvl with diff pars
			}

			// Add used prop to the prop list if duplicate properties are not allowed
			// Always add aura to the prop list because multiple auras on an item are broken
			if !s.opts.AllowDupeProps || propIDString == "aura" {
				propList[propIDString] = true
			}
		}
		i := s.propOffset + currentNumProps*4
		f.Rows[idx][i] = prop.Name
		f.Rows[idx][i+1] = prop.Par
		f.Rows[idx][i+2] = prop.Min
		f.Rows[idx][i+3] = prop.Max

	}
}

// Beware: min <= randint < max   i.e. never returns max
func randInt(rng *rand.Rand, min int, max int) int {
	if min == max {
		return min
	}
	return min + rng.Intn(max-min)
}
