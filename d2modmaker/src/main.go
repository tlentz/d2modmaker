package main

const dataDir = "../assets/113c-data/"
const outDir = "../dist/"

func main() {
	var cfg = ReadCfg()
	var d2files = map[string]D2File{}

	PP(cfg)

	if cfg.IncreaseStackSizes {
		d2file := GetOrCreateFile(&d2files, "Misc.txt")
		IncreaseStackSizes(d2file)
	}

	if cfg.IncreaseMonsterDensity > 1 {
		d2file := GetOrCreateFile(&d2files, "Levels.txt")
		IncreaseMonsterDensity(d2file, cfg.IncreaseMonsterDensity)
	}

	if cfg.LinearRuneDrops {
		GetOrCreateFile(&d2files, "TreasureClassEx.txt")
	}

	if cfg.EnableTownSkills {
		d2file := GetOrCreateFile(&d2files, "Skills.txt")
		EnableTownSkills(d2file)
	}

	WriteFiles(&d2files)
}
