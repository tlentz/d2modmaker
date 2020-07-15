package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"sort"
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

// Prop represents an item
type Prop struct {
	Name string
	Par  string
	Min  string
	Max  string
	Lvl  int
}

// Props is a slice of Prop
type Props = []Prop

// BucketedProps is a map holding Props for each bucket
type BucketedProps = map[int]Props

// BucketedPropsMap is a map with the prop name as the key holding a BucketedProps
type BucketedPropsMap = map[string]BucketedProps

// RandomOptions are the options for the randomizer
type RandomOptions struct {
	Randomize    bool  `json:"Randomize"`
	Seed         int64 `json:"Seed"`
	IsBalanced   bool  `json:"IsBalanced"`   // bucketizes props [0-30] [31-60] [61+]
	MinProps     int   `json:"MinProps"`     // minimum number of non blank props on an item
	MaxProps     int   `json:"MaxProps"`     // maximum number of non blank props on an item
	PerfectProps bool  `json:"PerfectProps"` // sets min/max to max
	UseOSkills   bool  `json:"UseOSkills"`   // +3 Fireball (Sorceress Only) -> +3 Fireball
}

const (
	bucketAll = 0
	bucket0   = 1
	bucket30  = 2
	bucket60  = 3
)

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

	props, propKeys := getAllProps(opts, d2files)
	miscBuckets := getBucketsForMisc(d2files)

	randomizeUniqueProps(opts, d2files, props, propKeys)
	randomizeSetProps(opts, d2files, props, propKeys)
	randomizeSetItemsProps(opts, d2files, props, propKeys)
	randomizeRWProps(opts, miscBuckets, d2files, props, propKeys)
	// writePropBuckets(props)
}

func writePropBuckets(props BucketedPropsMap) {
	filePath := outDir + "buckets.json"
	fmt.Println("Writing " + filePath)
	file, _ := json.MarshalIndent(props, "", " ")
	_ = ioutil.WriteFile(filePath, file, 0644)
}

// Adds prop to correct map/bucket
func addOrCreateProp(props BucketedPropsMap, prop Prop) BucketedPropsMap {
	buckets := getBalanceBuckets(prop.Lvl)
	if _, ok := props[prop.Name]; !ok {
		props[prop.Name] = make(map[int][]Prop)
		props[prop.Name][bucketAll] = make([]Prop, 0)
		props[prop.Name][bucket0] = make([]Prop, 0)
		props[prop.Name][bucket30] = make([]Prop, 0)
		props[prop.Name][bucket60] = make([]Prop, 0)
	}
	for bucket := range buckets {
		props[prop.Name][bucket] = append(props[prop.Name][bucket], prop)
	}
	props[prop.Name][bucketAll] = append(props[prop.Name][bucketAll], prop)
	return props
}

