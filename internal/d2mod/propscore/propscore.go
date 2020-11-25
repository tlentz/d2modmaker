package propscore

import (
	//"os"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/tlentz/d2modmaker/internal/d2fs"

	//"github.com/tlentz/d2modmaker/internal/d2fs/assets"
	//"github.com/tlentz/d2modmaker/internal/util"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/armor"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/itemTypes"
	propscores "github.com/tlentz/d2modmaker/internal/d2fs/txts/propScores"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/runes"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/setItems"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/sets"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/uniqueItems"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/weapons"
	"github.com/tlentz/d2modmaker/internal/d2mod/config"
	"github.com/tlentz/d2modmaker/internal/d2mod/d2items"
	"github.com/tlentz/d2modmaker/internal/d2mod/propscore/propscorepartype"
)

const (
	propScoresPath = "../propscores/" // Placing propscores in assets/propscores/
)

type ItemScoreMap map[string]int // Map to quickly get a Score based on its Score.Item.Name

//var scorefile *d2fs.File // file currently being scored

//type propscoremap map[string][]propscores.Line

var ScoreLines propscores.ScoreMap // Maps Prop name to list of Lines from propscores.txt

var parentItemType map[string]string // From Armor, Weapons, itemtype: Maps membership of items in groups for doing itype/etype calcs

type scorer struct {
	d2files    d2fs.Files           //
	opts       config.RandomOptions //
	itemScores ItemScoreMap         //
	scorefile  *d2fs.File           // PropScores.txt
	filename   string               // Current filename being scored
	TypeMap    map[string][]string  // Maps from a type to its parents.  From Weapons, Armor & ItemTypes
}

func newscorer(d2files d2fs.Files, opts config.RandomOptions) *scorer {
	s := scorer{
		d2files:    d2files,
		opts:       opts,
		itemScores: ItemScoreMap{},
		TypeMap:    map[string][]string{},
	}

	return &s
}
func ScoreAll(d2files d2fs.Files, opts config.RandomOptions) {
	s := newscorer(d2files, opts)

	// Read in/process PropScores.txt
	s.scorefile = d2files.Read(path.Join(propScoresPath, propscores.FileName))
	ScoreLines = propscores.ScoreMap{}
	for _, r := range s.scorefile.Rows {
		line := propscores.NewLine(r)
		ScoreLines[line.Prop.Name] = append(ScoreLines[line.Prop.Name], *line)
	}

	// Read in/Process Weapons.txt
	armortxt := d2files.Get(armor.FileName)
	for _, r := range armortxt.Rows {
		addTypeMap(s, r[armor.Normcode], r[armor.Type])
		addTypeMap(s, r[armor.Ubercode], r[armor.Type])
		addTypeMap(s, r[armor.Ultracode], r[armor.Type])
	}

	// Read in/Process Weapons.txt
	weaponstxt := d2files.Get(weapons.FileName)
	for _, r := range weaponstxt.Rows {
		addTypeMap(s, r[weapons.Normcode], r[weapons.Type])
		addTypeMap(s, r[weapons.Ubercode], r[weapons.Type])
		addTypeMap(s, r[weapons.Ultracode], r[weapons.Type])
	}

	// Read in/Process Weapons.txt
	itemTypestxt := d2files.Get(itemTypes.FileName)
	for _, r := range itemTypestxt.Rows {
		// fmt.Printf("%s | %s | %s", r[itemTypes.ItemType], r[itemTypes.Code], r[itemTypes.Equiv1])
		addTypeMap(s, r[itemTypes.Code], r[itemTypes.Equiv1])
		addTypeMap(s, r[itemTypes.Code], r[itemTypes.Equiv2])
	}

	scoreFile(s, uniqueItems.FileName)
	scoreFile(s, setItems.FileName)
	scoreFile(s, runes.FileName)
	scoreFile(s, sets.FileName)
	debugshow(s)
}

// Add child/parent relationship to TypeMap
func addTypeMap(s *scorer, child string, parent string) {
	if (child == "") || (parent == "") {
		return
	}
	if strings.TrimSpace(child) != child {
		fmt.Printf("addTypeMap: Bad child string [%s]", child)
		panic(1)
	}
	if strings.TrimSpace(parent) != parent {
		fmt.Printf("addTypeMap: Bad parent string [%s]", parent)
		panic(1)
	}
	if child == parent {
		return
	}

	plist := s.TypeMap[child]
	for _, p := range plist {
		if p == parent {
			return // New parent already exists for the child, bail
		}
	}
	s.TypeMap[child] = append(plist, parent)
}

