package itemscores

import (
	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/filenumbers"
)

// ItemScores.txt file constants
const (
	FileName    = "ItemScores.txt"
	Path        = "propscores/" // Placing in assets/propscores/
	MaxNumProps = 19
)

// FI File Info for setItems
var FI = d2fs.FileInfo{
	FileName:   FileName,
	FileNumber: filenumbers.ItemScores,
	NumColumns: NumColumns,
}

// Creating an IFI for this file would be deceptive because it does not follow the standard
// 4 columns per Prop patter, it has an extra column for Score

// itemscores.txt columns
const (
	File = iota
	Item
	Lvl
	Pbucket
	VanillaFlag
	ItemScore
	Prop1
	Par1
	Min1
	Max1
	Score1
	SMult1
	Prop2
	Par2
	Min2
	Max2
	Score2
	SMult2
	Prop3
	Par3
	Min3
	Max3
	Score3
	SMult3
	Prop4
	Par4
	Min4
	Max4
	Score4
	SMult4
	Prop5
	Par5
	Min5
	Max5
	Score5
	SMult5
	Prop6
	Par6
	Min6
	Max6
	Score6
	SMult6
	Prop7
	Par7
	Min7
	Max7
	Score7
	SMult7
	Prop8
	Par8
	Min8
	Max8
	Score8
	SMult8
	Prop9
	Par9
	Min9
	Max9
	Score9
	SMult9
	Prop10
	Par10
	Min10
	Max10
	Score10
	SMult10
	Prop11
	Par11
	Min11
	Max11
	Score11
	SMult11
	Prop12
	Par12
	Min12
	Max12
	Score12
	SMult12
	Prop13
	Par13
	Min13
	Max13
	Score13
	SMult13
	Prop14
	Par14
	Min14
	Max14
	Score14
	SMult14
	Prop15
	Par15
	Min15
	Max15
	Score15
	SMult15
	Prop16
	Par16
	Min16
	Max16
	Score16
	SMult16
	Prop17
	Par17
	Min17
	Max17
	Score17
	SMult17
	Prop18
	Par18
	Min18
	Max18
	Score18
	SMult18
	Prop19
	Par19
	Min19
	Max19
	Score19
	SMult19
	NumColumns = iota + 5
)
