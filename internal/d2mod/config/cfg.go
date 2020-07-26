package config

import (
	"encoding/json"
	"io/ioutil"
)

// RandomOptions are the options for the randomizer
type RandomOptions struct {
	Randomize         bool  `json:"Randomize"`
	Seed              int64 `json:"Seed"`
	IsBalanced        bool  `json:"IsBalanced"`          // Allows Props only from items up to 10 levels higher
	BalancedPropCount bool  `json:"BalancedPropCount"`   // Picks prop count from a vanilla item up to 10 levels higher
	AllowDupProps     bool  `json:"AllowDuplicateProps"` // Allow two props of the same type to be placed on an item
	MinProps          int   `json:"MinProps"`            // minimum number of non blank props on an item
	MaxProps          int   `json:"MaxProps"`            // maximum number of non blank props on an item
	PerfectProps      bool  `json:"PerfectProps"`        // sets min/max to max
	UseOSkills        bool  `json:"UseOSkills"`          // +3 Fireball (Sorceress Only) -> +3 Fireball
}

// data is the configuration used to build the mod
type Data struct {
	SourceDir              string        `json:"SourceDir"`
	MeleeSplash            bool          `json:"MeleeSplash"`
	IncreaseStackSizes     bool          `json:"IncreaseStackSizes"`
	IncreaseMonsterDensity float64       `json:"IncreaseMonsterDensity"`
	EnableTownSkills       bool          `json:"EnableTownSkills"`
	NoDropZero             bool          `json:"NoDropZero"`
	QuestDrops             bool          `json:"QuestDrops"`
	UniqueItemDropRate     float64       `json:"UniqueItemDropRate"`
	RuneDropRate           float64       `json:"RuneDropRate"`
	StartWithCube          bool          `json:"StartWithCube"`
	Cowzzz                 bool          `json:"Cowzzz"`
	EnterToExit            bool          `json:"EnterToExit"`
	RandomOptions          RandomOptions `json:"RandomOptions"`
}

// ReadCfg reads a ModConfig from the given json file
func Read(filePath string) Data {
	file, _ := ioutil.ReadFile(filePath)
	data := Data{}
	_ = json.Unmarshal([]byte(file), &data)
	return data
}
