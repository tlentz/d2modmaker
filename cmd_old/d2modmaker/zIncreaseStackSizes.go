package main

func IncreaseStackSizes(d2file *D2File) (*D2File, error) {

	// Sets TP Book size to 100
	tpBookIdx, err := GetItemFromRecords(d2file, "name", "Town Portal Book")
	CheckD2FileErr(d2file, err)
	d2file.Records[*tpBookIdx]["maxstack"] = "100"

	// Sets Id Book size to 100
	idBookIdx, err := GetItemFromRecords(d2file, "name", "Identify Book")
	CheckD2FileErr(d2file, err)
	d2file.Records[*idBookIdx]["maxstack"] = "100"

	return d2file, nil
}
