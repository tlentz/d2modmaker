package propscorespartype

// All types from PropScores.txt column PropParType
const (
	R   = 1 // Regular
	Req = 2
	Rp  = 3
	Rt  = 4
	Lvl = 5
	S   = 6
	Scl = 7
	Smm = 8
	Sch = 9
	C   = 10
)

// Types Map from string to the constant
var Types = map[string]int{
	"r":   R,   // _, Min, Max		 	Par is supposed to be empty
	"r=":  Req, // _, Min,Max			Min & Max on generated item to be same value
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
