package uniqueItemsTxt

// File Constants
const (
	FileName   = "UniqueItems.txt"
	NumColumns = 70
)

// Header Indexes
const (
	Index        = 0
	Version      = 1
	Enabled      = 2
	Ladder       = 3
	Rarity       = 4
	NoLimit      = 5
	Lvl          = 6
	LvlReq       = 7
	Code         = 8
	Type_        = 9
	Uber         = 10
	Carry1       = 11
	CostMult     = 12
	CostAdd      = 13
	ChrTransform = 14
	InvTransform = 15
	FlippyFile   = 16
	InvFile      = 17
	DropSound    = 18
	DropsFxFrame = 19
	UseSound     = 20
	Prop1        = 21
	Par1         = 22
	Min1         = 23
	Max1         = 24
	Prop2        = 25
	Par2         = 26
	Min2         = 27
	Max2         = 28
	Prop3        = 29
	Par3         = 30
	Min3         = 31
	Max3         = 32
	Prop4        = 33
	Par4         = 34
	Min4         = 35
	Max4         = 36
	Prop5        = 37
	Par5         = 38
	Min5         = 39
	Max5         = 40
	Prop6        = 41
	Par6         = 42
	Min6         = 43
	Max6         = 44
	Prop7        = 45
	Par7         = 46
	Min7         = 47
	Max7         = 48
	Prop8        = 49
	Par8         = 50
	Min8         = 51
	Max8         = 52
	Prop9        = 53
	Par9         = 54
	Min9         = 55
	Max9         = 56
	Prop10       = 57
	Par10        = 58
	Min10        = 59
	Max10        = 60
	Prop11       = 61
	Par11        = 62
	Min11        = 63
	Max11        = 64
	Prop12       = 65
	Par12        = 66
	Min12        = 67
	Max12        = 68
	Eol          = 69
)

type UProp struct {
	UPropProp string
	UPropPar  string
	UPropMin  string
	UPropMax  string
}
