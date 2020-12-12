package d2items

import (
	"log"
	"strings"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/armor"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/itemTypes"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/weapons"
)

// TypeTree Contains inheritance information of ItemTypes read from Armor.txt, Weapons.txt & ItemTypes.txt
type TypeTree struct {
	parentItemType map[string][]string // From Armor, Weapons, itemtype: Maps membership of items in groups for doing itype/etype calcs
}

// NewTypeTree Read in Armor.txt, Weapons.txt & ItemTypes.txt and return populated TypeTree
func NewTypeTree(d2files *d2fs.Files) *TypeTree {
	tt := TypeTree{}
	tt.parentItemType = map[string][]string{}
	// Add Type Maps from Armor.txt
	armortxt := d2files.Get(armor.FileName)
	for _, r := range armortxt.Rows {
		tt.addTypeParent(r[armor.Normcode], r[armor.Type])
		tt.addTypeParent(r[armor.Ubercode], r[armor.Type])
		tt.addTypeParent(r[armor.Ultracode], r[armor.Type])
	}

	// Add Type Maps from Weapons.txt
	weaponstxt := d2files.Get(weapons.FileName)
	for _, r := range weaponstxt.Rows {
		tt.addTypeParent(r[weapons.Normcode], r[weapons.Type])
		tt.addTypeParent(r[weapons.Ubercode], r[weapons.Type])
		tt.addTypeParent(r[weapons.Ultracode], r[weapons.Type])
	}

	// Add Type Maps from ItemTypes.txt
	itemTypestxt := d2files.Get(itemTypes.FileName)
	for _, r := range itemTypestxt.Rows {
		// fmt.Printf("%s | %s | %s", r[itemTypes.ItemType], r[itemTypes.Code], r[itemTypes.Equiv1])
		tt.addTypeParent(r[itemTypes.Code], r[itemTypes.Equiv1])
		tt.addTypeParent(r[itemTypes.Code], r[itemTypes.Equiv2])
	}

	return &tt
}

// AddTypeParent Add child/parent relationship to TypeMap
func (tt *TypeTree) addTypeParent(child string, parent string) {
	if (child == "") || (parent == "") {
		return
	}
	if strings.TrimSpace(child) != child {
		log.Fatalf("addTypeMap: Bad child string [%s]", child)
	}
	if strings.TrimSpace(parent) != parent {
		log.Fatalf("addTypeMap: Bad parent string [%s]", parent)
	}
	if child == parent {
		return
	}
	for _, p := range tt.parentItemType[child] {
		if p == parent {
			return // New parent already exists for the child, bail
		}
	}
	tt.parentItemType[child] = append(tt.parentItemType[child], parent)
}

// CheckTypeTree  Returns if there is a path from child->parent in TypeTree
// ***** BEWARE RECURSION ******
func CheckTypeTree(tt *TypeTree, child string, parent string) bool {
	if parent != strings.TrimSpace(parent) {
		log.Fatalf("checkTypeMap: bad type string (extra spaces): [%s] [%s]\n", child, parent)
	}
	//fmt.Printf("checkTypeMap: Checking %s %s\n", child, parent)
	if (child == "") || (parent == "") {
		return false
	}
	if child == parent {
		//fmt.Printf("checkTypeMap: Match:%s %s\n", child, parent)
		return true
	}
	plist := tt.parentItemType[child]
	for _, p := range plist {
		if p == parent {
			return true
		}
		if CheckTypeTree(tt, p, parent) {
			return true
		}
	}
	return false
}

// CheckTwoHander Check if an "Item" is 2hand only (gotta be careful that runes have multiple types)
func CheckTwoHander(tt *TypeTree, i Item) bool {
	if len(i.Types) == 0 {
		return false
	}
	twohanders := []string{"pole", "staf", "bow", "xbow", "abow", "aspe", "spea",
		// OBC: Friggen 1handed hammers and axes make me do this
		// TODO instead of all this garbage check 2handed vs 1or2hand columns in Weapons.txt
		"lax", "bax", "btx", "gax", "gix", "9la", "9ba", "9bt", "9ga", "9gi",
		"7la", "7ba", "7bt", "7ga", "7gi", "mau", "gma", "9m9", "9gm", "7m7", "7gm",
	}
NextItemType:
	for _, itemtype := range i.Types {
		for _, thtype := range twohanders {
			if CheckTypeTree(tt, itemtype, thtype) {
				continue NextItemType
			}
		}
		return false
	}
	return true
}

// CheckIETypes Check that any of item.Types is a child of propscores.Line.Itypes and not a child of propscore.Line.Etypes
func CheckIETypes(tt *TypeTree, itemTypes []string, Itypes []string, Etypes []string) bool {
	//fmt.Printf("checkIETypes: %s/%s/%s\n", itemTypes, Itypes, Etypes)

	isExcluded := false
	if len(Etypes) > 0 {
		isExcluded = true
	}
	for _, itemType := range itemTypes {
		for _, etype := range Etypes {
			if !CheckTypeTree(tt, itemType, etype) && (etype != "") {
				//fmt.Printf("efail\n")
				isExcluded = false
			}
		}
	}
	if isExcluded {
		//fmt.Printf("isExcluded\n")
		return false
	}

	if len(Itypes) == 0 {
		return true
	}
	for _, itemType := range itemTypes {
		for _, itype := range Itypes {
			if CheckTypeTree(tt, itemType, itype) {
				//fmt.Printf("isucceed\n")
				return true
			}
		}
	}
	//fmt.Printf("ifail(%s -> %s/%s checked)\n", itemTypes, Itypes, Etypes)
	return false // has item.Types but no match
}