func scoreFile(s *scorer, filename string) {
	s.filename = filename
	pgetter := d2items.NewPropGetter(s.d2files, s.opts, filename)
	_, items := d2items.GetProps(pgetter)
	for _, item := range items {
		s.itemScores[item.Name] = scoreItem(s, item)
	}
	return
}

func scoreItem(s *scorer, item d2items.Item) int {
	score := 0
	for _, p := range item.Affixes {
		score += scoreProp(s, item, p)
	}
	if score == 0 {
		panic(1)
	}
	return score
}

// Passing whole row so that Synergy can be calculated
func scoreProp(s *scorer, item d2items.Item, p d2items.Prop) int {
	//fmt.Printf(" %s/%s/%s/%s ", p.name, p.par, p.min, p.max)

	if p.Name[0] == '*' { // Blizz comments out with *
		return 0
	}

	for _, line := range ScoreLines[p.Name] {
		if checkPropScore(s, p, item, line) {
			return calcPropScore(s, item, p, line)

		}
	}
	for _, l := range ScoreLines[p.Name] {
		d2items.PrintProp(l.Prop)
		//fmt.Printf("L%d - %s|%s|%s\n", ln, l.Prop.Name, l.Prop.Par)
	}
	fmt.Printf("<Props\n")
	fmt.Printf("len = %d\n", len(ScoreLines[p.Name]))
	fmt.Printf("Couldn't Score %s<%d>: %s|%s|%s|%s\n", item.Name, item.Lvl, p.Name, p.Par, p.Min, p.Max)

	panic(-1)
}

// checkPropScore: Determin if a given prop matches a given line from PropScores.txt
// Checks (based on PropParType) Par, Min < Avg(Prop & Max) < Max and Itype/Etype restrictions
func checkPropScore(s *scorer, p d2items.Prop, item d2items.Item, line propscores.Line) bool {
	//fmt.Printf("checkPropScore: %s|%s>%s|%s\n", p.Name, p.Par, line.Prop.Name, line.Prop.Par)

	if item.Lvl < line.MinLvl {
		return false
	}
	switch line.PropParType {
	case propscorepartype.R, propscorepartype.Rp, propscorepartype.C:
		avg := (p.PropVal.Min + p.PropVal.Max) / 2

		if (avg < line.Prop.PropVal.Min) || (avg > line.Prop.PropVal.Max) {
			return false
		}

	case propscorepartype.Rt:
		if p.Par != line.Prop.Par {
			return false
		}
		avg := (p.PropVal.Min + p.PropVal.Max) / 2

		if (avg < line.Prop.PropVal.Min) || (avg > line.Prop.PropVal.Max) {
			return false
		}

	case propscorepartype.Lvl:
		// Par = N.  Min & Max are don't care

	case propscorepartype.S:
		if p.Par != line.Prop.Par {
			return false
		}
		if p.Min != line.Prop.Min {
			return false
		}
		if p.Max != line.Prop.Max {
			return false
		}

	case propscorepartype.Scl, propscorepartype.Sch:
		if p.Par != line.Prop.Par {
			return false
		}

	case propscorepartype.Smm:
		if p.Par != line.Prop.Par {
			return false
		}
		avg := (p.PropVal.Min + p.PropVal.Max) / 2
		//fmt.Printf("$%d", avg)
		if (avg < line.Prop.PropVal.Min) || (avg > line.Prop.PropVal.Max) {
			//fmt.Printf("%d %d %d", avg, line.Prop.PropVal.Min, line.Prop.PropVal.Max)
			return false
		}

	default:
		fmt.Printf("PropScores.txt has a line with missing PropParType")
		os.Exit(-1)
		//return false
	}
	return checkIETypes(s, p, item, line)
}

