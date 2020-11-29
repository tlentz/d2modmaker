package propscores

// File Constants
const (
	FileName   = "PropScores.txt"
	Path       = "../propscores/" // Placing propscores in assets/propscores/
	NumColumns = 21
)

// Header Indexes
const (
	Prop         = 0  // Property name from Properties
	Par          = 1  // Prop Parameter
	Min          = 2  // Prop Min (Can be # charges, depends on PropParType)
	Max          = 3  // Prop Max (Can be 0 for example %/lvl)
	PropParType  = 4  // See PropParTypes
	ScoreMin     = 5  // Score for minimum roll of prop
	ScoreMax     = 6  // Score for maximum roll of prop
	MinLvl       = 7  // prop cannot be applied to items whose Req Level is below this
	NoTypeOver   = 8  // Can't override itype/etype.  (Example: replenish on armor)
	Itype1       = 9  // Include Type, looked up from armor,weapons Normcode UltraCode, UberCode, and from ItemTypes
	Itype2       = 10 // If non-blank these columns restrict to just these types
	Itype3       = 11 // MagicPrefix.txt & MagicSuffix.txt have same setup.
	Itype4       = 12
	Itype5       = 13
	Itype6       = 14
	Etype1       = 15 // Looked up same way as itype.
	Etype2       = 16 // If the item matches itype, but is of etype, then prop is not allowed
	Etype3       = 17
	Group        = 18
	SynergyGroup = 19
	SourceItem   = 20 // Example of item that contains this prop (not necessarily with same min/max)
	SourceFile   = 21 // File the SourceItem came from
	Eol          = 22
)
