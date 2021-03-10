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
	EnhancedSets      bool  `json:"EnhancedSets"`      // Propcount & score from Unique Item, not Set Item.
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
	ModName                 string           `json:"ModName"`
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
	UseOSkills              bool             `json:"UseOSkills"`   // +3 Fireball (Sorceress Only) -> +3 Fireball
	PerfectProps            bool             `json:"PerfectProps"` // sets min/max to max
	SafeUnsocket            bool             `json:"SafeUnsocket"`
	PropDebug               bool             `json:"PropDebug"`
	EnterToExit             bool             `json:"EnterToExit"`
	RandomOptions           RandomOptions    `json:"RandomOptions"`
	GeneratorOptions        GeneratorOptions `json:"GeneratorOptions"`
}

// DefaultData Default configuration should the cfg.json not read/be missing anything.
func DefaultData() Data {
	return Data{
		Version:                 "v0.6.0",
		ModName:                 "113c",
		SourceDir:               "",
		OutputDir:               "",
		MeleeSplash:             false,
		IncreaseStackSizes:      false,
		IncreaseMonsterDensity:  1,
		EnableTownSkills:        false,
		BiggerGoldPiles:         false,
		NoFlawGems:              false,
		NoDropZero:              false,
		QuestDrops:              false,
		UniqueItemDropRate:      1,
		RuneDropRate:            1,
		StartWithCube:           true,
		Cowzzz:                  false,
		RemoveLevelRequirements: false,
		RemoveAttRequirements:   false,
		RemoveUniqCharmLimit:    false,
		UseOSkills:              false,
		PerfectProps:            false,
		SafeUnsocket:            false,
		PropDebug:               false,
		EnterToExit:             true,
		RandomOptions: RandomOptions{
			Randomize:         false,
			UseSeed:           false,
			Seed:              -1,
			EnhancedSets:      true,
			IsBalanced:        true,
			BalancedPropCount: true,
			AllowDupeProps:    false,
			MinProps:          3,
			MaxProps:          10,
			NumClones:         9,
			ElementalSkills:   true,
		},
		GeneratorOptions: GeneratorOptions{
			Generate:            false,
			UseSeed:             false,
			Seed:                -1,
			EnhancedSets:        true,
			BalancedPropCount:   true,
			MinProps:            3,
			MaxProps:            10,
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
	data.Version = DefaultData().Version // force version # to current version
	return data
}

// Parse configuration from given json data
func Parse(jsonData []byte) Data {
	data := DefaultData()
	_ = json.Unmarshal(jsonData, &data)
	data.Version = DefaultData().Version // force version # to current version
	return data
}