// checkIETypes: Check than any of item.Types is a child of propscores.Line.Itypes and not a child of propscore.Line.Etypes
func checkIETypes(s *scorer, p d2items.Prop, item d2items.Item, line propscores.Line) bool {
	//fmt.Printf("checkEITypes: %s %s <%d>%s <%d>%s\n", item.Name, item.Types, len(line.Itypes), line.Itypes, len(line.Etypes), line.Etypes)
	for _, t := range item.Types {
		for _, etype := range line.Etypes {
			if checkTypeMap(s, t, etype) {
				//fmt.Printf("efail\n")
				return false
			}
		}
	}

	if len(item.Types) > 0 {
		if len(line.Itypes) == 0 {
			return true
		}
		for _, t := range item.Types {
			for _, itype := range line.Itypes {
				if checkTypeMap(s, t, itype) {
					//fmt.Printf("isucceed\n")
					return true
				}
			}
		}
		//fmt.Printf("ifail(%s -> %s checked)\n", item.Types, line.Itypes)
		return false // has item.Types but no match
	}
	//fmt.Printf("checkIETypes passed check\n")
	return true // no item.Types
}

// Returns if there is a path from child->parent in TypeMap
// ***** BEWARE RECURSION ******
func checkTypeMap(s *scorer, child string, parent string) bool {
	if parent != strings.TrimSpace(parent) {
		fmt.Printf("checkTypeMap: bad type string (extra spaces): [%s] [%s]\n", child, parent)
		panic(0)
	}
	//fmt.Printf("checkTypeMap: Checking %s %s\n", child, parent)
	if (child == "") || (parent == "") {
		return false
	}
	if child == parent {
		//fmt.Printf("checkTypeMap: Match:%s %s\n", child, parent)
		return true
	}
	plist := s.TypeMap[child]
	for _, p := range plist {
		if p == parent {
			return true
		}
		if checkTypeMap(s, p, parent) {
			return true
		}
	}
	return false
}
func calcPropScore(s *scorer, item d2items.Item, p d2items.Prop, line propscores.Line) int {
	switch line.PropParType {
	case propscorepartype.R, propscorepartype.Rp, propscorepartype.Rt, propscorepartype.Smm, propscorepartype.C:
		return interpolate(p.PropVal.Min, p.PropVal.Max, line.Prop.PropVal.Min, line.Prop.PropVal.Max, line.ScoreMin, line.ScoreMax)
	case propscorepartype.Lvl: // (pts or %)/lvl prop min & max are empty/ignored
		return interpolate(0, p.PropVal.Par, 0, line.Prop.PropVal.Par, line.ScoreMin, line.ScoreMax)
	case propscorepartype.Scl: // Skill, %chance, Level
		if p.Par != line.Prop.Par {
			fmt.Printf("calcPropScore: SCL par mismatch %s %s\n", item.Name, p.Name)
			panic(1)
		}
		v := int(((float32(p.PropVal.Min)) * .01) * float32(p.PropVal.Max))
		return interpolate(v, v, 0, 100, line.ScoreMin, line.ScoreMax)
	case propscorepartype.Sch: // Skill, #charges, Level
		return interpolate(p.PropVal.Max, p.PropVal.Max, 0, line.Prop.PropVal.Max, line.ScoreMin, line.ScoreMax)
	case propscorepartype.S:
		return line.ScoreMax
	default:
		fmt.Printf("calcPropScore: Can't score type %s:%s:%d\n", item.Name, p.Name, line.PropParType)
		panic(1)

	}
}
func interpolate(pmin int, pmax int, lmin int, lmax int, smin int, smax int) int {
	avg := pmin + pmax
	if avg != 0 {
		avg = avg / 2
	}
	if lmin == lmax {
		return smax
	}
	return ((avg-lmin)/(lmax-lmin))*(smax-smin) + smin
}
func debugshow(s *scorer) {
	//fmt.Printf("-----------------------Debugging------------------------\n")
	/*
		for propname, lines := range ScoreLines {
			fmt.Printf("%s\t%d\n", propname, len(lines))
		}
	*/
	/*
		for _, lines := range ScoreLines {
			for _, l := range lines {
				fmt.Printf("%s %d %d\n", l.Prop.Name, l.Prop.PropVal.Min, l.Prop.PropVal.Max)
			}
		}
	*/
	/*
		fmt.Printf("--------------------------------------------------------\n")
		fmt.Printf("*****\nDumping typemap\n*****\n")
		for ctype, ptype := range s.TypeMap {
			fmt.Printf("%s\t%s\n", ctype, ptype)
		}
	*/
	fmt.Printf("--------------------------------------------------------\n")
	fmt.Printf("Dumping Item Scores\n")
	for key, el := range s.itemScores {
		fmt.Printf("%s\t%d\n", key, el)
	}
	//fmt.Printf("--------------------------------------------------------\n")
}
