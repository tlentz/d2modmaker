package config

import (
	"encoding/json"
	"io/ioutil"
)

// RandomOptions are the options for the randomizer
type RandomOptions struct {
	Randomize         bool  `json:"Randomize"`
	UseSeed           bool  `json:"UseSeed"`
	Seed              int64 `json:"Seed"`
	UseSetsSeed       bool  `json:"UseSetsSeed"`
	SetsSeed          int64 `json:"SetsSeed"`
	IsBalanced        bool  `json:"IsBalanced"`        // Allows Props only from items up to 10 levels higher
	BalancedPropCount bool  `json:"BalancedPropCount"` // Picks prop count from a vanilla item up to 10 levels higher
	AllowDupeProps    bool  `json:"AllowDupeProps"`    // Allow two props of the same type to be placed on an item
	MinProps          int   `json:"MinProps"`          // minimum number of non blank props on an item
	MaxProps          int   `json:"MaxProps"`          // maximum number of non blank props on an item
	NumClones         int   `json:"NumClones"`         // # of times to copy all rows in UniqueItems table before randomizing props
	ElementalSkills   bool  `json:"ElementalSkills"`
}

// GeneratorOptions control the Scorer/Generator combo
type GeneratorOptions struct {
	Generate            bool    `json:"Generate"` // Turn On/Off Generator
	UseSeed             bool    `json:"UseSeed"`
	Seed                int64   `json:"Seed"`
	UseSetsSeed         bool    `json:"UseSetsSeed"`
	SetsSeed            int64   `json:"SetsSeed"`            // seed forSetItems.txt
	EnhancedSets        bool    `json:"EnhancedSets"`        // Propcount & score from Unique Item, not Set Item.
	BalancedPropCount   bool    `json:"BalancedPropCount"`   // Picks prop count from a vanilla item up to 10 levels higher
	MinProps            int     `json:"MinProps"`            // minimum number of non blank props on an item
	MaxProps            int     `json:"MaxProps"`            // maximum number of non blank props on an item
	NumClones           int     `json:"NumClones"`           // # of clones to generate in UniqueItems table
	PropScoreMultiplier float64 `json:"PropScoreMultiplier"` // Multiplier against the vanilla prop score.  > 1 better item, < 1 worse
	ElementalSkills     bool    `json:"ElementalSkills"`
}

// Data is the configuration used to build the mod
type Data struct {
	Version                 string           `json:"Version"`
	SourceDir               string           `json:"SourceDir"`
	OutputDir               string           `json:"OutputDir"`
	MeleeSplash             bool             `json:"MeleeSplash"`
	IncreaseStackSizes      bool             `json:"IncreaseStackSizes"`
	IncreaseMonsterDensity  float64          `json:"IncreaseMonsterDensity"`
	EnableTownSkills        bool             `json:"EnableTownSkills"`
	BiggerGoldPiles         bool             `json:"BiggerGoldPiles"`
	NoFlawGems              bool             `json:"NoFlawGems"`
	NoDropZero              bool             `json:"NoDropZero"`
	QuestDrops              bool             `json:"QuestDrops"`
	UniqueItemDropRate      float64          `json:"UniqueItemDropRate"`
	RuneDropRate            float64          `json:"RuneDropRate"`
	StartWithCube           bool             `json:"StartWithCube"`
	Cowzzz                  bool             `json:"Cowzzz"`
	RemoveLevelRequirements bool             `json:"RemoveLevelRequirements"`
	RemoveAttRequirements   bool             `json:"RemoveAttRequirements"`
	RemoveUniqCharmLimit    bool             `json:"RemoveUniqCharmLimit"`
	PerfectProps            bool             `json:"PerfectProps"` // sets min/max to max
	UseOSkills              bool             `json:"UseOSkills"`   // +3 Fireball (Sorceress Only) -> +3 Fireball
	SafeUnsocket            bool             `json:"SafeUnsocket"`
	EnterToExit             bool             `json:"EnterToExit"`
	RandomOptions           RandomOptions    `json:"RandomOptions"`
	GeneratorOptions        GeneratorOptions `json:"GeneratorOptions"`
}

// DefaultData Default configuration should the cfg.json not read/be missing anything.
func DefaultData() Data {
	return Data{
		Version:                 "v0.5.2-alpha-13",
		SourceDir:               "",
		OutputDir:               "",
		MeleeSplash:             true,
		IncreaseStackSizes:      true,
		IncreaseMonsterDensity:  1,
		EnableTownSkills:        true,
		BiggerGoldPiles:         true,
		NoFlawGems:              true,
		NoDropZero:              false,
		QuestDrops:              true,
		UniqueItemDropRate:      1,
		RuneDropRate:            1,
		StartWithCube:           true,
		Cowzzz:                  true,
		RemoveLevelRequirements: false,
		RemoveAttRequirements:   false,
		RemoveUniqCharmLimit:    false,
		PerfectProps:            false,
		UseOSkills:              true,
		SafeUnsocket:            true,
		EnterToExit:             false,
		RandomOptions: RandomOptions{
			Randomize:         false,
			UseSeed:           false,
			Seed:              -1,
			UseSetsSeed:       true,
			SetsSeed:          1234,
			IsBalanced:        true,
			BalancedPropCount: true,
			AllowDupeProps:    false,
			MinProps:          2,
			MaxProps:          20,
			NumClones:         9,
			ElementalSkills:   true,
		},
		GeneratorOptions: GeneratorOptions{
			Generate:            true,
			UseSeed:             false,
			Seed:                -1,
			UseSetsSeed:         true,
			SetsSeed:            1234,
			EnhancedSets:        true,
			BalancedPropCount:   true,
			MinProps:            2,
			MaxProps:            20,
			NumClones:           9,
			PropScoreMultiplier: 1, // 1 == Vanilla
			ElementalSkills:     true,
		},
	}
}

// Read reads a ModConfig from the given json file
func Read(filePath string) Data {
	file, _ := ioutil.ReadFile(filePath)
	data := DefaultData()
	_ = json.Unmarshal([]byte(file), &data)
	return data
}

// Parse configuration from given json data
func Parse(jsonData []byte) Data {
	data := DefaultData()
	_ = json.Unmarshal(jsonData, &data)
	return data
}
