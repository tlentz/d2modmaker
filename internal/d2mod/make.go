package d2mod

import (
	"fmt"
	"os"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2mod/config"
	"github.com/tlentz/d2modmaker/internal/d2mod/cows"
	"github.com/tlentz/d2modmaker/internal/d2mod/monsterdensity"
	"github.com/tlentz/d2modmaker/internal/d2mod/qol"
	"github.com/tlentz/d2modmaker/internal/d2mod/randomizer"
	"github.com/tlentz/d2modmaker/internal/d2mod/reqs"
	"github.com/tlentz/d2modmaker/internal/d2mod/splash"
	"github.com/tlentz/d2modmaker/internal/d2mod/stacksizes"
	"github.com/tlentz/d2modmaker/internal/d2mod/townskills"
	"github.com/tlentz/d2modmaker/internal/d2mod/treasure"
	"github.com/tlentz/d2modmaker/internal/util"
)

func MakeFromCfgPath(defaultOutDir string, cfgPath string) {
	cfg := config.Read(cfgPath)
	Make(defaultOutDir, cfg)
}

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

	if cfg.RemoveLevelRequirements {
		reqs.RemoveLevelRequirements(d2files)
	}

	if cfg.RemoveAttRequirements {
		reqs.RemoveAttRequirements(d2files)
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
		randomizer.Run(&cfg, d2files)
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
