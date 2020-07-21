package splash

import (
	"github.com/tlentz/d2modmaker/internal/d2file"
	"github.com/tlentz/d2modmaker/internal/util"
)

func SplashJewels(d2files d2file.D2Files) {
	panic("")
}

func mergeSplashFile(fileName string, d2files d2file.D2Files) {
	splashDir := "/splash/"
	_, err := d2file.ReadD2File(fileName, splashDir)
	util.Check(err)

	// modFile, err := d2file.GetOrCreateFile()
}

func mergeRows(f1 d2file.D2File, f2 d2file.D2File) d2file.D2File {
	return d2file.D2File{
		FileName: f1.FileName,
		Headers:  f1.Headers,
		Rows:     append(f1.Rows, f2.Rows...),
	}
}
