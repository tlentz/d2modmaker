package sets

import (
	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/filenumbers"
)

// File Constants
const (
	FileName    = "Sets.txt"
	NumColumns  = 69
	MaxNumProps = 16
)

var FI = d2fs.FileInfo{
	FileName:   FileName,
	FileNumber: filenumbers.Sets,
	NumColumns: NumColumns,
}
var IFI = d2fs.ItemFileInfo{
	FI:               FI,
	ItemName:         Index,
	Lvl:              Level,
	FirstProp:        PCode2a,
	NumProps:         MaxNumProps,
	HasEnabledColumn: false,
}

// Header Indexes
const (
	Index    = 0
	Name     = 1
	Version  = 2
	Level    = 3
	PCode2a  = 4
	PParam2a = 5
	PMin2a   = 6
	PMax2a   = 7
	PCode2b  = 8
	PParam2b = 9
	PMin2b   = 10
	PMax2b   = 11
	PCode3a  = 12
	PParam3a = 13
	PMin3a   = 14
	PMax3a   = 15
	PCode3b  = 16
	PParam3b = 17
	PMin3b   = 18
	PMax3b   = 19
	PCode4a  = 20
	PParam4a = 21
	PMin4a   = 22
	PMax4a   = 23
	PCode4b  = 24
	PParam4b = 25
	PMin4b   = 26
	PMax4b   = 27
	PCode5a  = 28
	PParam5a = 29
	PMin5a   = 30
	PMax5a   = 31
	PCode5b  = 32
	PParam5b = 33
	PMin5b   = 34
	PMax5b   = 35
	FCode1   = 36
	FParam1  = 37
	FMin1    = 38
	FMax1    = 39
	FCode2   = 40
	FParam2  = 41
	FMin2    = 42
	FMax2    = 43
	FCode3   = 44
	FParam3  = 45
	FMin3    = 46
	FMax3    = 47
	FCode4   = 48
	FParam4  = 49
	FMin4    = 50
	FMax4    = 51
	FCode5   = 52
	FParam5  = 53
	FMin5    = 54
	FMax5    = 55
	FCode6   = 56
	FParam6  = 57
	FMin6    = 58
	FMax6    = 59
	FCode7   = 60
	FParam7  = 61
	FMin7    = 62
	FMax7    = 63
	FCode8   = 64
	FParam8  = 65
	FMin8    = 66
	FMax8    = 67
	Eol      = 68
)
