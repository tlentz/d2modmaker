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
	IsBalanced        bool  `json:"IsBalanced"`        // Allows Props only from items up to 10 levels higher
	BalancedPropCount bool  `json:"BalancedPropCount"` // Picks prop count from a vanilla item up to 10 levels higher
	AllowDupeProps    bool  `json:"AllowDupeProps"`    // Allow two props of the same type to be placed on an item
	MinProps          int   `json:"MinProps"`          // minimum number of non blank props on an item
	MaxProps          int   `json:"MaxProps"`          // maximum number of non blank props on an item
	UseOSkills        bool  `json:"UseOSkills"`        // +3 Fireball (Sorceress Only) -> +3 Fireball
	PerfectProps      bool  `json:"PerfectProps"`      // sets min/max to max
	ElementalSkills   bool  `json:"ElementalSkills"`
}

// data is the configuration used to build the mod
type Data struct {
	Version                 string        `json:"Version"`
	SourceDir               string        `json:"SourceDir"`
	OutputDir               string        `json:"OutputDir"`
	MeleeSplash             bool          `json:"MeleeSplash"`
	IncreaseStackSizes      bool          `json:"IncreaseStackSizes"`
	IncreaseMonsterDensity  float64       `json:"IncreaseMonsterDensity"`
	EnableTownSkills        bool          `json:"EnableTownSkills"`
	BiggerGoldPiles         bool          `json:"BiggerGoldPiles"`
	NoFlawGems              bool          `json:"NoFlawGems"`
	NoDropZero              bool          `json:"NoDropZero"`
	QuestDrops              bool          `json:"QuestDrops"`
	UniqueItemDropRate      float64       `json:"UniqueItemDropRate"`
	RuneDropRate            float64       `json:"RuneDropRate"`
	StartWithCube           bool          `json:"StartWithCube"`
	Cowzzz                  bool          `json:"Cowzzz"`
	RemoveLevelRequirements bool          `json:"RemoveLevelRequirements"`
	RemoveAttRequirements   bool          `json:"RemoveAttRequirements"`
	RemoveUniqCharmLimit    bool          `json:"RemoveUniqCharmLimit"`
	SafeUnsocket            bool          `json:"SafeUnsocket"`
	EnterToExit             bool          `json:"EnterToExit"`
	RandomOptions           RandomOptions `json:"RandomOptions"`
}

func DefaultData() Data {
	return Data{
		Version:                 "v0.5.4",
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
		SafeUnsocket:            false,
		EnterToExit:             true,
		RandomOptions: RandomOptions{
			Randomize:         false,
			UseSeed:           false,
			Seed:              -1,
			IsBalanced:        true,
			BalancedPropCount: true,
			AllowDupeProps:    false,
			MinProps:          3,
			MaxProps:          8,
			UseOSkills:        false,
			PerfectProps:      false,
		},
	}
}

// Read configuration from the given json file
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
