package propscorepartype

const (
	R   = 1
	Rp  = 2
	Rt  = 3
	Lvl = 4
	S   = 5
	Scl = 6
	Smm = 7
	Sch = 8
	C   = 9
)

// Above & below must stay in synch
var Types = map[string]int{
	"r":   1, // _, Min, Max		 	Par is supposed to be empty
	"rp":  2, // ?, Min, Max		 	Don't touch Par
	"rt":  3, // Time, Min, Max
	"lvl": 4, // Lvl, _, _				%/Lvl: % or pts per Level
	"s":   5, // ?,?,?					Par, Min, Max must all match
	"scl": 6, // %chance, Min, Max
	"smm": 7, // ?, Min, Max
	"sch": 8, // skill, #charges, Level
	"c":   9, // _, Min, Max			Negative effect on player, ScoreMin & ScoreMax are negative example -200 -50
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
