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
)

type Prop struct {
	Name string
	Par  string
	Min  string
	Max  string
	Lvl  int
}
type Props = []Prop
type BalancedProps = map[int]Props
type Props2 = map[string]BalancedProps

type RandomOptions struct {
	Randomize  bool  `json:"Randomize"`
	Seed       int64 `json:"Seed"`
	IsBalanced bool  `json:"Balance"`
}

const (
	bucketAll = 0
	bucket0   = 1
	bucket30  = 2
	bucket60  = 3
)

func getRandomOptions(cfg *ModConfig) RandomOptions {
	defaultCfg := RandomOptions{
		Seed: time.Now().UnixNano(),
	}
	if cfg.RandomOptions.Seed > 0 {
		defaultCfg.Seed = cfg.RandomOptions.Seed
	}
	defaultCfg.IsBalanced = cfg.RandomOptions.IsBalanced
	return defaultCfg
}

func Randomize(cfg *ModConfig, d2files *d2file.D2Files) {
	opts := getRandomOptions(cfg)
	rand.Seed(opts.Seed)

	props := getAllProps(opts, d2files)
	miscBuckets := getBucketsForMisc(d2files)

	randomizeUniqueProps(opts, d2files, props)
	randomizeSetProps(opts, d2files, props)
	randomizeSetItemsProps(opts, d2files, props)
	randomizeRWProps(opts, miscBuckets, d2files, props)
}

