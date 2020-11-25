package runewordlevels

import (
	"fmt"
	"strconv"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/misc"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/runes"
)

var miscItemLevelsCache map[string]int // OBC: yeah ok it's static variable, thread safety yadda yadda

func GetRunewordLevel(row []string, miscItemLevels map[string]int) int {
	if len(miscItemLevels) == 0 {
		fmt.Printf("runewordlevels was not initialized before calling GetRunewordLevel\n")
		panic(1)
	}
	rl := 0
	for j := 0; j < 6; j++ {
		newrl := miscItemLevels[row[runes.Rune1+j]]
		if newrl > rl {
			rl = newrl
		}

	}
	return rl
}

func GetMiscItemLevels(d2files d2fs.Files) map[string]int {
	if len(miscItemLevelsCache) > 0 {
		return miscItemLevelsCache
	}
	f := d2files.Get(misc.FileName)
	miscItemLevelsCache = make(map[string]int)
	for _, row := range f.Rows {
		if row[misc.Code] == "" {
			continue
		}
		n, err := strconv.Atoi(row[misc.Level])
		if err == nil {
			miscItemLevelsCache[row[misc.Code]] = n
		} else {
			fmt.Printf("%s\n", row)
			if row[misc.Code][0] == 'r' {
				fmt.Printf("GetMiscItemLevels Row:%s\n", row[misc.Code])
				d2fs.DebugDumpFiles(d2files, misc.FileName)

				panic(3)
			}
		}
	}
	return miscItemLevelsCache
}
