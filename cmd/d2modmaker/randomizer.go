package main

import (
	"github.com/tlentz/d2modmaker/internal/d2file"
	"github.com/tlentz/d2modmaker/internal/uniqueItemsTxt"
	"github.com/tlentz/d2modmaker/internal/util"
)

func Randomize(d2files *d2file.D2Files) {
	getAllProps(d2files)
}

func getAllProps(d2files *d2file.D2Files) {
	getAllUniqueProps(d2files)
}

func getAllUniqueProps(d2files *d2file.D2Files) {
	f := d2file.GetOrCreateFile(dataDir, d2files, uniqueItemsTxt.FileName)
	propOffset := uniqueItemsTxt.Prop1
	props := []uniqueItemsTxt.UProp{}
	for _, row := range f.Rows {
		for i := propOffset; i < len(row)-propOffset; i += 4 {
			props = append(props, uniqueItemsTxt.UProp{
				UPropProp: row[i],
				UPropPar:  row[i+1],
				UPropMin:  row[i+2],
				UPropMax:  row[i+3],
			})
		}
	}
	util.PP(props)
}

func getAllSetProps() {

}

func getAllRWProps() {

}
