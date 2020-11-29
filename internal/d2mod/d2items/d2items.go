package d2items

import (
	"fmt"
	"log"
	"strconv"

	"github.com/tlentz/d2modmaker/internal/d2fs/filenumbers"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/propscores"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/sets"
	"github.com/tlentz/d2modmaker/internal/d2mod/prop"
	"github.com/tlentz/d2modmaker/internal/d2mod/runewordlevels"
)

// Prop convenience alias to prop.Prop
type Prop = prop.Prop

// Props convenience alias to prop.Props
type Props = prop.Props

// Affix Contains Prop + it's position on an item and matching PropScores.txt Line
type Affix struct {
	P                  prop.Prop
	PSet               int
	SetBonusMultiplier float32
	ColIdx             int
	Line               *propscores.Line
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
	Enabled    bool     // false when row[2] == 1 in Runes or UniqueItems, true otherwise
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
	item := Item{}
	item.Name = row[pg.IFI.ItemName]
	item.Lvl = lvl
	item.Enabled = true
	item.Affixes = []Affix{}
	if pg.IFI.HasEnabledColumn {
		// Third Column (row[2]) in Runes and in Unique Items must be 1 for an item to exist
		//fmt.Printf("row[2] %s - %s\n", row[0], row[2])
		if row[2] != "1" {
			//fmt.Printf("NewItem: Not Enabled: %s\n", row[0])
			// fmt.Println(row)
			item.Enabled = false
			//log.Fatalf(row[0] + row[1] + row[2])
			return nil
		}
	}
	// Sets has hard-coded type "fset", UniqueItems & SetItems have 1 type, Runewords have multiple types
	for _, t := range pg.typeOffsets {
		//fmt.Printf("NewItem: Type = %s", row[t])
		item.Types = append(item.Types, row[t])
	}
	if pg.IFI.FI.FileName == sets.FileName {
		item.Types = append(item.Types, "fset")
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
		// FIXME:TODO: Need to re-implement perfect props AFTER scoring.
		// Don't forget to add exclusion for skill-rand
		/*
			if pg.Opts.PerfectProps {
				prop.Min = prop.Max
			}
			if pg.Opts.UseOSkills {
				if prop.Name == "skill" {
					prop.Name = "oskill"
				}
			}
		*/
		if aff.Line == nil {
			panic(1)
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
			log.Fatal("Item.ToRow: colIdx == 0")
		}
		newrow[colIdx] = aff.P.Name
		newrow[colIdx+1] = aff.P.Par
		newrow[colIdx+2] = aff.P.Min
		newrow[colIdx+3] = aff.P.Max

	}
	return newrow
}

// NewAffixFromRow Create an Affix from an Item Row
func NewAffixFromRow(pg PropGetter, item Item, row []string, colIdx int) *Affix {
	aff := Affix{
		P: prop.NewProp(row[colIdx], row[colIdx+1], row[colIdx+2], row[colIdx+3]),
	}
	aff.SetBonusMultiplier = CalcSetBonusMultiplier(pg.IFI.FI.FileNumber, colIdx)
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
		log.Fatalf("NewAffix: Couldn't find line in PropScores.txt for %s[%d] %s|%s|%s|%s", item.Name, item.Lvl, aff.P.Name, aff.P.Par, aff.P.Min, aff.P.Max)
	}
	return &aff
}

// NewAffixFromLine Used by Generator, create an Affix from a PropScores.txt Line
func NewAffixFromLine(line *propscores.Line, colIdx int, setBonusMultiplier float32) *Affix {
	aff := Affix{
		P: line.Prop,
	}
	aff.Line = line
	aff.ColIdx = colIdx
	aff.SetBonusMultiplier = setBonusMultiplier
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
