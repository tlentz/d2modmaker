package main

import (
	"encoding/json"
	"io/ioutil"
)

type ModConfig struct {
	IncreaseStackSizes     bool `json:"IncreaseStackSizes"`
	IncreaseMonsterDensity int  `json:"IncreaseMonsterDensity"`
	LinearRuneDrops        bool `json:"LinearRuneDrops"`
	EnableTownSkills       bool `json:"EnableTownSkills"`
}

func ReadCfg() ModConfig {
	file, _ := ioutil.ReadFile("cfg.json")
	data := ModConfig{}
	_ = json.Unmarshal([]byte(file), &data)
	return data
}
