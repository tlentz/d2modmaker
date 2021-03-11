package ignore

import (
	"log"
	"strconv"
	"strings"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/filenumbers"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/ignoretxt"
)

var ignoreMap map[string]bool
var isInitialized = false

// IsIgnored Returns true if the specified filenumber/type/name is in Ignore.txt
func IsIgnored( /*d2files *d2fs.Files,*/ fileNumber int, typeToCheck string, nameToCheck string) bool {
	//fmt.Printf("IsIgnored:%d %s %s\n", fileNumber, typeToCheck, nameToCheck)
	//log.Printf("%+v\n", ignoreMap)
	switch typeToCheck {
	case "prop", "item":
		return ignoreMap[strconv.Itoa(fileNumber)+"/"+typeToCheck+"/"+nameToCheck]
	default:
		log.Panicf("unknown typeToCheck argument: contact developer")
		return false
	}
}

// Init .
func Init(d2files d2fs.Files) {
	ignoreMap = make(map[string]bool, 0)
	f := d2files.Get(ignoretxt.FileName)
	for rowIdx := range f.Rows {
		fnumber := filenumbers.NameToNumberMap[strings.ToLower(f.Rows[rowIdx][ignoretxt.File])]
		if fnumber == 0 {
			log.Panicf("Filename %s isn't recognized\n", f.Rows[rowIdx][ignoretxt.File])
		}
		ignoreMap[strconv.Itoa(fnumber)+"/"+f.Rows[rowIdx][ignoretxt.Type]+"/"+f.Rows[rowIdx][ignoretxt.Name]] = true
	}
}
