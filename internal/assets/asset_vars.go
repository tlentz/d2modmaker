package assets

import (
	"net/http"

	"github.com/tlentz/d2modmaker/internal/modcfg"
)

const (
	PatchStringDest = "data/local/LNG/ENG/"
	DataGlobalExcel = "data/global/excel/"
)

var (
	DataDir                   = "/113c-data/"
	AssetFS   http.FileSystem = Assets
	DataDirFS http.FileSystem = Assets
)

func SetDataDirFS(cfg modcfg.ModConfig) {
	if cfg.PathToDataDir != "" {
		DataDirFS = http.Dir(cfg.PathToDataDir)
		DataDir = "/data/"
	}
}
