package main

import "strconv"

func IncreaseMonsterDensity(d2file *D2File, m int) {
	var monDen = "MonDen"
	var monUMin = "MonUMin"
	var monUMax = "MonUMax"
	var maxDensity = 10000
	var maxM = 30
	multiplier := Min(maxM, m)

	var difficulties = []string{"", "(N)", "(H)"}
	for idx, _ := range d2file.Records {
		for _, diff := range difficulties {

			// MonDen
			oldMonDen, err := strconv.Atoi(d2file.Records[idx][monDen+diff])
			if err == nil {
				newMonDen := Min(maxDensity, oldMonDen*multiplier)
				d2file.Records[idx][monDen+diff] = strconv.Itoa(newMonDen)
			}

			// MonUMin
			oldMonUMin, err := strconv.Atoi(d2file.Records[idx][monUMin+diff])
			if err == nil {
				newMonUMin := oldMonUMin * multiplier
				d2file.Records[idx][monUMin+diff] = strconv.Itoa(newMonUMin)
			}

			// MonUMax
			oldMonUMax, err := strconv.Atoi(d2file.Records[idx][monUMax+diff])
			if err == nil {
				newMonUMax := oldMonUMax * multiplier
				d2file.Records[idx][monUMax+diff] = strconv.Itoa(newMonUMax)
			}
		}
	}
}
