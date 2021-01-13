package propscores

// File Constants
const (
	FileName   = "PropScores.txt"
	Path       = "../propscores/" // Placing propscores in assets/propscores/
	NumColumns = 25
)

// Header Indexes
const (
	Prop         = 0 // Property name from Properties
	Par          = 1 // Prop Parameter
	Min          = 2 // Prop Min (Can be # charges, depends on PropParType)
	Max          = 3 // Prop Max (Can be 0 for example %/lvl)
	PropParType  = 4 // See PropParTypes
	ScoreMin     = 5 // Score for minimum roll of prop
	ScoreMax     = 6 // Score for maximum roll of prop
	ScoreLim     = 7
	MinLvl       = 8 // prop cannot be applied to items whose Req Level is below this
	MaxLvl       = 9
	LvlScale     = 10
	NoTypeOver   = 11 // Can't override itype/etype.  (Example: replenish on armor)
	Itype1       = 12 // Include Type, looked up from armor,weapons Normcode UltaCode, UberCode, and from ItemTypes
	Itype2       = 13 // If non-blank these columns restrict to just these ypes
	Itype3       = 14 // MagicPrefix.txt & MagicSuffix.txt have same setup.
	Itype4       = 15
	Itype5       = 16
	Itype6       = 17
	Etype1       = 18 // Looked up same way as itype.
	Etype2       = 19 // If the item matches itype, but is of etype, then prop is not allowed
	Etype3       = 20
	Group        = 21
	SynergyGroup = 22
	SourceItem   = 23 // Example of item that containsthis prop (not necessarily with same min/max)
	SourceFile   = 24 // File the SourceItem came from
	Eol          = 25 // End of line, should be 0 (to match blizz's format)
)
