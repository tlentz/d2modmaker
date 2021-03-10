package filenumbers

// List of filenumbers for well known files
const (
	UniqueItems = 1
	SetItems    = 2
	Sets        = 3
	Runes       = 4
	ItemScores  = 5
)

// NameToNumberMap Maps lowercase filenames to filenumbers
var NameToNumberMap = map[string]int{
	"uniqueitems.txt": UniqueItems,
	"setitems.txt":    SetItems,
	"sets.txt":        Sets,
	"runes.txt":       Runes,
	"itemscores.txt":  ItemScores,
}

//NumberToNameMap Maps filenumbers to mixed case filenames
var NumberToNameMap = map[int]string{
	UniqueItems: "UniqueItems.txt",
	SetItems:    "SetItems.txt",
	Sets:        "Sets.txt",
	Runes:       "Runes.txt",
	ItemScores:  "ItemScores.txt",
}
