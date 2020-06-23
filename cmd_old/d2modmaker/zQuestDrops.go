package main

func QuestDrops(d2file *D2File) (*D2File, error) {

	pairs := make(map[string]string)
	pairs["Andariel"] = "Andarielq"
	pairs["Andariel (N)"] = "Andarielq (N)"
	pairs["Andariel (H)"] = "Andarielq (H)"

	pairs["Duriel - Base"] = "Durielq - Base"
	pairs["Duriel (N) - Base"] = "Durielq (N) - Base"
	pairs["Duriel (H) - Base"] = "Durielq (H) - Base"

	pairs["Mephisto"] = "Mephistoq"
	pairs["Mephisto (N)"] = "Mephistoq (N)"
	pairs["Mephisto (H)"] = "Mephistoq (H)"

	pairs["Diablo"] = "Diabloq"
	pairs["Diablo (N)"] = "Diabloq (N)"
	pairs["Diablo (H)"] = "Diabloq (H)"

	pairs["Baal"] = "Baalq"
	pairs["Baal (N)"] = "Baalq (N)"
	pairs["Baal (H)"] = "Baalq (H)"

	for idx, record := range d2file.Records {
		tc, ok := record["Treasure Class"]
		if ok && tc != "" {
			qKey, ok2 := pairs[tc]
			if ok2 {
				qItemIdx, err := GetItemFromRecords(d2file, "Treasure Class", qKey)
				if err == nil {
					qRow := d2file.Records[*qItemIdx]
					newRow := copyMap(qRow)
					newRow["Treasure Class"] = tc
					d2file.Records[idx] = newRow
				}
			}
		}
	}

	return d2file, nil
}
