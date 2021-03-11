package modsupport

import (
	"fmt"
	"log"
	"path"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/assets"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/propscorestxt"
	"github.com/tlentz/d2modmaker/internal/d2mod/config"
	"github.com/tlentz/d2modmaker/internal/d2mod/ignore"
)

// ModSupport constants
const (
	DefaultModName = "113c"
)

// Run Activate general mod support features, based on ModName chosen
// May disable options in the config based on cfg_override.json
func Run(cfg *config.Data, d2files d2fs.Files) {
	fmt.Println("Mod Support Start")
	sourceDir := cfg.SourceDir
	if cfg.ModName == "" {
		log.Panicf("Module name cannot be blank.  Use 113c for no mod.")
	}
	if cfg.ModName == DefaultModName {
		sourceDir = assets.AssetDir + assets.DataDir
	}
	if cfg.ModName != DefaultModName && sourceDir == "" {
		log.Fatalln("Modsupport.Run: Error: Must specify Source Directory if not " + DefaultModName)
	}
	dataFiles := d2files.List(path.Join(sourceDir, assets.GlobalExcelDir))
	dataFilesMap := make(map[string]int, 0)
	for idx := range dataFiles {
		if !dataFiles[idx].IsDir() {
			dataFilesMap[dataFiles[idx].Name()] = idx + 1 // Default value is 0 and not very easy to change, it's easier just to -1
		}
	}
	modsupportDir := path.Join(assets.ModSupportDir, cfg.ModName) // This path is relative to assets.AssetDir
	modFiles := d2files.List(path.Join(assets.AssetDir, modsupportDir))
	for idx := range modFiles {
		fmt.Printf("    Loading Filename %s: ", modFiles[idx].Name())
		switch {
		case modFiles[idx].IsDir():
			// Ignore directories
		case modFiles[idx].Name() == "PropScores.txt",
			modFiles[idx].Name() == "PBucketList.txt":
			d2fs.MergeRows(d2files.GetAsset(propscorestxt.Path, modFiles[idx].Name()),
				*d2fs.ReadAsset(modsupportDir, modFiles[idx].Name()))
			fmt.Println("Merged")
		case modFiles[idx].Name() == "ignore.txt":
			d2files.GetAsset(modsupportDir, modFiles[idx].Name())
			fmt.Println("loaded.")
		case dataFilesMap[modFiles[idx].Name()] > 0:
			d2fs.MergeRows(d2files.Get(modFiles[idx].Name()), *d2fs.ReadAsset(modsupportDir, modFiles[idx].Name()))
		case modFiles[idx].Name() == "cfg_override.json":
			fmt.Println("not yet supported")
		default:
			fmt.Printf("Skipping (unknown file)\n")
		}

	}
	ignore.Init(d2files)

	// d2fs.MergeRows(d2files.Get(itemStatCost.FileName), *d2fs.ReadAsset(elementalAssetsDir, itemStatCost.FileName))
	fmt.Println("Mod Support Done")
}
