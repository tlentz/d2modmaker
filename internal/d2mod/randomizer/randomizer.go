package randomizer

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/tlentz/d2modmaker/internal/d2mod/config"
	"github.com/tlentz/d2modmaker/internal/d2mod/elementalskills"
	"github.com/tlentz/d2modmaker/internal/d2mod/prop"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/gems"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/misc"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/runes"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/setItems"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/sets"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/uniqueItems"

	"github.com/tlentz/d2modmaker/internal/util"
)

// Prop == d2items.Prop
type Prop = prop.Prop

// Props == d2items.Prop
type Props = prop.Props

// Item represents an item
type Item struct {
	Name    string
	Lvl     int
	Affixes Props
}

// Items is a slice of Item
type Items = []Item

// Randomize randomizes all items based on the RandomOptions
func Run(cfg *config.Data, d2files d2fs.Files) {
	opts := getRandomOptions(cfg)
	rand.Seed(opts.Seed)

	props, items := getAllProps(opts, d2files)
	if cfg.RandomOptions.ElementalSkills {
		props = append(props, elementalskills.Props()...)
	}

	s := scrambler{
		opts:    opts,
		d2files: d2files,
		props:   props,
		items:   items,
	}

	randomizeUniqueProps(s)
	randomizeSetProps(s)
	randomizeSetItemsProps(s)
	randomizeRWProps(s)
	randomizeGemAndRuneProps(s)
}

func getRandomOptions(cfg *config.Data) config.RandomOptions {
	defaultCfg := config.RandomOptions{
		Seed:     time.Now().UnixNano(),
		MinProps: -1,
		MaxProps: -1,
	}
	if cfg.RandomOptions.Seed >= 0 {
		defaultCfg.Seed = cfg.RandomOptions.Seed
	}
	defaultCfg.IsBalanced = cfg.RandomOptions.IsBalanced
	if cfg.RandomOptions.MaxProps >= 0 {
		defaultCfg.MaxProps = cfg.RandomOptions.MaxProps
	}
	if cfg.RandomOptions.MinProps >= 0 {
		defaultCfg.MinProps = cfg.RandomOptions.MinProps
	}
	defaultCfg.PerfectProps = cfg.RandomOptions.PerfectProps
	defaultCfg.UseOSkills = cfg.RandomOptions.UseOSkills
	// TODO: add configuration for Gems and Runes
	//defaultCfg.RandomizeRunes = cfg.RandomOptions.RandomizeRunes

	cfg.RandomOptions.Seed = defaultCfg.Seed

	defaultCfg.BalancedPropCount = cfg.RandomOptions.BalancedPropCount

	return defaultCfg
}

// Returns all props bucketized
func getAllProps(opts config.RandomOptions, d2files d2fs.Files) (Props, Items) {
	props := Props{}
	items := Items{}

	p := propGetter{
		d2files: d2files,
		opts:    opts,
	}

	uniqueItemProps, uniqueItems := getAllUniqueProps(p)
	props = append(props, uniqueItemProps...)
	items = append(items, uniqueItems...)

	setProps, _ := getAllSetProps(p)
	props = append(props, setProps...)
	//items = append(items, setBonuses...) //These aren't really items

	setItemProps, setItems := getAllSetItemsProps(p)
	props = append(props, setItemProps...)
	items = append(items, setItems...)

	runeWordProps, runewords := getAllRWProps(p)
	props = append(props, runeWordProps...)
	items = append(items, runewords...)

	// do not add gem & rune properties to global shuffle
	// gemProps, gems := getAllGemsAndRunes(p)
	// props = append(props, gemProps...)
	// items = append(items, gems...)

	return props, items
}

type propGetter struct {
	d2files     d2fs.Files
	opts        config.RandomOptions
	props       Props
	fileName    string
	propOffset  int
	levelOffset int
	nameOffset  int
}

func getProps(p propGetter) (Props, Items) {
	f := p.d2files.Get(p.fileName)
	props := Props{}
	items := Items{}
	for _, row := range f.Rows {
		lvl := 0
		if p.levelOffset >= 0 {
			mbLvl, err := strconv.Atoi(row[p.levelOffset])
			if err == nil {
				lvl = mbLvl
			}
		} else if p.fileName == runes.FileName { //Runes.txt doesn't have a level column
			//TODO: It would be nice to not call getMiscItemLevels every time
			lvl = getRunewordLevel(row, getMiscItemLevels(p.d2files))
		}

		item := Item{}
		item.Name = row[p.nameOffset]
		item.Lvl = lvl

		for i := p.propOffset; i < len(row)-3; i += 4 {
			prop := Prop{
				Name: row[i],
				Par:  row[i+1],
				Min:  row[i+2],
				Max:  row[i+3],
				Lvl:  lvl,
			}
			if prop.Name != "" {
				if p.opts.PerfectProps {
					prop.Min = prop.Max
				}
				if p.opts.UseOSkills {
					if prop.Name == "skill" {
						prop.Name = "oskill"
					}
				}

				props = append(props, prop)
				item.Affixes = append(item.Affixes, prop)
			}
		}

		if item.Name != "" && len(item.Affixes) > 0 {
			items = append(items, item)
		}
	}
	return props, items
}

