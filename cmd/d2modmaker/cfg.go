package main

import (
	"encoding/json"
	"io/ioutil"
)

// ModConfig is the config used to build the mod
type ModConfig struct {
	IncreaseStackSizes     bool          `json:"IncreaseStackSizes"`
	IncreaseMonsterDensity float64       `json:"IncreaseMonsterDensity"`
	EnableTownSkills       bool          `json:"EnableTownSkills"`
	NoDropZero             bool          `json:"NoDropZero"`
	QuestDrops             bool          `json:"QuestDrops"`
	UniqueItemDropRate     float64       `json:"UniqueItemDropRate"`
	StartWithCube          bool          `json:"StartWithCube"`
	Cowzzz                 bool          `json:"Cowzzz"`
	EnterToExit            bool          `json:"EnterToExit"`
	RandomOptions          RandomOptions `json:"RandomOptions"`
}

// ReadCfg reads a ModConfig from the given json file
func ReadCfg(filePath string) ModConfig {
	file, _ := ioutil.ReadFile(filePath)
	data := ModConfig{}
	_ = json.Unmarshal([]byte(file), &data)
	return data
}
