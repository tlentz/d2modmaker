package propscorespartype

// All types from PropScores.txt column PropParType
const (
	R   = 1 // Regular
	Rp  = 2
	Rt  = 3
	Lvl = 4
	S   = 5
	Scl = 6
	Smm = 7
	Sch = 8
	C   = 9
)

// Types Map from string to the constant
var Types = map[string]int{
	"r":   R,   // _, Min, Max		 	Par is supposed to be empty
	"rp":  Rp,  // ?, Min, Max		 	Don't touch Par
	"rt":  Rt,  // Time, Min, Max
	"lvl": Lvl, // Lvl, _, _				%/Lvl: % or pts per Level
	"s":   S,   // ?,?,?					Par, Min, Max must all match
	"scl": Scl, // %chance, Min, Max
	"smm": Smm, // ?, Min, Max
	"sch": Sch, // skill, #charges, Level
	"c":   C,   // _, Min, Max			Negative effect on player, ScoreMin & ScoreMax are negative example -200 -50
}

/* type PropScoreParTypes map[string]int

func NewPropScoreParTypes() *PropScoreParTypes {
	p := make(PropScoreParTypes)
	p["r"] = r
	p["rp"] = rp
	p["rt"] = rt
	p["lvl"] = lvl
	p["s"] = s
	p["scl"] = scl
	p["smm"] = smm
	p["sch"] = sch
	p["c"] = c
	return &p
}
*/