// Get Unique Props
func getAllUniqueProps(p propGetter) (Props, Items) {
	p.fileName = uniqueItems.FileName
	p.propOffset = uniqueItems.Prop1
	p.levelOffset = uniqueItems.Lvl
	p.nameOffset = uniqueItems.Index
	return getProps(p)
}

// Randomize Unique Props
func randomizeUniqueProps(s scrambler) {
	s.fileName = uniqueItems.FileName
	s.propOffset = uniqueItems.Prop1
	s.itemMaxProps = uniqueItems.MaxNumProps
	s.minMaxProps = getMinMaxProps(s.opts, uniqueItems.MaxNumProps)
	s.lvl = uniqueItems.Lvl
	scramble(s)
}

// Get Set Props
func getAllSetProps(p propGetter) (Props, Items) {
	p.fileName = sets.FileName
	p.propOffset = sets.PCode2a
	p.levelOffset = sets.Level
	p.nameOffset = sets.Index
	return getProps(p)
}

// Randomize Set Props
func randomizeSetProps(s scrambler) {
	s.fileName = sets.FileName
	s.propOffset = sets.PCode2a
	s.itemMaxProps = sets.MaxNumProps
	s.minMaxProps = getMinMaxProps(s.opts, sets.MaxNumProps)
	s.lvl = sets.Level
	scramble(s)
}

// Get Set Items Props
func getAllSetItemsProps(p propGetter) (Props, Items) {
	p.fileName = setItems.FileName
	p.propOffset = setItems.Prop1
	p.levelOffset = setItems.Lvl
	p.nameOffset = setItems.Index
	return getProps(p)
}

// Randomize Set Items Props
func randomizeSetItemsProps(s scrambler) {
	s.fileName = setItems.FileName
	s.propOffset = setItems.Prop1
	s.itemMaxProps = setItems.MaxNumProps
	s.minMaxProps = getMinMaxProps(s.opts, setItems.MaxNumProps)
	s.lvl = setItems.Lvl
	scramble(s)
}

// Get RW Props
func getAllRWProps(p propGetter) (Props, Items) {
	p.fileName = runes.FileName
	p.propOffset = runes.T1Code1
	p.levelOffset = -1
	p.nameOffset = runes.RuneName
	return getProps(p)
}

// Randomize RW Props
func randomizeRWProps(s scrambler) {
	s.fileName = runes.FileName
	s.propOffset = runes.T1Code1
	s.itemMaxProps = runes.MaxNumProps
	s.minMaxProps = getMinMaxProps(s.opts, runes.MaxNumProps)

	f := s.d2files.Get(runes.FileName)
	miscLevels := getMiscItemLevels(s.d2files)
	for idx, row := range f.Rows {
		level := getRunewordLevel(row, miscLevels)
		scrambleRow(s, f, idx, level)
	}
}

func getRunewordLevel(row []string, miscLevels map[string]int) int {
	runeLevels := []int{}
	for j := 0; j < 6; j++ {
		runeLevels = append(runeLevels, miscLevels[row[runes.Rune1+j]])
	}
	return getMaxInt(runeLevels)
}

func getAllGemsAndRunes(p propGetter) (Props, Items) {
	p.fileName = gems.FileName
	p.propOffset = gems.WeaponMod1Code
	p.levelOffset = -1
	p.nameOffset = gems.Name
	return getProps(p)
}

func randomizeGemAndRuneProps(s scrambler) {
	s.fileName = gems.FileName
	s.propOffset = gems.NumMods

	// do gem shuffle by category/color
	//shuffleGemsByColor(s)
	shuffleGemsByTier(s)

	// shuffle rune properties
	shuffleRuneProperties(s)
}

func shuffleGemsByColor(s scrambler) {
	f := s.d2files.Get(s.fileName)
	origRows := make([][]string, len(f.Rows))
	for i := range origRows {
		origRows[i] = make([]string, len(f.Rows[i]))
		copy(origRows[i], f.Rows[i])
	}

	// 5 rows = 1 color
	// 7 colors total
	order := makeRange(0, 6)
	rand.Shuffle(len(order), func(i, j int) { order[i], order[j] = order[j], order[i] })

	for i, o := range order {
		for j := 0; j < 5; j++ {
			swapGemPropertiesFromCache(s, f, origRows, (i*5)+j, (o*5)+j)
		}
	}
}

