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
	MinLvl       = 7 // prop cannot be applied to items whose Req Level is below this
	LvlScale     = 8
	NoTypeOver   = 9  // Can't override itype/etype.  (Example: replenish on armor)
	Itype1       = 10 // Include Type, looked up from armor,weapons Normcode UltraCode, UberCode, and from ItemTypes
	Itype2       = 11 // If non-blank these columns restrict to just these types
	Itype3       = 12 // MagicPrefix.txt & MagicSuffix.txt have same setup.
	Itype4       = 13
	Itype5       = 14
	Itype6       = 15
	Etype1       = 16 // Looked up same way as itype.
	Etype2       = 17 // If the item matches itype, but is of etype, then prop is not allowed
	Etype3       = 18
	Group        = 19
	SynergyGroup = 20
	SourceItem   = 21 // Example of item that contains this prop (not necessarily with same min/max)
	SourceFile   = 22 // File the SourceItem came from
	Eol          = 23
)
