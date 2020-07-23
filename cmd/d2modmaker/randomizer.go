package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"

	"github.com/tlentz/d2modmaker/internal/d2file"
	"github.com/tlentz/d2modmaker/internal/gemsTxt"
	misctxt "github.com/tlentz/d2modmaker/internal/miscTxt"
	"github.com/tlentz/d2modmaker/internal/runesTxt"
	"github.com/tlentz/d2modmaker/internal/setItemsTxt"
	"github.com/tlentz/d2modmaker/internal/setsTxt"
	"github.com/tlentz/d2modmaker/internal/uniqueItemsTxt"
	"github.com/tlentz/d2modmaker/internal/util"
)

// Prop represents an item affix
type Prop struct {
	Name string
	Par  string
	Min  string
	Max  string
	Lvl  int
}

// Props is a slice of Prop
type Props = []Prop

// Item represents an item
type Item struct {
	Name    string
	Lvl     int
	Affixes Props
}

// Items is a slice of Item
type Items = []Item

// RandomOptions are the options for the randomizer
type RandomOptions struct {
	Randomize         bool  `json:"Randomize"`
	Seed              int64 `json:"Seed"`
	IsBalanced        bool  `json:"IsBalanced"`        // Allows Props only from items up to 10 levels higher
	BalancedPropCount bool  `json:"BalancedPropCount"` // Picks prop count from a vanilla item up to 10 levels higher
	MinProps          int   `json:"MinProps"`          // minimum number of non blank props on an item
	MaxProps          int   `json:"MaxProps"`          // maximum number of non blank props on an item
	PerfectProps      bool  `json:"PerfectProps"`      // sets min/max to max
	UseOSkills        bool  `json:"UseOSkills"`        // +3 Fireball (Sorceress Only) -> +3 Fireball
}

