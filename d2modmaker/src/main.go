package main

const dataDir = "../../assets/113c-data/"
const outDir = "../dist/"

func main() {
	var cfg = ReadCfg()
	var d2files = map[string]D2File{}

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

	WriteFiles(&d2files)
}
