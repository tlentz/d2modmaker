package reqs

import (
	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/armor"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/autoMagic"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/magicPrefix"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/magicSuffix"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/misc"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/setItems"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/uniqueItems"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/weapons"
)

func RemoveLevelRequirements(d2files d2fs.Files) {
	armortxt := d2files.Get(armor.FileName)
	for i := range armortxt.Rows {
		armortxt.Rows[i][armor.Levelreq] = ""
	}

	amagic := d2files.Get(autoMagic.FileName)
	for i := range amagic.Rows {
		amagic.Rows[i][autoMagic.Levelreq] = ""
		amagic.Rows[i][autoMagic.Classlevelreq] = ""
	}
	magicpref := d2files.Get(magicPrefix.FileName)
	for i := range magicpref.Rows {
		magicpref.Rows[i][magicPrefix.LevelReq] = ""
		magicpref.Rows[i][magicPrefix.ClassLevelReq] = ""
	}
	magicsuf := d2files.Get(magicSuffix.FileName)
	for i := range magicsuf.Rows {
		magicsuf.Rows[i][magicSuffix.LevelReq] = ""
		magicsuf.Rows[i][magicSuffix.ClassLevelReq] = ""
	}

	misctxt := d2files.Get(misc.FileName)
	for i := range misctxt.Rows {
		misctxt.Rows[i][misc.LevelReq] = ""
	}

	sets := d2files.Get(setItems.FileName)
	for i := range sets.Rows {
		sets.Rows[i][setItems.LvlReq] = ""
	}

	uitems := d2files.Get(uniqueItems.FileName)
	for i := range uitems.Rows {
		uitems.Rows[i][uniqueItems.LvlReq] = ""
	}

	weps := d2files.Get(weapons.FileName)
	for i := range weps.Rows {
		weps.Rows[i][weapons.Levelreq] = ""
	}

}

func RemoveAttRequirements(d2files d2fs.Files) {

	armortxt := d2files.Get(armor.FileName)
	for i := range armortxt.Rows {
		armortxt.Rows[i][armor.Reqstr] = ""
	}

	weptxt := d2files.Get(weapons.FileName)
	for i := range weptxt.Rows {
		weptxt.Rows[i][weapons.Reqstr] = ""
		weptxt.Rows[i][weapons.Reqdex] = ""
	}

}
