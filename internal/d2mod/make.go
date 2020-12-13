package d2mod

import (
	"fmt"
	"os"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2mod/config"
	"github.com/tlentz/d2modmaker/internal/d2mod/cows"
	"github.com/tlentz/d2modmaker/internal/d2mod/generator"
	"github.com/tlentz/d2modmaker/internal/d2mod/monsterdensity"
	"github.com/tlentz/d2modmaker/internal/d2mod/oskills"
	"github.com/tlentz/d2modmaker/internal/d2mod/qol"
	"github.com/tlentz/d2modmaker/internal/d2mod/randomizer"
	"github.com/tlentz/d2modmaker/internal/d2mod/reqs"
	"github.com/tlentz/d2modmaker/internal/d2mod/scorer"
	"github.com/tlentz/d2modmaker/internal/d2mod/splash"
	"github.com/tlentz/d2modmaker/internal/d2mod/stacksizes"
	"github.com/tlentz/d2modmaker/internal/d2mod/townskills"
	"github.com/tlentz/d2modmaker/internal/d2mod/treasure"
	"github.com/tlentz/d2modmaker/internal/util"
)

//MakeFromCfgPath ??
func MakeFromCfgPath(defaultOutDir string, cfgPath string) {
	cfg := config.Read(cfgPath)
	Make(defaultOutDir, cfg)
}

// Make Run all the enabled d2 modules
func Make(defaultOutDir string, cfg config.Data) {
	if cfg.OutputDir == "" {
		cfg.OutputDir = defaultOutDir
	}
	d2files := d2fs.NewFiles(cfg.SourceDir, cfg.OutputDir)

	if cfg.MeleeSplash {
		splash.Jewels(cfg.OutputDir, d2files)
	}

	if cfg.IncreaseStackSizes {
		stacksizes.Increase(d2files)
	}

	if cfg.IncreaseMonsterDensity > 0 {
		monsterdensity.Scale(d2files, cfg.IncreaseMonsterDensity)
	}

	if cfg.EnableTownSkills {
		townskills.Enable(d2files)
	}

	if cfg.NoDropZero {
		treasure.SetNoDropZero(d2files)
	}

	if cfg.QuestDrops {
		treasure.EnableQuestDrops(d2files)
	}

	if cfg.Cowzzz {
		cows.AddTpRecipe(d2files)
		cows.AllowKingKill(d2files)
	}

	// Calculate scores  before any alterations to items.
	var s *scorer.Scorer
	if cfg.RandomOptions.UsePropScores {
		s = scorer.Run(&d2files, cfg.RandomOptions)
	}

	if cfg.UniqueItemDropRate > 0 {
		treasure.ScaleUniqueDropRate(d2files, cfg.UniqueItemDropRate)
	}

	if cfg.RuneDropRate > 0 {
		treasure.ScaleRuneDropRate(d2files, cfg.RuneDropRate)
	}

	if cfg.StartWithCube {
		qol.StartWithCube(d2files)
	}

	if cfg.RemoveUniqCharmLimit {
		qol.RemoveUniqCharmLimit(d2files)
	}

	if cfg.RandomOptions.Randomize {
		if cfg.RandomOptions.UsePropScores {
			g := generator.NewGenerator(&d2files, cfg.RandomOptions, s.TypeTree, s.PSI, s.Statistics)
			g.Run()
		} else {
			randomizer.Run(&cfg, &d2files)
		}
	}

	if cfg.RemoveLevelRequirements {
		reqs.RemoveLevelRequirements(d2files)
	}

	if cfg.RemoveAttRequirements {
		reqs.RemoveAttRequirements(d2files)
	}
	if cfg.RandomOptions.UseOSkills {
		oskills.ConvertSkillsToOSkills(&d2files, cfg)
	}

	d2files.Write()
	writeSeed(cfg)
	util.PP(cfg)
	fmt.Println("===========================")
	fmt.Println("Done!")
	if cfg.EnterToExit {
		fmt.Println("\n[Press enter to exit]")
		fmt.Scanln() // wait for Enter Key
	}
}

func writeSeed(cfg config.Data) {
	filePath := cfg.OutputDir + "Seed.txt"
	f, err := os.Create(filePath)
	util.Check(err)
	defer f.Close()
	f.WriteString(fmt.Sprintf("%d\n", cfg.RandomOptions.Seed))
}
