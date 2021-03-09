package d2items

import (
	"fmt"
	"log"
	"strconv"

	"github.com/tlentz/d2modmaker/internal/d2fs/filenumbers"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/ignoretxt"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/runes"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/sets"
	"github.com/tlentz/d2modmaker/internal/d2mod/ignore"
	"github.com/tlentz/d2modmaker/internal/d2mod/prop"
	"github.com/tlentz/d2modmaker/internal/d2mod/propscores"
	"github.com/tlentz/d2modmaker/internal/d2mod/runewordlevels"
)

// Prop convenience alias to prop.Prop
type Prop = prop.Prop

// Props convenience alias to prop.Props
type Props = prop.Props

// Affix Contains Prop + it's position on an item and matching PropScores.txt Line
type Affix struct {
	P        prop.Prop
	PSet     int
	ColIdx   int
	Line     *propscores.Line
	RawScore int
	//ScoreMult          float32
	SetBonusMultiplier float32
	SynergyMultiplier  float32
	AdjustedScore      int
	Ignored            bool // This affix exists in Ignore.txt.  It therefore won't have a valid Line
}

// ItemAffixes list of Affixes (containing props) on an Item
type ItemAffixes = []Affix

// Item Armor, Weapon, Runeword, etc coming from or to be written to one of the 4 Item files (UniqueItems.txt, Sets.txt, etc)
type Item struct {
	Name       string
	Lvl        int
	Affixes    ItemAffixes
	FileNumber int      // filenumber this item came from/going to
	RowIdx     int      // Row# this item came from/going to
	Types      []string // Similar to typeOffsets, can be empty, 1 or up to 6.
	Code       string   // Item Code
	Enabled    bool     // false when row[2] == 1 in Runes or UniqueItems, true otherwise
	Score      int
}

// Items is a slice of Item
type Items = []Item

// CloneWithoutAffixes Clone an Item except the Affixes
func (item *Item) CloneWithoutAffixes() *Item {
	var n Item = *item
	return &n
}

// NewItem Instantiates a new item from an Item file
func NewItem(pg PropGetter, rowIdx int, row []string) *Item {
	//fmt.Printf("NewItem %s\n", row[0])
	lvl := 0
	if pg.IFI.Lvl >= 0 {
		lvl, _ = strconv.Atoi(row[pg.IFI.Lvl])
	} else if pg.IFI.FI.FileNumber == filenumbers.Runes { //Runes.txt doesn't have a level column
		lvl = runewordlevels.GetRunewordLevel(row, pg.rwlevels)
		//lvl = -1 // Temporary for testing
	}
	if pg.IFI.HasEnabledColumn {
		// Third Column (row[2]) in Runes and in Unique Items must be 1 for an item to exist
		//fmt.Printf("row[2] %s - %s\n", row[0], row[2])
		if row[2] != "1" {
			//fmt.Printf("NewItem: Not Enabled: %s\n", row[0])
			// fmt.Println(row)
			//log.Fatalf(row[0] + row[1] + row[2])
			return nil
		}
	}
	if ignore.IsIgnored(pg.IFI.FI.FileNumber, ignoretxt.IgnoreTypeItem, row[0]) {
		// Item is ignored
		return nil
	}

	item := Item{}
	item.Name = row[pg.IFI.ItemName]
	item.Lvl = lvl
	item.Enabled = true
	item.Affixes = []Affix{}
	// Sets has hard-coded type "fset", UniqueItems & SetItems have 1 type, Runewords have multiple types
	for _, t := range pg.typeOffsets {
		//fmt.Printf("NewItem: Type = %s", row[t])
		item.Types = append(item.Types, row[t])
	}
	if pg.IFI.FI.FileName == sets.FileName {
		item.Types = append(item.Types, "fset")
	}
	if pg.IFI.FI.FileName == runes.FileName {
		item.Types = append(item.Types, "rune")
	}
	switch {
	case pg.IFI.FI.FileNumber == filenumbers.Runes:
		// Fudging it by using item Type instead of Code.  This means codes must be loaded into pBuckets too.
		item.Code = row[pg.typeOffsets[0]]
	case pg.IFI.FI.FileNumber == filenumbers.Sets:
		item.Code = "fset"
	default:
		item.Code = row[pg.IFI.Code]
	}
	item.FileNumber = pg.IFI.FI.FileNumber
	item.RowIdx = rowIdx

	// FIXME: don't use len(row), use pg.MaxNumProps * 4
	for i := pg.IFI.FirstProp; i < len(row)-3; i += 4 {
		if row[i] == "" {
			continue
		}
		if row[i][0] == '*' {
			continue
		}
		//prop := prop.NewProp(row[i], row[i+1], row[i+2], row[i+3])
		aff := NewAffixFromRow(pg, item, row, i)
		if aff.Line == nil && !(aff.Ignored) {
			log.Panicf("Item %s column#%d no Affix created? (contact developer)\n", row[0], i)
		}
		//item.SetBonus = append(item.SetBonus, sbn)
		item.Affixes = append(item.Affixes, *aff)
	}
	if len(item.Types) == 0 {
		log.Fatalf("New Item: No Types")
	}
	return &item
}

