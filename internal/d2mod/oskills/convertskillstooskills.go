package oskills

import (
	"fmt"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/runes"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/setItems"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/sets"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/uniqueItems"
	"github.com/tlentz/d2modmaker/internal/d2mod/config"
)

// ConvertSkillsToOSkills will change all of the skills props to oskills for the 4 main item files
// UniqueItems, Sets, SetItems and Runes.
func ConvertSkillsToOSkills(d2files *d2fs.Files, cfg config.Data) {
	convertFile(d2files, cfg, &uniqueItems.IFI)
	convertFile(d2files, cfg, &setItems.IFI)
	convertFile(d2files, cfg, &sets.IFI)
	convertFile(d2files, cfg, &runes.IFI)

}
func convertFile(d2files *d2fs.Files, cfg config.Data, ifi *d2fs.ItemFileInfo) {
	conversionCounter := 0
	file := d2files.Get(ifi.FI.FileName)
	//log.Printf("Converting %s", ifi.FI.FileName)
	for rowIdx := range file.Rows {
		if file.Rows[rowIdx][1] == "" {
			continue
		}
		for colIdx := ifi.FirstProp; colIdx < (ifi.FirstProp + ((ifi.NumProps - 1) * 4)); colIdx += 4 {
			if cfg.RandomOptions.UseOSkills && (file.Rows[rowIdx][colIdx] == "skill") {
				//log.Printf("convertFile: %s", file.Rows[rowIdx][colIdx+1]) // Par is the skill name or number
				file.Rows[rowIdx][colIdx] = "oskill"
				conversionCounter++
			}

		}
	}
	fmt.Printf("%s Skills converted to oskills: %d\n", ifi.FI.FileName, conversionCounter)
}