func getAllProps(opts RandomOptions, d2files *d2file.D2Files) BalancedProps {
	props := BalancedProps{}
	props[bucketAll] = Props{}
	props[bucket0] = Props{}
	props[bucket30] = Props{}
	props[bucket60] = Props{}

	// uniques
	uniqueProps := getAllUniqueProps(d2files, []Prop{})
	if opts.IsBalanced {
		for _, prop := range uniqueProps {
			bucket := getBalancebucket(prop.Lvl)
			props[bucket] = append(props[bucket], prop)
		}
	}
	props[bucketAll] = append(props[bucketAll], uniqueProps...)

	// sets
	setProps := getAllSetProps(d2files, []Prop{})
	if opts.IsBalanced {
		for _, prop := range setProps {
			bucket := getBalancebucket(prop.Lvl)
			props[bucket] = append(props[bucket], prop)
		}
	}
	props[bucketAll] = append(props[bucketAll], setProps...)

	// sets items
	setItemsProps := getAllSetItemsProps(d2files, []Prop{})
	if opts.IsBalanced {
		for _, prop := range setItemsProps {
			bucket := getBalancebucket(prop.Lvl)
			props[bucket] = append(props[bucket], prop)
		}
	}
	props[bucketAll] = append(props[bucketAll], setProps...)

	// rw
	rwProps := getAllRWProps(d2files, []Prop{})
	if opts.IsBalanced {
		for _, prop := range rwProps {
			bucket := getBalancebucket(prop.Lvl)
			props[bucket] = append(props[bucket], prop)
		}
	}
	props[bucketAll] = append(props[bucketAll], setProps...)

	// gems
	gemsProps := getAllGemsProps(d2files, []Prop{})
	if opts.IsBalanced {
		for _, prop := range gemsProps {
			bucket := getBalancebucket(prop.Lvl)
			props[bucket] = append(props[bucket], prop)
		}
	}
	props[bucketAll] = append(props[bucketAll], setProps...)

	return props
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
		for i := propOffset; i < len(row)-propOffset; i += 4 {
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

// Randomize Unique Props
func randomizeUniqueProps(opts RandomOptions, d2files *d2file.D2Files, props BalancedProps) {
	f := d2file.GetOrCreateFile(dataDir, d2files, uniqueItemsTxt.FileName)
	propOffset := uniqueItemsTxt.Prop1
	for idx, row := range f.Rows {
		for i := propOffset; i < len(row)-propOffset; i += 4 {
			prop := getBalancedRandomProp(opts, row[uniqueItemsTxt.Lvl], props)
			if row[i] != "" {
				f.Rows[idx][i] = prop.Name
				f.Rows[idx][i+1] = prop.Par
				f.Rows[idx][i+2] = prop.Min
				f.Rows[idx][i+3] = prop.Max
			}
		}
	}
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
		for i := propOffset; i < len(row)-propOffset; i += 4 {
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

// Randomize Set Props
func randomizeSetProps(opts RandomOptions, d2files *d2file.D2Files, props BalancedProps) {
	f := d2file.GetOrCreateFile(dataDir, d2files, setsTxt.FileName)
	propOffset := setsTxt.PCode2a
	for idx, row := range f.Rows {
		for i := propOffset; i < len(row)-propOffset; i += 4 {
			prop := getBalancedRandomProp(opts, row[setsTxt.Level], props)
			if row[i] != "" {
				f.Rows[idx][i] = prop.Name
				f.Rows[idx][i+1] = prop.Par
				f.Rows[idx][i+2] = prop.Min
				f.Rows[idx][i+3] = prop.Max
			}
		}
	}
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
		for i := propOffset; i < len(row)-propOffset; i += 4 {
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
func randomizeSetItemsProps(opts RandomOptions, d2files *d2file.D2Files, props BalancedProps) {
	f := d2file.GetOrCreateFile(dataDir, d2files, setItemsTxt.FileName)
	propOffset := setItemsTxt.Prop1
	for idx, row := range f.Rows {
		for i := propOffset; i < len(row)-propOffset; i += 4 {
			prop := getBalancedRandomProp(opts, row[setItemsTxt.Lvl], props)
			if row[i] != "" {
				f.Rows[idx][i] = prop.Name
				f.Rows[idx][i+1] = prop.Par
				f.Rows[idx][i+2] = prop.Min
				f.Rows[idx][i+3] = prop.Max
			}
		}
	}
}

// Get RW Props
func getAllRWProps(d2files *d2file.D2Files, props Props) Props {
	f := d2file.GetOrCreateFile(dataDir, d2files, runesTxt.FileName)
	propOffset := runesTxt.T1Code1
	for _, row := range f.Rows {
		for i := propOffset; i < len(row)-propOffset; i += 4 {
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
func randomizeRWProps(opts RandomOptions, miscBuckets map[string]int, d2files *d2file.D2Files, props BalancedProps) {
	f := d2file.GetOrCreateFile(dataDir, d2files, runesTxt.FileName)
	propOffset := runesTxt.T1Code1
	for idx, row := range f.Rows {
		runeBuckets := []int{}
		for j := 0; j < 6; j++ {
			runeBuckets = append(runeBuckets, miscBuckets[row[runesTxt.Rune1+j]])
		}
		bucket := getMaxBucket(runeBuckets)
		for i := propOffset; i < len(row)-propOffset; i += 4 {
			prop := getRandomProp(opts, bucket, props)
			if row[i] != "" {
				f.Rows[idx][i] = prop.Name
				f.Rows[idx][i+1] = prop.Par
				f.Rows[idx][i+2] = prop.Min
				f.Rows[idx][i+3] = prop.Max
			}
		}
	}
}

// Get Gem Props
func getAllGemsProps(d2files *d2file.D2Files, props Props) Props {
	f := d2file.GetOrCreateFile(dataDir, d2files, gemsTxt.FileName)
	propOffset := gemsTxt.WeaponMod1Code
	for _, row := range f.Rows {
		for i := propOffset; i < len(row)-propOffset; i += 4 {
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

// Randomize Gem Props
func randomizeGemsProps(opts RandomOptions, miscBuckets map[string]int, d2files *d2file.D2Files, props BalancedProps) {
	f := d2file.GetOrCreateFile(dataDir, d2files, gemsTxt.FileName)
	propOffset := gemsTxt.WeaponMod1Code
	for idx, row := range f.Rows {
		for i := propOffset; i < len(row)-propOffset; i += 4 {
			prop := getRandomProp(opts, miscBuckets[row[gemsTxt.Name]], props)
			for ok := false; ok; ok = prop.Par == "" {
				if prop.Par == "" {
					break
				}
				prop = getRandomProp(opts, miscBuckets[row[gemsTxt.Name]], props)
			}

			if row[i] != "" {
				f.Rows[idx][i] = prop.Name
				f.Rows[idx][i+1] = prop.Par
				f.Rows[idx][i+2] = prop.Max // min/max need to be the same values
				f.Rows[idx][i+3] = prop.Max
			}
		}
	}
}

func getRandomProp(opts RandomOptions, bucket int, props BalancedProps) Prop {
	if opts.IsBalanced {
		return props[bucket][rand.Intn(len(props[bucket]))]
	}
	return props[bucketAll][rand.Intn(len(props[bucketAll]))]
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

func getBalancedRandomProp(opts RandomOptions, lvl string, props BalancedProps) Prop {
	n, err := strconv.Atoi(lvl)
	if err == nil {
		bucket := getBalancebucket(n)
		return getRandomProp(opts, bucket, props)
	}
	return getRandomProp(opts, bucketAll, props)
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
		buckets[row[misctxt.Name]] = bucket
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

func maxBucket(a, b string) string {
	if bucketToInt(a) > bucketToInt(b) {
		return a
	}
	return b
}

func bucketToInt(x string) int {
	if x == "0-30" {
		return 0
	}
	if x == "31-60" {
		return 1
	}
	return 2
}
