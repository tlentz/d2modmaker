package main

import (
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

const dataDir = "./assets/113c-data/"
const outDir = "./dist/"

func main() {
	var cfg = ReadCfg()
	makeMod(cfg)
}

func makeMod(cfg ModConfig) {
	var d2files = map[string]D2File{}

	PP(cfg)

	if cfg.IncreaseStackSizes {
		d2file := GetOrCreateFile(&d2files, "Misc.txt")
		IncreaseStackSizes(d2file)
	}

	if cfg.IncreaseMonsterDensity > 1 {
		d2file := GetOrCreateFile(&d2files, "Levels.txt")
		IncreaseMonsterDensity(d2file, cfg.IncreaseMonsterDensity)
	}

	// if cfg.LinearRuneDrops {
	// 	GetOrCreateFile(&d2files, "TreasureClassEx.txt")
	// }

	if cfg.EnableTownSkills {
		d2file := GetOrCreateFile(&d2files, "Skills.txt")
		EnableTownSkills(d2file)
	}

	if cfg.NoDropZero {
		d2file := GetOrCreateFile(&d2files, "TreasureClassEx.txt")
		NoDropZero(d2file)
	}

	if cfg.QuestDrops {
		d2file := GetOrCreateFile(&d2files, "TreasureClassEx.txt")
		QuestDrops(d2file)
	}

	WriteFiles(&d2files)
}

func runGUI() {
	cfg := ModConfig{}

	myApp := app.New()
	myWindow := myApp.NewWindow("Choice Widgets")

	check1 := widget.NewCheck("Increase Stack Sizes", func(value bool) {
		cfg.IncreaseStackSizes = value
	})
	check2 := widget.NewCheck("Enable Town Skills", func(value bool) {
		cfg.EnableTownSkills = value
	})
	check3 := widget.NewCheck("No Drop = 0", func(value bool) {
		cfg.NoDropZero = value
	})

	slider1 := widget.NewSlider(1, 30)
	slider1Label := widget.NewLabel("Increase Monster Density by " + strconv.Itoa(int(slider1.Value)) + "x")
	monsterDensityBox := widget.NewHBox(slider1Label, slider1)

	// radio := widget.NewRadio([]string{"Option 1", "Option 2"}, func(value string) {
	// 	log.Println("Radio set to", value)
	// })
	// combo := widget.NewSelect([]string{"Option 1", "Option 2"}, func(value string) {
	// 	log.Println("Select set to", value)
	// })

	myWindow.SetContent(widget.NewVBox(check1, check2, check3, monsterDensityBox))
	myWindow.Resize(fyne.NewSize(600, 600))
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()
}