func getRandomOptions(cfg *ModConfig) RandomOptions {
	defaultCfg := RandomOptions{
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

	cfg.RandomOptions.Seed = defaultCfg.Seed
	return defaultCfg
}

// Randomize randomizes all items based on the RandomOptions
func Randomize(cfg *ModConfig, d2files d2file.D2Files) {
	opts := getRandomOptions(cfg)
	rand.Seed(opts.Seed)

	props, items := getAllProps(opts, d2files)

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
	// writePropsToFile(props)
}

func writePropsToFile(props Props) {
	filePath := outDir + "props.json"
	fmt.Println("Writing " + filePath)
	file, _ := json.MarshalIndent(props, "", " ")
	_ = ioutil.WriteFile(filePath, file, 0644)
}

// Returns all props bucketized
func getAllProps(opts RandomOptions, d2files d2file.D2Files) (Props, Items) {
	props := Props{}
	items := Items{}

	uniqueItemProps, uniqueItems  := getAllUniqueProps(d2files)
	props = append(props, uniqueItemProps...)
	items = append(items, uniqueItems...)

	setProps, setBonuses = getAllSetProps(d2files)
	props = append(props, setProps...)
	//items = append(items, setBonuses...) //These aren't really items

	setItemProps, setItems = getAllSetItemsProps(d2files)
	props = append(props, setItemProps...)
	items = append(items, setItems...)

	runeWordProps, runewords = getAllRWProps(d2files)
	props = append(props, runeWordProps...)
	items = append(items, runewords...)

	//	gemProps, gems = getAllGemsProps(d2files);
	//	props = append(props, gemProps...)
	//	items = append(items, gems...)

	//TODO: move this to getProps so that we don't need to apply it to Items separately
	for i := range props {
		// Set all props Min to the Max value
		if opts.PerfectProps {
			props[i].Min = props[i].Max
		}
		// sets skill = oskill
		if opts.UseOSkills {
			if props[i].Name == "skill" {
				props[i].Name = "oskill"
			}
		}
	}

	return props, items
}

type propGetter struct {
	d2files     d2file.D2Files
	props       Props
	fileName    string
	propOffset  int
	levelOffset int
	nameOffset  int
}

func getProps(p propGetter) (Props, Items) {
	f := d2file.GetOrCreateFile(p.d2files, p.fileName)
	props := Props{}
	items := Items{}
	for _, row := range f.Rows {
		lvl := 0
		if p.levelOffset >= 0 {
			mbLvl, err := strconv.Atoi(row[p.levelOffset])
			if err == nil {
				lvl = mbLvl
			}
		} else { //Runes.txt doesn't have a level column
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
func getAllUniqueProps(d2files d2file.D2Files) (Props, Items) {
	p := propGetter{
		d2files:     d2files,
		fileName:    uniqueItemsTxt.FileName,
		propOffset:  uniqueItemsTxt.Prop1,
		levelOffset: uniqueItemsTxt.Lvl,
		nameOffset:  uniqueItemsTxt.Index,
	}
	return getProps(p)
}

// Randomize Unique Props
func randomizeUniqueProps(s scrambler) {
	s.fileName = uniqueItemsTxt.FileName
	s.propOffset = uniqueItemsTxt.Prop1
	s.itemMaxProps = uniqueItemsTxt.MaxNumProps
	s.minMaxProps = getMinMaxProps(s.opts, uniqueItemsTxt.MaxNumProps)
	s.lvl = uniqueItemsTxt.Lvl
	scramble(s)
}

// Get Set Props
func getAllSetProps(d2files d2file.D2Files) (Props, Items) {
	p := propGetter{
		d2files:     d2files,
		fileName:    setsTxt.FileName,
		propOffset:  setsTxt.PCode2a,
		levelOffset: setsTxt.Level,
		nameOffset:  setsTxt.Index,
	}
	return getProps(p)
}

// Randomize Set Props
func randomizeSetProps(s scrambler) {
	s.fileName = setsTxt.FileName
	s.propOffset = setsTxt.PCode2a
	s.itemMaxProps = setsTxt.MaxNumProps
	s.minMaxProps = getMinMaxProps(s.opts, setsTxt.MaxNumProps)
	s.lvl = setsTxt.Level
	scramble(s)
}

// Get Set Items Props
func getAllSetItemsProps(d2files d2file.D2Files) (Props, Items) {
	p := propGetter{
		d2files:     d2files,
		fileName:    setItemsTxt.FileName,
		propOffset:  setItemsTxt.Prop1,
		levelOffset: setItemsTxt.Lvl,
		nameOffset:  setItemsTxt.Index,
	}
	return getProps(p)
}

// Randomize Set Items Props
func randomizeSetItemsProps(s scrambler) {
	s.fileName = setItemsTxt.FileName
	s.propOffset = setItemsTxt.Prop1
	s.itemMaxProps = setItemsTxt.MaxNumProps
	s.minMaxProps = getMinMaxProps(s.opts, setItemsTxt.MaxNumProps)
	s.lvl = setItemsTxt.Lvl
	scramble(s)
}

// Get RW Props
func getAllRWProps(d2files d2file.D2Files) (Props, Items) {
	p := propGetter{
		d2files:     d2files,
		fileName:    runesTxt.FileName,
		propOffset:  runesTxt.T1Code1,
		levelOffset: -1,
		nameOffset:  runesTxt.RuneName,
	}
	return getProps(p)
}

// Randomize RW Props
func randomizeRWProps(s scrambler) {
	s.fileName = runesTxt.FileName
	s.propOffset = runesTxt.T1Code1
	s.itemMaxProps = runesTxt.MaxNumProps
	s.minMaxProps = getMinMaxProps(s.opts, runesTxt.MaxNumProps)

	f := d2file.GetOrCreateFile(s.d2files, runesTxt.FileName)
	miscLevels := getMiscItemLevels(s.d2files)
	for idx, row := range f.Rows {
		level := getRunewordLevel(row, miscLevels)
		scrambleRow(s, f, idx, level)
	}
}

func getRunewordLevel(row []string, miscLevels map[string]int) int {
	runeLevels := []int{}
	for j := 0; j < 6; j++ {
		runeLevels = append(runeLevels, miscLevels[row[runesTxt.Rune1+j]])
	}
	return getMaxInt(runeLevels)
}

// Get Gem Props
func getAllGemsProps(d2files d2file.D2Files) Props {
	f := d2file.GetOrCreateFile(d2files, gemsTxt.FileName)
	propOffset := gemsTxt.WeaponMod1Code
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

func getBalancedRandomProp(opts RandomOptions, lvl int, props Props) Prop {
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

func getMiscItemLevels(d2files d2file.D2Files) map[string]int {
	f := d2file.GetOrCreateFile(d2files, misctxt.FileName)
	itemMap := make(map[string]int)
	for _, row := range f.Rows {
		n, err := strconv.Atoi(row[misctxt.Level])
		if err == nil {
			itemMap[row[misctxt.Code]] = n
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

func getAdjustNumProps(opts RandomOptions) bool {
	return opts.MinProps >= 0 || opts.MaxProps >= 0
}

func getMinMaxProps(opts RandomOptions, maxItemProps int) minMaxProps {
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
	opts         RandomOptions
	d2files      d2file.D2Files
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
	f := d2file.GetOrCreateFile(s.d2files, s.fileName)
	for idx, row := range f.Rows {
		level := 0
		rowLvl, err := strconv.Atoi(row[s.lvl])
		if err == nil {
			level = rowLvl
		}
		scrambleRow(s, f, idx, level)
	}
}

func scrambleRow(s scrambler, f *d2file.D2File, idx int, level int) {
	//Choose a random number of props between min and max
	numProps := randInt(s.minMaxProps.minNumProps, s.minMaxProps.maxNumProps+1)

	//If using balanced prop counts, override the random count
	if opts.BalancedPropCount {
		item := s.items[randInt(0, len(s.items))]
		if s.opts.IsBalanced {
			for item.Lvl-s.lvl > 10 {
				item = s.items[randInt(0, len(s.items))]
			}
		}
		numProps = util.MinInt(len(item.Affixes), s.minMaxProps.maxNumProps)
	}

	// fill in the props
	for currentNumProps := 0; currentNumProps < s.itemMaxProps; currentNumProps++ {
		prop := Prop{Name: "", Par: "", Min: "", Max: ""}
		if currentNumProps < numProps {
			prop = getBalancedRandomProp(s.opts, level, s.props)
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
