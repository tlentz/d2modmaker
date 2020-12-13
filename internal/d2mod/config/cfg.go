package config

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/tlentz/d2modmaker/internal/util"
)

// RandomOptions are the options for the randomizer
type RandomOptions struct {
	Randomize           bool    `json:"Randomize"`
	UseSeed             bool    `json:"UseSeed"`
	Seed                int64   `json:"Seed"`
	IsBalanced          bool    `json:"IsBalanced"`          // Allows Props only from items up to 10 levels higher
	BalancedPropCount   bool    `json:"BalancedPropCount"`   // Picks prop count from a vanilla item up to 10 levels higher
	AllowDupProps       bool    `json:"AllowDuplicateProps"` // Allow two props of the same type to be placed on an item
	MinProps            int     `json:"MinProps"`            // minimum number of non blank props on an item
	MaxProps            int     `json:"MaxProps"`            // maximum number of non blank props on an item
	PerfectProps        bool    `json:"PerfectProps"`        // sets min/max to max
	UseOSkills          bool    `json:"UseOSkills"`          // +3 Fireball (Sorceress Only) -> +3 Fireball
	NumClones           int     `json:"NumClones"`           // # of times to copy all rows in UniqueItems table before randomizing props
	UsePropScores       bool    `json:"UsePropScores"`       // Use Scorer/Generator not Randomizer/scramble
	PropScoreMultiplier float64 `json:"PropScoreMultiplier"` // Multiplier against the vanilla prop score.  > 1 better item, < 1 worse
}

// Data is the configuration used to build the mod
type Data struct {
	Version                 string        `json:"Version"`
	SourceDir               string        `json:"SourceDir"`
	OutputDir               string        `json:"OutputDir"`
	MeleeSplash             bool          `json:"MeleeSplash"`
	IncreaseStackSizes      bool          `json:"IncreaseStackSizes"`
	IncreaseMonsterDensity  float64       `json:"IncreaseMonsterDensity"`
	EnableTownSkills        bool          `json:"EnableTownSkills"`
	NoDropZero              bool          `json:"NoDropZero"`
	QuestDrops              bool          `json:"QuestDrops"`
	UniqueItemDropRate      float64       `json:"UniqueItemDropRate"`
	RuneDropRate            float64       `json:"RuneDropRate"`
	StartWithCube           bool          `json:"StartWithCube"`
	Cowzzz                  bool          `json:"Cowzzz"`
	RemoveLevelRequirements bool          `json:"RemoveLevelRequirements"`
	RemoveAttRequirements   bool          `json:"RemoveAttRequirements"`
	RemoveUniqCharmLimit    bool          `json:"RemoveUniqCharmLimit"`
	EnterToExit             bool          `json:"EnterToExit"`
	RandomOptions           RandomOptions `json:"RandomOptions"`
}

func DefaultData() Data {
	return Data{
		Version:                 "v0.5.2-alpha-8",
		SourceDir:               "",
		OutputDir:               "",
		MeleeSplash:             true,
		IncreaseStackSizes:      true,
		IncreaseMonsterDensity:  1,
		EnableTownSkills:        true,
		NoDropZero:              true,
		QuestDrops:              true,
		UniqueItemDropRate:      1,
		RuneDropRate:            1,
		StartWithCube:           true,
		Cowzzz:                  true,
		RemoveLevelRequirements: false,
		RemoveAttRequirements:   false,
		RemoveUniqCharmLimit:    false,
		EnterToExit:             false,
		RandomOptions: RandomOptions{
			Randomize:           true,
			UseSeed:             false,
			Seed:                -1,
			IsBalanced:          true,
			BalancedPropCount:   true,
			AllowDupProps:       false,
			MinProps:            2,
			MaxProps:            20,
			PerfectProps:        false,
			UseOSkills:          true,
			NumClones:           9,
			UsePropScores:       true,
			PropScoreMultiplier: 1, // 1 == Vanilla
		},
	}
}

// Read reads a ModConfig from the given json file
func Read(filePath string) Data {
	file, _ := ioutil.ReadFile(filePath)
	data := DefaultData()
	_ = json.Unmarshal([]byte(file), &data)
	data.audit()
	return data
}

func Parse(jsonData []byte) Data {
	data := DefaultData()
	_ = json.Unmarshal(jsonData, &data)
	data.audit()
	return data
}

func (d *Data) audit() {
	if !d.RandomOptions.UseSeed {
		d.RandomOptions.Seed = time.Now().UnixNano()
	}
	d.RandomOptions.NumClones = util.MaxInt(0, d.RandomOptions.NumClones)
}
