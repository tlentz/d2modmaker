package d2items

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/runes"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/setItems"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/sets"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/uniqueItems"
	"github.com/tlentz/d2modmaker/internal/d2mod/config"
	"github.com/tlentz/d2modmaker/internal/d2mod/runewordlevels"
)

type PropVal struct {
	Par int
	Min int
	Max int
}

type Prop struct {
	Name    string
	Par     string
	Min     string
	Max     string
	Lvl     int
	PropVal PropVal
}

// Props is a slice of Prop
type Props = []Prop

// Item represents an item
type Item struct {
	Name    string
	Lvl     int
	Affixes Props
	Types   []string // Like typeOffsets, can be empty, 1 or up to 6.
}

// Items is a slice of Item
type Items = []Item

type PropGetter struct {
	D2files d2fs.Files
	Opts    config.RandomOptions
	//Props       Props
	FileName    string
	propOffset  int
	maxNumProps int
	levelOffset int
	nameOffset  int
	typeOffsets []int // Can be empty (Sets.txt), 1 item (UniqueItems.txt & SetItems.txt) or 6 items (Runes.txt)
}

func GetProps(p PropGetter) (Props, Items) {
	rwlevels := runewordlevels.GetMiscItemLevels(p.D2files)
	f := p.D2files.Get(p.FileName)
	props := Props{}
	items := Items{}
	for _, row := range f.Rows {
		lvl := 0
		if p.levelOffset >= 0 {
			lvl, _ = strconv.Atoi(row[p.levelOffset])
		} else if p.FileName == runes.FileName { //Runes.txt doesn't have a level column
			lvl = runewordlevels.GetRunewordLevel(row, rwlevels)
			//lvl = -1 // Temporary for testing
		}

		item := Item{}
		item.Name = row[p.nameOffset]
		item.Lvl = lvl
		if (p.FileName == uniqueItems.FileName) || (p.FileName == runes.FileName) {
			// Third Column (row[2]) in Runes and in Set Items must be 1 for an item to exist
			//fmt.Printf("row[2] %s - %s\n", row[0], row[2])
			if strings.Compare(row[2], "1") != 0 {
				continue
			}
		}
		for i := p.propOffset; i < len(row)-3; i += 4 {
			/*
				prop := Prop{
					Name: row[i],
					Par:  row[i+1],
					Min:  row[i+2],
					Max:  row[i+3],
					Lvl:  lvl,
				}
				prop.PropVal.Par, _ = strconv.Atoi(prop.Par)
				prop.PropVal.Min, _ = strconv.Atoi(prop.Min)
				prop.PropVal.Max, _ = strconv.Atoi(prop.Max)
			*/
			prop := NewProp(row[i], row[i+1], row[i+2], row[i+3])
			if prop.Name == "" {
				continue
			}
			// FIXME:TODO: Need to re-implement perfect props AFTER scoring.
			// Don't forget to add exclusion for skill-rand (maybe use PropScores.txt:PropParType?)
			/*
				if p.Opts.PerfectProps {
					prop.Min = prop.Max
				}
				if p.Opts.UseOSkills {
					if prop.Name == "skill" {
						prop.Name = "oskill"
					}
				}
			*/
			props = append(props, prop)
			item.Affixes = append(item.Affixes, prop)
		}
		// Sets has hard-coded type "fset", UniqueItems & SetItems have 1 type, Runewords have multiple types
		for _, t := range p.typeOffsets {
			//fmt.Printf("GetProps: Type = %s", row[t])
			item.Types = append(item.Types, row[t])
		}
		if p.FileName == sets.FileName {
			item.Types = append(item.Types, "fset")
		}
		if item.Name != "" && len(item.Affixes) > 0 {
			items = append(items, item)
		}
	}
	return props, items
}

// Initialize propOffset, numProps, levelOffset, nameOffset
func NewPropGetter(d2files d2fs.Files, opts config.RandomOptions, filename string) PropGetter {
	pg := PropGetter{}
	pg.D2files = d2files
	pg.FileName = filename
	pg.Opts = opts
	switch {
	case pg.FileName == uniqueItems.FileName:
		pg.propOffset = uniqueItems.Prop1
		pg.maxNumProps = uniqueItems.MaxNumProps
		pg.levelOffset = uniqueItems.Lvl
		pg.nameOffset = uniqueItems.Index
		pg.typeOffsets = append(pg.typeOffsets, uniqueItems.Code)
	case pg.FileName == sets.FileName:
		pg.propOffset = sets.PCode2a
		pg.maxNumProps = sets.MaxNumProps
		pg.levelOffset = sets.Level
		pg.nameOffset = sets.Index
	case strings.Compare(pg.FileName, "SetItems.txt") == 0:
		pg.propOffset = setItems.Prop1
		pg.maxNumProps = setItems.MaxNumProps
		pg.levelOffset = setItems.LvlReq
		pg.nameOffset = setItems.Index
		pg.typeOffsets = append(pg.typeOffsets, setItems.Item)
	case strings.Compare(pg.FileName, "Runes.txt") == 0:
		pg.propOffset = runes.T1Code1
		pg.maxNumProps = runes.MaxNumProps
		pg.levelOffset = -1
		pg.nameOffset = runes.RuneName
		// Runes can go into different item types
		pg.typeOffsets = append(pg.typeOffsets, runes.IType1)
		pg.typeOffsets = append(pg.typeOffsets, runes.IType2)
		pg.typeOffsets = append(pg.typeOffsets, runes.IType3)
		pg.typeOffsets = append(pg.typeOffsets, runes.IType4)
		pg.typeOffsets = append(pg.typeOffsets, runes.IType5)
		pg.typeOffsets = append(pg.typeOffsets, runes.IType6)
	}
	return pg
}

func (p *Prop) GetId() string {
	if p.Name == "aura" {
		// Two auras do not work even if they are different types
		return p.Name
	} else {
		// Otherwise include both the prop type and sub-type
		return p.Name + p.Par
	}
}
func NewProp(name string, par string, min string, max string) Prop {
	/*
		if name == "sock" {
			fmt.Printf("NewProp: [%s][%s][%s][%s]", name, par, min, max)
		}
	*/
	prop := Prop{
		Name: name,
		Par:  par,
		Min:  min,
		Max:  max,
	}
	prop.PropVal.Par, _ = strconv.Atoi(prop.Par)
	prop.PropVal.Min, _ = strconv.Atoi(prop.Min)
	prop.PropVal.Max, _ = strconv.Atoi(prop.Max)
	return prop
}
func PrintProp(p Prop) {
	fmt.Printf("%s|%s|%s|%s %d|%d|%d\n", p.Name, p.Par, p.Min, p.Max, p.PropVal.Par, p.PropVal.Min, p.PropVal.Max)
}
