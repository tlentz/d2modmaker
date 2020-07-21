package modcfg

import (
	"encoding/json"
	"io/ioutil"
)

// ModConfig is the config used to build the mod
type ModConfig struct {
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
	PathToDataDir          string        `json:"PathToDataDir"`
	RandomOptions          RandomOptions `json:"RandomOptions"`
}

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

// ReadCfg reads a ModConfig from the given json file
func ReadCfg(filePath string) ModConfig {
	file, _ := ioutil.ReadFile(filePath)
	data := ModConfig{}
	_ = json.Unmarshal([]byte(file), &data)
	return data
}