// ToRow creates a row suitable for writing to an Item file from an Item
func (item *Item) ToRow(pg PropGetter, row []string) []string {

	// First copy everything but the props
	newrow := make([]string, len(row))

	// Erase existing props
	for colIdx, c := range row {
		if (colIdx >= pg.IFI.FirstProp) && (colIdx < (pg.IFI.FirstProp + (pg.IFI.NumProps * 4))) {
			newrow[colIdx] = ""
		} else {
			newrow[colIdx] = c
		}
	}
	if len(item.Affixes) > pg.IFI.NumProps {
		log.Fatalf("Item.ToRow: Too many affixes generated for item %s(%d)", row[0], len(item.Affixes))
	}
	for _, aff := range item.Affixes {
		if item.FileNumber != pg.IFI.FI.FileNumber {
			// Reason for this assertion is that column #s generated for one file don't match another file.
			log.Panic("Mismatched FileNumber")
		}
		colIdx := aff.ColIdx
		if colIdx == 0 {
			log.Panicf("Item.ToRow: colIdx == 0: %s|%s|%s|%s", aff.P.Name, aff.P.Par, aff.P.Min, aff.P.Max)
		}
		newrow[colIdx] = aff.P.Name
		newrow[colIdx+1] = aff.P.Par
		newrow[colIdx+2] = aff.P.Min
		newrow[colIdx+3] = aff.P.Max
		// if pg.Opts.UseOSkills && (aff.P.Name == "skill") {
		// 	newrow[colIdx] = "oskill"
		// }

	}
	return newrow
}

// NewAffixFromRow Create an Affix from an Item Row
func NewAffixFromRow(pg PropGetter, item Item, row []string, colIdx int) *Affix {
	lvl := pg.IFI.Lvl
	if pg.IFI.Lvl == -1 {
		lvl = runewordlevels.GetRunewordLevel(row, pg.rwlevels)
	}
	aff := Affix{
		P: prop.NewProp(row[colIdx], row[colIdx+1], row[colIdx+2], row[colIdx+3], lvl),
	}
	if ignore.IsIgnored(item.FileNumber, ignoretxt.IgnoreTypeProp, aff.P.Name) {
		// It's in ignore.txt
		aff.Ignored = true
		aff.SetBonusMultiplier = 1
		aff.SynergyMultiplier = 1
		return &aff
	}
	aff.SetBonusMultiplier = CalcSetBonusMultiplier(pg.IFI.FI.FileNumber, colIdx)
	aff.ColIdx = colIdx
	for _, line := range pg.psi.PropLines[aff.P.Name] {
		//line := pg.psi.PropLines[aff.P.Name][idx]
		if checkPropScore(&pg.tt, aff.P, item, line) {
			//s.Group[p] = line.Group
			//s.SynergyGroup[p] = line.SynergyGroup
			//return calcPropScore(p, line), &line
			aff.Line = line
			return &aff
		}
	}

	if (item.Lvl > 0) && ((pg.IFI.HasEnabledColumn && row[2] == "1") || (!pg.IFI.HasEnabledColumn)) {
		log.Printf("Item Type: %v\n", item.Types)
		for _, line := range pg.psi.PropLines[aff.P.Name] {
			//tmatch := "  "
			if CheckIETypes(&pg.tt, item.Types, line.Itypes, line.Etypes) {
				//tmatch = "->"
			}
			//log.Printf("%s %s %v|%v\n", line.Prop.Name, tmatch, line.Itypes, line.Etypes)
		}

		// for x, y := range pg.tt.parentItemType {
		// 	fmt.Printf("%s", x)
		// 	for _, foo := range y {
		// 		fmt.Printf("\t%s", foo)
		// 	}
		// 	fmt.Println()
		// }

		log.Fatalf("NewAffixFromRow: Couldn't find line in PropScores.txt for %s[%d] %s|%s|%s|%s", item.Name, item.Lvl, aff.P.Name, aff.P.Par, aff.P.Min, aff.P.Max)
	}
	return &aff
}

// NewAffixFromLine Used by Generator, create an Affix from a PropScores.txt Line
func NewAffixFromLine(line *propscores.Line, colIdx int, filenumber int) *Affix {
	aff := Affix{
		P:                  line.Prop,
		Line:               line,
		ColIdx:             colIdx,
		SetBonusMultiplier: CalcSetBonusMultiplier(filenumber, colIdx),
	}
	// aff.SetBonusMultiplier = CalcSetBonusMultiplier(pg.IFI.FI.FileNumber, colIdx)
	return &aff
}

// AppendProps append all props in p2 to p1.  For collisions p2 wins.
func AppendProps(p1 Props, p2 Props) {
	for key, val := range p2 {
		p1[key] = val
	}
}

// PrintProp Debuggin Prop printer
func PrintProp(p Prop) {
	fmt.Printf("Prop>%s|%s|%s|%s %d|%d|%d\n", p.Name, p.Par, p.Min, p.Max, p.Val.Par, p.Val.Min, p.Val.Max)
}
