package main

import (
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
func Randomize(cfg *ModConfig, d2files *d2file.D2Files) {
	opts := getRandomOptions(cfg)
	rand.Seed(opts.Seed)

	props, propKeys := getAllProps(opts, d2files)
	miscBuckets := getBucketsForMisc(d2files)

	randomizeUniqueProps(opts, d2files, props, propKeys)
	randomizeSetProps(opts, d2files, props, propKeys)
	randomizeSetItemsProps(opts, d2files, props, propKeys)
	randomizeRWProps(opts, miscBuckets, d2files, props, propKeys)
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
func getAllProps(opts RandomOptions, d2files *d2file.D2Files) (BucketedPropsMap, []string) {

	props := BucketedPropsMap{}

	// uniques
	uniqueProps := getAllUniqueProps(d2files, []Prop{})
	for _, prop := range uniqueProps {
		props = addOrCreateProp(props, prop)
	}

	// sets
	setProps := getAllSetProps(d2files, []Prop{})
	for _, prop := range setProps {
		props = addOrCreateProp(props, prop)
	}

	// sets items
	setItemsProps := getAllSetItemsProps(d2files, []Prop{})
	for _, prop := range setItemsProps {
		props = addOrCreateProp(props, prop)
	}

	// rw
	rwProps := getAllRWProps(d2files, []Prop{})
	for _, prop := range rwProps {
		props = addOrCreateProp(props, prop)
	}

	// gems
	gemsProps := getAllGemsProps(d2files, []Prop{})
	for _, prop := range gemsProps {
		props = addOrCreateProp(props, prop)
	}

	for k := range props {
		for b := range props[k] {
			for i, p := range props[k][b] {
				// Set all props Min to the Max value
				if opts.PerfectProps {
					props[k][b][i].Min = p.Max
				}
				// sets skill = oskill
				if opts.UseOSkills {
					if p.Name == "skill" {
						props[k][b][i].Name = "oskill"
					}
				}
			}
		}
	}

	var keys []string
	for k := range props {
		totalProps := 0
		for b := range props[k] {
			totalProps += len(props[k][b])
		}
		for i := 0; i < totalProps; i++ {
			keys = append(keys, k)
		}
	}
	return props, keys
}

// Get Unique Props
func getAllUniqueProps(d2files *d2file.D2Files, props Props) Props {
	f := d2file.GetOrCreateFile(dataDir, d2files, uniqueItemsTxt.FileName)
	propOffset := uniqueItemsTxt.Prop1
	for _, row := range f.Rows {
		mbLvl, err := strconv.Atoi(row[uniqueItemsTxt.Lvl])
		lvl := 0
		if err == nil {
			lvl = mbLvl
		}
		for i := propOffset; i < len(row)-3; i += 4 {
			props = append(props, Prop{
				Name: row[i],
				Par:  row[i+1],
				Min:  row[i+2],
				Max:  row[i+3],
				Lvl:  lvl,
			})
		}
	}
	return props
}

// Randomize Unique Props
func randomizeUniqueProps(opts RandomOptions, d2files *d2file.D2Files, props BucketedPropsMap, propKeys []string) {
	s := Scrambler{
		opts:           opts,
		d2files:        d2files,
		props:          props,
		propKeys:       propKeys,
		fileName:       uniqueItemsTxt.FileName,
		propOffset:     uniqueItemsTxt.Prop1,
		adjustNumProps: getAdjustNumProps(opts),
		minMaxProps:    getMinMaxProps(opts, uniqueItemsTxt.MaxNumProps),
		lvl:            uniqueItemsTxt.Lvl,
	}
	scramble(s)
}

// Get Set Props
func getAllSetProps(d2files *d2file.D2Files, props Props) Props {
	f := d2file.GetOrCreateFile(dataDir, d2files, setsTxt.FileName)
	propOffset := setsTxt.PCode2a
	for _, row := range f.Rows {
		mbLvl, err := strconv.Atoi(row[setsTxt.Level])
		lvl := 0
		if err == nil {
			lvl = mbLvl
		}
		for i := propOffset; i < len(row)-3; i += 4 {
			props = append(props, Prop{
				Name: row[i],
				Par:  row[i+1],
				Min:  row[i+2],
				Max:  row[i+3],
				Lvl:  lvl,
			})
		}
	}
	return props
}

// Randomize Set Props
func randomizeSetProps(opts RandomOptions, d2files *d2file.D2Files, props BucketedPropsMap, propKeys []string) {
	s := Scrambler{
		opts:           opts,
		d2files:        d2files,
		props:          props,
		propKeys:       propKeys,
		fileName:       setsTxt.FileName,
		propOffset:     setsTxt.PCode2a,
		adjustNumProps: getAdjustNumProps(opts),
		minMaxProps:    getMinMaxProps(opts, setsTxt.MaxNumProps),
		lvl:            setsTxt.Level,
	}
	scramble(s)
}

// Get Set Items Props
func getAllSetItemsProps(d2files *d2file.D2Files, props Props) Props {
	f := d2file.GetOrCreateFile(dataDir, d2files, setItemsTxt.FileName)
	propOffset := setItemsTxt.Prop1
	for _, row := range f.Rows {
		mbLvl, err := strconv.Atoi(row[setItemsTxt.Lvl])
		lvl := 0
		if err == nil {
			lvl = mbLvl
		}
		for i := propOffset; i < len(row)-3; i += 4 {
			if row[i] != "" {
				props = append(props, Prop{
					Name: row[i],
					Par:  row[i+1],
					Min:  row[i+2],
					Max:  row[i+3],
					Lvl:  lvl,
				})
			}
		}
	}
	return props
}

// Randomize Set Items Props
func randomizeSetItemsProps(opts RandomOptions, d2files *d2file.D2Files, props BucketedPropsMap, propKeys []string) {
	s := Scrambler{
		opts:           opts,
		d2files:        d2files,
		props:          props,
		propKeys:       propKeys,
		fileName:       setItemsTxt.FileName,
		propOffset:     setItemsTxt.Prop1,
		adjustNumProps: getAdjustNumProps(opts),
		minMaxProps:    getMinMaxProps(opts, setItemsTxt.MaxNumProps),
		lvl:            setItemsTxt.Lvl,
	}
	scramble(s)
}

// Get RW Props
func getAllRWProps(d2files *d2file.D2Files, props Props) Props {
	f := d2file.GetOrCreateFile(dataDir, d2files, runesTxt.FileName)
	propOffset := runesTxt.T1Code1
	for _, row := range f.Rows {
		for i := propOffset; i < len(row)-3; i += 4 {
			if row[i] != "" {
				props = append(props, Prop{
					Name: row[i],
					Par:  row[i+1],
					Min:  row[i+2],
					Max:  row[i+3],
					Lvl:  0,
				})
			}
		}
	}
	return props
}

// Randomize RW Props
func randomizeRWProps(opts RandomOptions, miscBuckets map[string]int, d2files *d2file.D2Files, props BucketedPropsMap, propKeys []string) {
	f := d2file.GetOrCreateFile(dataDir, d2files, runesTxt.FileName)
	s := Scrambler{
		opts:           opts,
		d2files:        d2files,
		props:          props,
		propKeys:       propKeys,
		fileName:       runesTxt.FileName,
		propOffset:     runesTxt.T1Code1,
		adjustNumProps: getAdjustNumProps(opts),
		minMaxProps:    getMinMaxProps(opts, runesTxt.MaxNumProps),
		lvl:            0,
	}
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
func getAllGemsProps(d2files *d2file.D2Files, props Props) Props {
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
	}
	if lvl > 30 {
		buckets = append(buckets, bucket30)
	}
	buckets = append(buckets, bucket0)
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
	k := propKeys[rand.Intn(numPropKeys)]
	if len(props[k][bucket]) > 0 {
		return props[k][bucket][rand.Intn(len(props[k][bucket]))]
	}
	return props[k][bucketAll][rand.Intn(len(props[k][bucketAll]))]

}

func getBucketsForMisc(d2files *d2file.D2Files) map[string]int {
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
	min := util.MinInt(util.MaxInt(0, opts.MinProps), maxItemProps)
	max := util.MaxInt(util.MinInt(maxItemProps, opts.MaxProps), 0)
	a := MinMaxProps{
		minNumProps: min,
		maxNumProps: util.MaxInt(min, max),
	}
	return a
}

type Scrambler struct {
	opts           RandomOptions
	d2files        *d2file.D2Files
	props          BucketedPropsMap
	propKeys       []string
	fileName       string
	propOffset     int
	adjustNumProps bool
	minMaxProps    MinMaxProps
	lvl            int
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
	numProps := 0
	for i := s.propOffset; i < len(row)-3; i += 4 {
		prop := getBalancedRandomProp(s.opts, row[s.lvl], s.props, s.propKeys)
		if prop.Name == "" && numProps < s.minMaxProps.minNumProps && numProps < s.minMaxProps.maxNumProps && s.adjustNumProps {
			i -= 4
		} else if numProps >= s.minMaxProps.minNumProps && numProps >= s.minMaxProps.maxNumProps && s.adjustNumProps {
			f.Rows[idx][i] = ""
			f.Rows[idx][i+1] = ""
			f.Rows[idx][i+2] = ""
			f.Rows[idx][i+3] = ""
			numProps++
		} else {
			f.Rows[idx][i] = prop.Name
			f.Rows[idx][i+1] = prop.Par
			f.Rows[idx][i+2] = prop.Min
			f.Rows[idx][i+3] = prop.Max
			numProps++
		}
	}
}
