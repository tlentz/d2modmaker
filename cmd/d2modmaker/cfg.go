package main

import (
	"encoding/json"
	"io/ioutil"
)

// ModConfig is the config used to build the mod
type ModConfig struct {
	IncreaseStackSizes     bool `json:"IncreaseStackSizes"`
	IncreaseMonsterDensity int  `json:"IncreaseMonsterDensity"`
	EnableTownSkills       bool `json:"EnableTownSkills"`
	NoDropZero             bool `json:"NoDropZero"`
	QuestDrops             bool `json:"QuestDrops"`
}

// ReadCfg reads a ModConfig from the given json file
func ReadCfg(filePath string) ModConfig {
	file, _ := ioutil.ReadFile(filePath)
	data := ModConfig{}
	_ = json.Unmarshal([]byte(file), &data)
	return data
}