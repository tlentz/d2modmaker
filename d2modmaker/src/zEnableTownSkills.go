package main

func EnableTownSkills(d2file *D2File) (*D2File, error) {

	var skills = []string{"Teleport"}

	PP(skills)

	for _, skill := range skills {
		skillIdx, err := GetItemFromRecords(d2file, "skill", skill)
		CheckD2FileErr(d2file, err)
		d2file.Records[*skillIdx]["InTown"] = "1"
	}

	return d2file, nil
}
