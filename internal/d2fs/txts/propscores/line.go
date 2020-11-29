package propscores

import (
	"path"
	"strconv"
	"strings"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/propscores/propscorespartype"
	"github.com/tlentz/d2modmaker/internal/d2mod/prop"
)

// Line Contains 1 line from PropScore.txt
type Line struct {
	RowIndex     int
	Prop         prop.Prop
	PropParType  int // See constants in propscorepartype
	ScoreMin     int
	ScoreMax     int
	MinLvl       int
	NoTypeOvr    bool
	Itypes       []string
	Etypes       []string
	Group        string
	SynergyGroup string
	SourceItem   string
	SourceFile   string
}

// ScoreMap map of Prop Name's to  PropScores.txt Lines
// Beware that this list is only prop name, not prop&par, so there will be a lot of duplicates for things like "skill"
// The list must be parsed with level, etype, itype restrictions etc. to find the right Line.
type ScoreMap map[string][]*Line // Used to grab all the Lines materialized from PropScores.txt for a given prop name

type PropScoresIndex struct {
	PropLines ScoreMap // Map from Prop name to  array of PropScore.txt []Lines
	RowLines  []*Line  // Array of Lines by RowIdx i.e. matching d2files.Rows[]
}

// NewLine Create a new Line from a given row in PropScores.txt
func NewLine(Row []string, RowIndex int) *Line {
	var l Line
	l.RowIndex = RowIndex
	l.Prop = prop.NewProp(Row[Prop], Row[Par], Row[Min], Row[Max])
	l.PropParType = propscorespartype.Types[Row[PropParType]]
	l.ScoreMin, _ = strconv.Atoi(Row[ScoreMin])
	l.ScoreMax, _ = strconv.Atoi(Row[ScoreMax])
	l.MinLvl, _ = strconv.Atoi(Row[MinLvl])
	l.NoTypeOvr = (strings.Compare(Row[NoTypeOver], "Y") == 0)
	for colIdx := Itype1; colIdx < Itype6; colIdx++ {
		itype := Row[colIdx]
		if itype != "" {
			l.Itypes = append(l.Itypes, strings.TrimSpace(itype))
		}
	}
	for colIdx := Etype1; colIdx < Etype3; colIdx++ {
		etype := Row[colIdx]
		if etype != "" {
			l.Etypes = append(l.Etypes, strings.TrimSpace(etype))
		}
	}
	l.Group = Row[Group]
	l.SynergyGroup = Row[SynergyGroup]
	l.SourceItem = Row[SourceItem]
	l.SourceFile = Row[SourceFile]
	return &l
}

// NewPropScoresIndex Reads PropScores.txt and populates PropScoresIndex,
// a structure containing Lines indexed by both Prop Name and by PropScores.txt Row #
func NewPropScoresIndex(d2files *d2fs.Files) *PropScoresIndex {

	scorefile := d2files.Read(path.Join(Path, FileName))

	psi := PropScoresIndex{}
	psi.PropLines = ScoreMap{}
	psi.RowLines = make([]*Line, len(scorefile.Rows))

	for rowIndex, r := range scorefile.Rows {
		line := NewLine(r, rowIndex)
		psi.RowLines[rowIndex] = line
		psi.PropLines[line.Prop.Name] = append(psi.PropLines[line.Prop.Name], line)
	}
	return &psi
}
