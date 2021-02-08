package propscores

import (
	"log"
	"strconv"
	"strings"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/propscorestxt"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/propscorestxt/propscorespartype"
	"github.com/tlentz/d2modmaker/internal/d2mod/prop"
)

// Line Contains 1 line from PropScore.txt
type Line struct {
	RowIndex     int
	Prop         prop.Prop
	PropParType  int      // See constants in propscorepartype
	ScoreMin     int      // Score for Prop.Min or 0, type dependent
	ScoreMax     int      // Score for Prop.Max, Prop.Par, or Prop.Min * Prop.Max:  type dependent
	ScoreLimit   int      // Specifies a % limit which generated item cannot exceed for this prop
	MinLvl       int      // Minimum level of item this prop can appear on
	MaxLvl       int      // Maximum level of item this prop can appear on
	LvlScale     bool     // scale this item for 1/1 at level 50
	NoTypeOvr    bool     // Don't allow type overriding for this prop
	Itypes       []string // Include Types  If non-empty item must be one of these types or a child
	Etypes       []string // Exclude Types  If non-empty item must not be one of these types or a child
	Group        string   // Cannot roll 2 props from same group on same item.
	SynergyGroup string   // This prop synergizes with other props in this SynergyGroup
	SourceItem   string   // 1 example of this prop
	SourceFile   string   // File for SourceItem
}

// ScoreMap map of Prop Name's to  PropScores.txt Lines
// Beware that this list is only prop name, not prop&par, so there will be a lot of duplicates for things like "skill"
// The list must be parsed with level, etype, itype restrictions etc. to find the right Line.
type ScoreMap map[string][]*Line // Used to grab all the Lines materialized from PropScores.txt for a given prop name

// Maps Maps to help with lookups into PropScores.txt
type Maps struct {
	PropLines ScoreMap // Map from Prop name to  array of PropScore.txt []Lines
	RowLines  []*Line  // Array of Lines by RowIdx i.e. matching d2files.Rows[]
}

// NewLine Create a new Line from a given row in PropScores.txt
func NewLine(Row []string, RowIndex int) *Line {
	var l Line
	l.RowIndex = RowIndex
	l.Prop = prop.NewProp(Row[propscorestxt.Prop], Row[propscorestxt.Par], Row[propscorestxt.Min], Row[propscorestxt.Max], 0)
	l.PropParType = propscorespartype.Types[Row[propscorestxt.PropParType]]
	l.ScoreMin, _ = strconv.Atoi(Row[propscorestxt.ScoreMin])
	l.ScoreMax, _ = strconv.Atoi(Row[propscorestxt.ScoreMax])
	if l.PropParType == propscorespartype.S {
		l.ScoreMin = l.ScoreMax
	}
	if Row[propscorestxt.ScoreLim] == "" {
		l.ScoreLimit = 1000
	} else {
		scoreLimit, ok := strconv.Atoi(Row[propscorestxt.ScoreLim])
		if ok != nil {
			log.Fatalf("propscores.NewLine: non-numeric ScoreLim on PropScores.txt line #%d", RowIndex)
		}
		l.ScoreLimit = scoreLimit
	}
	l.MinLvl, _ = strconv.Atoi(Row[propscorestxt.MinLvl])
	l.MaxLvl, _ = strconv.Atoi(Row[propscorestxt.MaxLvl])
	l.LvlScale = (Row[propscorestxt.LvlScale] == "Y")
	l.NoTypeOvr = (Row[propscorestxt.NoTypeOver] == "Y")
	for colIdx := propscorestxt.Itype1; colIdx < propscorestxt.Itype6; colIdx++ {
		itype := Row[colIdx]
		if itype != "" {
			l.Itypes = append(l.Itypes, strings.TrimSpace(itype))
		}
	}
	for colIdx := propscorestxt.Etype1; colIdx < propscorestxt.Etype3; colIdx++ {
		etype := Row[colIdx]
		if etype != "" {
			l.Etypes = append(l.Etypes, strings.TrimSpace(etype))
		}
	}
	l.Group = Row[propscorestxt.Group]
	l.SynergyGroup = Row[propscorestxt.SynergyGroup]
	l.SourceItem = Row[propscorestxt.SourceItem]
	l.SourceFile = Row[propscorestxt.SourceFile]
	return &l
}

// NewPropScoresIndex Reads PropScores.txt and populates PropScoresIndex,
// a structure containing Lines indexed by both Prop Name and by PropScores.txt Row #
func NewPropScoresIndex(d2files *d2fs.Files) *Maps {

	scorefile := d2files.GetAsset(propscorestxt.Path, propscorestxt.FileName)

	psi := Maps{}
	psi.PropLines = ScoreMap{}
	psi.RowLines = make([]*Line, len(scorefile.Rows))

	for rowIndex, r := range scorefile.Rows {
		line := NewLine(r, rowIndex)
		psi.RowLines[rowIndex] = line
		psi.PropLines[line.Prop.Name] = append(psi.PropLines[line.Prop.Name], line)
	}
	return &psi
}