func shuffleGemsByTier(s scrambler) {
	f := s.d2files.Get(s.fileName)
	origRows := make([][]string, len(f.Rows))
	for i := range origRows {
		origRows[i] = make([]string, len(f.Rows[i]))
		copy(origRows[i], f.Rows[i])
	}

	// 5 rows = 1 color
	// 7 colors total
	order := [][]int{makeRange(0, 6), makeRange(0, 6), makeRange(0, 6), makeRange(0, 6), makeRange(0, 6)}

	// each tier has their own order
	for tier := 0; tier < len(order); tier++ {
		rand.Shuffle(len(order[tier]), func(i, j int) { order[tier][i], order[tier][j] = order[tier][j], order[tier][i] })
	}

	// progresses sequentially through 7 chips, then 7 flawed, etc.
	for i, o := range order {
		for j := 0; j < 7; j++ {
			swapGemPropertiesFromCache(s, f, origRows, (j*5)+i, (o[j]*5)+i)
		}
	}
}

func shuffleRuneProperties(s scrambler) {
	f := s.d2files.Get(s.fileName)
	origRows := make([][]string, len(f.Rows))
	for i := range origRows {
		origRows[i] = make([]string, len(f.Rows[i]))
		copy(origRows[i], f.Rows[i])
	}

	const firstRuneRow = 36
	const lastRuneRow = 68
	order := makeRange(firstRuneRow, lastRuneRow)
	rand.Shuffle(len(order), func(i, j int) { order[i], order[j] = order[j], order[i] })

	for i := firstRuneRow; i <= lastRuneRow; i++ {
		swapGemPropertiesFromCache(s, f, origRows, i, order[i-firstRuneRow])
	}
}

func swapGemPropertiesFromCache(s scrambler, f *d2fs.File, origRows [][]string, destRow int, sourceRow int) {
	//fmt.Printf("Swapping\n%+v\nWith\n%+v\n", f.Rows[destRow], origRows[sourceRow])
	for i := range f.Rows[destRow] {
		if i < s.propOffset {
			continue
		}
		f.Rows[destRow][i] = origRows[sourceRow][i]
	}
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
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

func getBalancedRandomProp(opts config.RandomOptions, lvl int, props Props) Prop {
	prop := Prop{}
	numProps := len(props)

	for prop.Name == "" {
		prop = props[randInt(0, numProps)]
		//Check if this prop is balanced if using that feature
		if opts.IsBalanced && prop.Lvl-lvl > 10 {
			// Blank the prop name and pick again
			prop.Name = ""
		}
	}

	return prop
}

func getMiscItemLevels(d2files d2fs.Files) map[string]int {
	f := d2files.Get(misc.FileName)
	itemMap := make(map[string]int)
	for _, row := range f.Rows {
		n, err := strconv.Atoi(row[misc.Level])
		if err == nil {
			itemMap[row[misc.Code]] = n
		}
	}
	return itemMap
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
	d2files      d2fs.Files
	props        Props
	items        Items
	fileName     string
	propOffset   int
	minMaxProps  minMaxProps
	itemMaxProps int
	lvl          int
}

type minMaxProps struct {
	minNumProps int
	maxNumProps int
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
	//Choose a random number of props between min and max
	numProps := randInt(s.minMaxProps.minNumProps, s.minMaxProps.maxNumProps+1)

	//If using balanced prop counts, override the random count
	if s.opts.BalancedPropCount {
		item := s.items[randInt(0, len(s.items))]
		if s.opts.IsBalanced {
			for item.Lvl-s.lvl > 10 {
				item = s.items[randInt(0, len(s.items))]
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
			prop = getBalancedRandomProp(s.opts, level, s.props)

			propIdString := prop.GetID()
			for propList[propIdString] {
				prop = getBalancedRandomProp(s.opts, s.lvl, s.props)
				propIdString = prop.GetID()
			}

			// Add used prop to the prop list if duplicate properties are not allowed
			// Always add aura to the prop list because multiple auras on an item are broken
			if !s.opts.AllowDupeProps || propIdString == "aura" {
				propList[propIdString] = true
			}
		}
		i := s.propOffset + currentNumProps*4
		f.Rows[idx][i] = prop.Name
		f.Rows[idx][i+1] = prop.Par
		f.Rows[idx][i+2] = prop.Min
		f.Rows[idx][i+3] = prop.Max

	}
}

func randInt(min int, max int) int {
	if min == max {
		return min
	}
	return min + rand.Intn(max-min)
}
