package propscores

// File Constants
const (
	FileName   = "PropScores.txt"
	Path       = "../propscores/" // Placing propscores in assets/propscores/
	NumColumns = 21
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
	ScoreLim 	 = 7
	MinLvl       = 8// prop cannot be applied to items whose Req Level is below this
	LvlScale     = 9
	NoTypeOver   = 10  // Can't override itype/etype.  (Example: replenish on armor)
	Itype1       = 11 // Include Type, looked up from armor,weapons Normcode UltaCode, UberCode, and from ItemTypes
	Itype2       = 12 // If non-blank these columns restrict to just these ypes
	Itype3       = 13// MagicPrefix.txt & MagicSuffix.txt have same setup.
	Itype4       = 14
	Itype5       = 15
	Itype6       = 16
	Etype1       = 17 // Looked up same way as itype.
	Etype2       = 18// If the item matches itype, but is of etype, then prop is not allowed
	Etype3       = 19
	Group        = 20
	SynergyGroup = 21
	SourceItem   = 22 // Example of item that containsthis prop (not necessarily with same min/max)
	SourceFile   = 23// File the SourceItem came from
	Eol          = 24
)
