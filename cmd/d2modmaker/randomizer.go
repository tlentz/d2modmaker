package main

import (
	"fmt"

	"github.com/tlentz/d2modmaker/internal/d2file"
	"github.com/tlentz/d2modmaker/internal/runesTxt"
	"github.com/tlentz/d2modmaker/internal/setsTxt"
	"github.com/tlentz/d2modmaker/internal/uniqueItemsTxt"
)

type Prop struct {
	Name string
	Par  string
	Min  string
	Max  string
}
type Props = []Prop

func Randomize(d2files *d2file.D2Files) {
	props := getAllProps(d2files)
	fmt.Println(len(props))
	// util.PP(props)
}

func getAllProps(d2files *d2file.D2Files) Props {
	props := Props{}

	// uniques
	uniqueProps := getAllUniqueProps(d2files, []Prop{})
	props = append(props, uniqueProps...)

	// sets
	setProps := getAllSetProps(d2files, []Prop{})
	props = append(props, setProps...)

	// rw
	rwProps := getAllRWProps(d2files, []Prop{})
	props = append(props, rwProps...)

	return props
}

func getAllUniqueProps(d2files *d2file.D2Files, props Props) Props {
	f := d2file.GetOrCreateFile(dataDir, d2files, uniqueItemsTxt.FileName)
	propOffset := uniqueItemsTxt.Prop1
	for _, row := range f.Rows {
		for i := propOffset; i < len(row)-propOffset; i += 4 {
			props = append(props, Prop{
				Name: row[i],
				Par:  row[i+1],
				Min:  row[i+2],
				Max:  row[i+3],
			})
		}
	}
	return props
}

func getAllSetProps(d2files *d2file.D2Files, props Props) Props {
	f := d2file.GetOrCreateFile(dataDir, d2files, setsTxt.FileName)
	propOffset := setsTxt.PCode2a
	for _, row := range f.Rows {
		for i := propOffset; i < len(row)-propOffset; i += 4 {
			props = append(props, Prop{
				Name: row[i],
				Par:  row[i+1],
				Min:  row[i+2],
				Max:  row[i+3],
			})
		}
	}
	return props
}

func getAllRWProps(d2files *d2file.D2Files, props Props) Props {
	f := d2file.GetOrCreateFile(dataDir, d2files, runesTxt.FileName)
	propOffset := runesTxt.T1Code1
	for _, row := range f.Rows {
		for i := propOffset; i < len(row)-propOffset; i += 4 {
			props = append(props, Prop{
				Name: row[i],
				Par:  row[i+1],
				Min:  row[i+2],
				Max:  row[i+3],
			})
		}
	}
	return props
}
