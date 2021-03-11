package d2items

import (
	"log"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/filenumbers"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/runes"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/setItems"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/uniqueItems"
	"github.com/tlentz/d2modmaker/internal/d2mod/config"
	"github.com/tlentz/d2modmaker/internal/d2mod/propscores"
	"github.com/tlentz/d2modmaker/internal/d2mod/runewordlevels"
)

// PropGetter Bookkeeping structure for reading Items and Props from Item files
type PropGetter struct {
	D2files     *d2fs.Files
	Opts        config.GeneratorOptions
	IFI         *d2fs.ItemFileInfo
	typeOffsets []int          // Can be empty (Sets.txt), 1 item (UniqueItems.txt & SetItems.txt) or 6 items (Runes.txt)
	rwlevels    map[string]int // Calculated based on MiscItems.txt Lvl column for the highest Rune
	psi         *propscores.Maps
	tt          TypeTree
}

// NewPropGetter Return new PropGetter for a particular file
func NewPropGetter(d2files *d2fs.Files, ifi *d2fs.ItemFileInfo, psi *propscores.Maps, tt TypeTree) *PropGetter {
	pg := PropGetter{
		D2files: d2files,
		IFI:     ifi,
		/*Opts:    opts,*/
		psi: psi,
	}
	switch {
	case pg.IFI.FI.FileNumber == filenumbers.UniqueItems:
		pg.typeOffsets = append(pg.typeOffsets, uniqueItems.Code)
	case pg.IFI.FI.FileNumber == filenumbers.Sets:
	case pg.IFI.FI.FileNumber == filenumbers.SetItems:
		pg.typeOffsets = append(pg.typeOffsets, setItems.Item)
	case pg.IFI.FI.FileNumber == filenumbers.Runes:
		// Runes can go into different item types
		pg.typeOffsets = append(pg.typeOffsets, runes.IType1)
		pg.typeOffsets = append(pg.typeOffsets, runes.IType2)
		pg.typeOffsets = append(pg.typeOffsets, runes.IType3)
		pg.typeOffsets = append(pg.typeOffsets, runes.IType4)
		pg.typeOffsets = append(pg.typeOffsets, runes.IType5)
		pg.typeOffsets = append(pg.typeOffsets, runes.IType6)
	default:
		log.Fatalf("unknown file number in NewPropGetter: %d\n", pg.IFI.FI.FileNumber)
	}
	pg.rwlevels = runewordlevels.GetMiscItemLevels(pg.D2files)
	pg.tt = tt
	return &pg
}

// GetProps Returns all Props and Items in the file the PropGetter is set up to read
func (p *PropGetter) GetProps() (Props, Items) {
	f := p.D2files.Get(p.IFI.FI.FileName)
	props := Props{}
	items := Items{}
	for rowIdx, row := range f.Rows {
		item := NewItem(*p, rowIdx, row)
		if item == nil {
			continue
		}
		if item.Name != "" && len(item.Affixes) > 0 {
			items = append(items, *item)
		}
		for _, aff := range item.Affixes {
			props = append(props, aff.P)
		}
	}
	return props, items
}