// Returns all props bucketized
func getAllProps(opts RandomOptions, d2files d2file.D2Files) (BucketedPropsMap, []string) {

	propMap := BucketedPropsMap{}
	props := [][]Prop{}
	props = append(props, getAllUniqueProps(d2files, []Prop{}))
	props = append(props, getAllSetProps(d2files, []Prop{}))
	props = append(props, getAllSetItemsProps(d2files, []Prop{}))
	props = append(props, getAllRWProps(d2files, []Prop{}))
	props = append(props, getAllGemsProps(d2files, []Prop{}))

	for i := range props {
		for j := range props[i] {
			propMap = addOrCreateProp(propMap, props[i][j])
		}
	}

	for k := range propMap {
		for b := range propMap[k] {
			for i, p := range propMap[k][b] {
				// Set all props Min to the Max value
				if opts.PerfectProps {
					propMap[k][b][i].Min = p.Max
				}
				// sets skill = oskill
				if opts.UseOSkills {
					if p.Name == "skill" {
						propMap[k][b][i].Name = "oskill"
					}
				}
			}
		}
	}

	var keys []string
	for k := range propMap {
		totalProps := 0
		for b := range propMap[k] {
			totalProps += len(propMap[k][b])
		}
		for i := 0; i < totalProps; i++ {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	return propMap, keys
}

type PropGetter struct {
	d2files    d2file.D2Files
	props      Props
	fileName   string
	propOffset int
	lvl        int
}

func getProps(p PropGetter) Props {
	f := d2file.GetOrCreateFile(dataDir, p.d2files, p.fileName)
	for _, row := range f.Rows {
		lvl := 0
		if p.lvl >= 0 {
			mbLvl, err := strconv.Atoi(row[p.lvl])
			if err == nil {
				lvl = mbLvl
			}
		}

		for i := p.propOffset; i < len(row)-3; i += 4 {
			prop := Prop{
				Name: row[i],
				Par:  row[i+1],
				Min:  row[i+2],
				Max:  row[i+3],
				Lvl:  lvl,
			}
			p.props = append(p.props, prop)
		}
	}
	return p.props
}

// Get Unique Props
func getAllUniqueProps(d2files d2file.D2Files, props Props) Props {
	p := PropGetter{
		d2files:    d2files,
		props:      props,
		fileName:   uniqueItemsTxt.FileName,
		propOffset: uniqueItemsTxt.Prop1,
		lvl:        uniqueItemsTxt.Lvl,
	}
	return getProps(p)
}

// Randomize Unique Props
func randomizeUniqueProps(opts RandomOptions, d2files d2file.D2Files, props BucketedPropsMap, propKeys []string) {
	s := Scrambler{
		opts:         opts,
		d2files:      d2files,
		props:        props,
		propKeys:     propKeys,
		fileName:     uniqueItemsTxt.FileName,
		propOffset:   uniqueItemsTxt.Prop1,
		itemMaxProps: uniqueItemsTxt.MaxNumProps,
		lvl:          uniqueItemsTxt.Lvl,
	}
	s.minMaxProps = getMinMaxProps(opts, uniqueItemsTxt.MaxNumProps)
	scramble(s)
}

// Get Set Props
func getAllSetProps(d2files d2file.D2Files, props Props) Props {
	p := PropGetter{
		d2files:    d2files,
		props:      props,
		fileName:   setsTxt.FileName,
		propOffset: setsTxt.PCode2a,
		lvl:        setsTxt.Level,
	}
	return getProps(p)
}

// Randomize Set Props
func randomizeSetProps(opts RandomOptions, d2files d2file.D2Files, props BucketedPropsMap, propKeys []string) {
	s := Scrambler{
		opts:         opts,
		d2files:      d2files,
		props:        props,
		propKeys:     propKeys,
		fileName:     setsTxt.FileName,
		propOffset:   setsTxt.PCode2a,
		itemMaxProps: setsTxt.MaxNumProps,
		lvl:          setsTxt.Level,
	}
	s.minMaxProps = getMinMaxProps(opts, setsTxt.MaxNumProps)
	scramble(s)
}

// Get Set Items Props
func getAllSetItemsProps(d2files d2file.D2Files, props Props) Props {
	p := PropGetter{
		d2files:    d2files,
		props:      props,
		fileName:   setItemsTxt.FileName,
		propOffset: setItemsTxt.Prop1,
		lvl:        setItemsTxt.Lvl,
	}
	return getProps(p)
}

// Randomize Set Items Props
func randomizeSetItemsProps(opts RandomOptions, d2files d2file.D2Files, props BucketedPropsMap, propKeys []string) {
	s := Scrambler{
		opts:         opts,
		d2files:      d2files,
		props:        props,
		propKeys:     propKeys,
		fileName:     setItemsTxt.FileName,
		propOffset:   setItemsTxt.Prop1,
		itemMaxProps: setItemsTxt.MaxNumProps,
		lvl:          setItemsTxt.Lvl,
	}
	s.minMaxProps = getMinMaxProps(opts, setItemsTxt.MaxNumProps)
	scramble(s)
}

// Get RW Props
func getAllRWProps(d2files d2file.D2Files, props Props) Props {
	p := PropGetter{
		d2files:    d2files,
		props:      props,
		fileName:   runesTxt.FileName,
		propOffset: runesTxt.T1Code1,
		lvl:        -1,
	}
	return getProps(p)
}

// Randomize RW Props
func randomizeRWProps(opts RandomOptions, miscBuckets map[string]int, d2files d2file.D2Files, props BucketedPropsMap, propKeys []string) {
	f := d2file.GetOrCreateFile(dataDir, d2files, runesTxt.FileName)
	s := Scrambler{
		opts:         opts,
		d2files:      d2files,
		props:        props,
		propKeys:     propKeys,
		fileName:     runesTxt.FileName,
		propOffset:   runesTxt.T1Code1,
		itemMaxProps: runesTxt.MaxNumProps,
	}
	s.minMaxProps = getMinMaxProps(opts, runesTxt.MaxNumProps)
	for idx, row := range f.Rows {
		runeBuckets := []int{}
		for j := 0; j < 5; j++ {
			runeBuckets = append(runeBuckets, miscBuckets[row[runesTxt.Rune1+j]])
		}
		bucket := getMaxBucket(runeBuckets)
		s.lvl = bucket
		scrambleRow(s, f, idx, row)

	}
}

// Get Gem Props
func getAllGemsProps(d2files d2file.D2Files, props Props) Props {
	f := d2file.GetOrCreateFile(dataDir, d2files, gemsTxt.FileName)
	propOffset := gemsTxt.WeaponMod1Code
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

func getBalancebucket(lvl int) int {
	if lvl > 60 {
		return bucket60
	}
	if lvl > 30 {
		return bucket30
	}
	return bucket0
}

func getBalanceBuckets(lvl int) []int {
	buckets := []int{}
	if lvl > 60 {
		buckets = append(buckets, bucket60)
	} else if lvl > 30 {
		buckets = append(buckets, bucket60)
		buckets = append(buckets, bucket30)
	} else {
		buckets = append(buckets, bucket60)
		buckets = append(buckets, bucket30)
		buckets = append(buckets, bucket0)
	}
	return buckets
}

func getBalancedRandomProp(opts RandomOptions, lvl string, props BucketedPropsMap, propKeys []string) Prop {
	// get our bucket
	bucket := bucketAll
	n, err := strconv.Atoi(lvl)
	if err == nil && opts.IsBalanced {
		bucket = getBalancebucket(n)
	}

	// get prop
	numPropKeys := len(propKeys)
	k := propKeys[randInt(0, numPropKeys)]
	if len(props[k][bucket]) > 0 {
		return props[k][bucket][randInt(0, len(props[k][bucket]))]
	}
	return props[k][bucketAll][randInt(0, len(props[k][bucketAll]))]

}

func getBucketsForMisc(d2files d2file.D2Files) map[string]int {
	f := d2file.GetOrCreateFile(dataDir, d2files, misctxt.FileName)
	buckets := make(map[string]int)
	for _, row := range f.Rows {
		n, err := strconv.Atoi(row[misctxt.Level])
		bucket := bucketAll
		if err == nil {
			bucket = getBalancebucket(n)
		}
		buckets[row[misctxt.Code]] = bucket
	}
	return buckets
}

func getMaxBucket(buckets []int) int {
	bucket := 0
	for i, e := range buckets {
		if i == 0 || e < bucket {
			bucket = e
		}
	}
	return bucket
}

func getAdjustNumProps(opts RandomOptions) bool {
	return opts.MinProps >= 0 || opts.MaxProps >= 0
}

func getMinMaxProps(opts RandomOptions, maxItemProps int) MinMaxProps {
	min := util.MinInt(maxItemProps, util.MaxInt(0, opts.MinProps))
	max := maxItemProps
	if opts.MaxProps > 0 {
		max = util.MinInt(opts.MaxProps, maxItemProps)
	}
	return MinMaxProps{
		minNumProps: min,
		maxNumProps: util.MaxInt(min, max),
	}
}

type Scrambler struct {
	opts         RandomOptions
	d2files      d2file.D2Files
	props        BucketedPropsMap
	propKeys     []string
	fileName     string
	propOffset   int
	minMaxProps  MinMaxProps
	itemMaxProps int
	lvl          int
}

type MinMaxProps struct {
	minNumProps int
	maxNumProps int
}

func scramble(s Scrambler) {
	f := d2file.GetOrCreateFile(dataDir, s.d2files, s.fileName)
	for idx, row := range f.Rows {
		scrambleRow(s, f, idx, row)
	}
}

func scrambleRow(s Scrambler, f *d2file.D2File, idx int, row []string) {
	numProps := randInt(s.minMaxProps.minNumProps, s.minMaxProps.maxNumProps)
	currentNumProps := 0
	// fill in the rest of the props with blanks
	for currentNumProps < s.itemMaxProps {
		prop := Prop{Name: "", Par: "", Min: "", Max: ""}
		for prop.Name == "" && currentNumProps < numProps {
			prop = getBalancedRandomProp(s.opts, row[s.lvl], s.props, s.propKeys)
		}
		i := s.propOffset + currentNumProps*4
		f.Rows[idx][i] = prop.Name
		f.Rows[idx][i+1] = prop.Par
		f.Rows[idx][i+2] = prop.Min
		f.Rows[idx][i+3] = prop.Max
		currentNumProps++
	}
}

func randInt(min int, max int) int {
	if min == max {
		return min
	}
	return min + rand.Intn(max-min)
}
