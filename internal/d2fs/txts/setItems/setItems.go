package setItems

import (
	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/filenumbers"
)

// File Constants
const (
	FileName    = "SetItems.txt"
	NumColumns  = 94
	MaxNumProps = 19
)

// FI File Info for setItems
var FI = d2fs.FileInfo{
	FileName:   FileName,
	FileNumber: filenumbers.SetItems,
	NumColumns: NumColumns,
}

// IFI  ItemFileInfo for setItems
var IFI = d2fs.ItemFileInfo{
	FI:               FI,
	ItemName:         Index,
	Code:             Item,
	Lvl:              LvlReq, // Use LvlReq, not Lvl
	FirstProp:        Prop1,
	NumProps:         MaxNumProps,
	HasEnabledColumn: false,
}

// Header Indexes
const (
	Index        = 0
	Set          = 1
	Item         = 2 // This is type or code?, column header is "item"
	StarItem     = 3 // This is the types name, column header is "*item"
	Rarity       = 4
	Lvl          = 5
	LvlReq       = 6
	ChrTransform = 7
	InvTransform = 8
	InvFile      = 9
	FlippyFile   = 10
	DropSound    = 11
	DropsFxFrame = 12
	UseSound     = 13
	CostMult     = 14
	CostAdd      = 15
	AddFunc      = 16
	Prop1        = 17
	Par1         = 18
	Min1         = 19
	Max1         = 20
	Prop2        = 21
	Par2         = 22
	Min2         = 23
	Max2         = 24
	Prop3        = 25
	Par3         = 26
	Min3         = 27
	Max3         = 28
	Prop4        = 29
	Par4         = 30
	Min4         = 31
	Max4         = 32
	Prop5        = 33
	Par5         = 34
	Min5         = 35
	Max5         = 36
	Prop6        = 37
	Par6         = 38
	Min6         = 39
	Max6         = 40
	Prop7        = 41
	Par7         = 42
	Min7         = 43
	Max7         = 44
	Prop8        = 45
	Par8         = 46
	Min8         = 47
	Max8         = 48
	Prop9        = 49
	Par9         = 50
	Min9         = 51
	Max9         = 52
	AProp1a      = 53
	APar1a       = 54
	AMin1a       = 55
	AMax1a       = 56
	AProp1b      = 57
	APar1b       = 58
	AMin1b       = 59
	AMax1b       = 60
	AProp2a      = 61
	APar2a       = 62
	AMin2a       = 63
	AMax2a       = 64
	AProp2b      = 65
	APar2b       = 66
	AMin2b       = 67
	AMax2b       = 68
	AProp3a      = 69
	APar3a       = 70
	AMin3a       = 71
	AMax3a       = 72
	AProp3b      = 73
	APar3b       = 74
	AMin3b       = 75
	AMax3b       = 76
	AProp4a      = 77
	APar4a       = 78
	AMin4a       = 79
	AMax4a       = 80
	AProp4b      = 81
	APar4b       = 82
	AMin4b       = 83
	AMax4b       = 84
	AProp5a      = 85
	APar5a       = 86
	AMin5a       = 87
	AMax5a       = 88
	AProp5b      = 89
	APar5b       = 90
	AMin5b       = 91
	AMax5b       = 92
	Eol          = 93
)
