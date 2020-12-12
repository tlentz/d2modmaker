package scorerstatistics

import (
	"log"

	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/armor"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/itemTypes"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/misc"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/pbucketlist"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/weapons"
)

func newpBucketMap(d2files *d2fs.Files) *pBucketMap {
	typeCodes := make(pBucketMap)
	/*
		// Below code is how I originally created the BucketList.txt
		// As I can't guarantee this will work with other mods, I wrote what this code built up into the  codes map
		// into the PBucketList.txt file.
		itemtypestxt := d2files.Get(itemTypes.FileName)
		for _, r := range itemtypestxt.Rows {
			if r[itemTypes.Equiv1] == "clas" || r[itemTypes.Equiv2] == "clas" {
				typeCodes[r[itemTypes.Code]] = r[itemTypes.Code]
			}
			bloc := r[itemTypes.BodyLoc1]
			if bloc == "rrin" || bloc == "neck" || bloc == "feet" || bloc == "glov" || bloc == "belt" {
				typeCodes[r[itemTypes.Code]] = r[itemTypes.Code]
			}
			if r[itemTypes.Equiv1] == "comb" {
				typeCodes[r[itemTypes.Code]] = "comb"
			}
			if r[itemTypes.Equiv1] == "armo" && r[itemTypes.Equiv2] == "" { // the "" is to exclude shld
				typeCodes[r[itemTypes.Code]] = r[itemTypes.Code]
			}
		}
		typeCodes["char"] = "char"
		typeCodes["wand"] = "wand"
		typeCodes["staf"] = "staf" // Staves are 2h, but they are more likely to get sorc bonuses
		typeCodes["circ"] = "circ"
		typeCodes["scep"] = "scep"
		typeCodes["shie"] = "shie"
		typeCodes["bow"] = "bow"
		typeCodes["xbow"] = "bow"
		typeCodes["pole"] = "2h"
		typeCodes["spea"] = "2h"
		typeCodes["club"] = "1h"
		typeCodes["mace"] = "1h"
		typeCodes["knif"] = "1h"
		typeCodes["tpot"] = "tpot" // this category should never be used TODO: Add away to not do this?

		// melee weapons are tricky to determine which are 1hand & which are 2hand, & which are both.
		// Grouping together all pure 2handers so that they have a higher chance at getting AC
		weaponstxt := d2files.Get(weapons.FileName)
		for _, r := range weaponstxt.Rows {
			wtype := r[weapons.Type]
			if wtype == "axe" || wtype == "swor" || wtype == "hamm" {
				switch {
				case r[weapons.Oneortwohanded] == "1":
					typeCodes[r[weapons.Code]] = "1h"
				case r[weapons.Twohanded] == "1":
					typeCodes[r[weapons.Code]] = "2h"
				default:
					typeCodes[r[weapons.Code]] = "1h"
				}
			}
		}
	*/
	pb := make(pBucketMap)

	// Instead of the above generation code, just read it in.
	pBucketFile := d2files.GetWithPath(pbucketlist.Path, pbucketlist.FileName)
	for _, row := range pBucketFile.Rows {
		typeCodes[row[pbucketlist.TypeCode]] = row[pbucketlist.PBucket]
	}

	itemtypestxt := d2files.Get(itemTypes.FileName)
	added := true
	for added {
		added = false
		for _, r := range itemtypestxt.Rows {
			if r[itemTypes.Code] == "" {
				continue
			}
			if typeCodes[r[itemTypes.Code]] != "" {
				// already added
				continue
			}
			if r[itemTypes.Equiv1] != "" {
				if typeCodes[r[itemTypes.Equiv1]] != "" {
					typeCodes[r[itemTypes.Code]] = typeCodes[r[itemTypes.Equiv1]]
					//fmt.Printf("added1: %s %s\n", r[itemTypes.Code], typeCodes[r[itemTypes.Equiv1]])
					added = true
				}
			}
			if r[itemTypes.Equiv2] != "" {
				if typeCodes[r[itemTypes.Equiv2]] != "" {
					typeCodes[r[itemTypes.Code]] = typeCodes[r[itemTypes.Equiv2]]
					//fmt.Printf("added2: %s %s\n", r[itemTypes.Code], typeCodes[r[itemTypes.Equiv2]])
					added = true
				}
			}
		}
	}

	// finally the real deal.. assign every Type in weapons.txt to a pbucket
	weaponstxt := d2files.Get(weapons.FileName)
	for _, r := range weaponstxt.Rows {
		wtype := r[weapons.Type]
		wcode := r[weapons.Code]
		if wtype == "" {
			continue
		}
		if pb[wtype] == "" {
			switch {
			case typeCodes[wcode] != "":
				pb[wcode] = typeCodes[wcode]
			case typeCodes[wtype] != "":
				pb[wcode] = typeCodes[wtype]
			default:
				log.Fatalf("Bucketlist missing type %s or code %s\n", wtype, wcode)
			}
		}
	}

	armortxt := d2files.Get(armor.FileName)
	for _, r := range armortxt.Rows {
		atype := r[armor.Type]
		acode := r[armor.Code]
		if acode == "" {
			continue
		}
		if pb[atype] == "" {
			switch {
			case typeCodes[acode] != "":
				pb[acode] = typeCodes[acode]
			case typeCodes[atype] != "":
				pb[acode] = typeCodes[atype]
			default:
				log.Fatalf("Bucketlist missing type %s or code %s", atype, acode)
			}
		}
	}
	misctxt := d2files.Get(misc.FileName)
	for _, r := range misctxt.Rows {
		itype := r[misc.Type_]
		icode := r[misc.Code]
		if icode == "" {
			continue
		}
		if pb[itype] == "" {
			switch {
			case typeCodes[icode] != "":
				pb[icode] = typeCodes[icode]
			case typeCodes[itype] != "":
				pb[icode] = typeCodes[itype]
			default:
				// Don't bomb because there are many item types we won't roll for in misc.txt
				//log.Fatalf("Bucketlist missing type %s or code %s", itype, icode)
			}
		}
	}

	for _, row := range pBucketFile.Rows {
		pb[row[pbucketlist.TypeCode]] = row[pbucketlist.PBucket]
	}

	// for key, r := range typeCodes {
	// 	fmt.Printf("%s\t%s\n", key, r)
	// }
	//fmt.Printf("%+v\n%+v\n", pb, typeCodes)
	return &pb
}

// // GetPbucket returns item bucket name
// func GetPbucket(sspb *scorerstatistics.pBucketMap, itemCode string) string {
// 	return (*sspb.pBucket)[itemCode]
// }
