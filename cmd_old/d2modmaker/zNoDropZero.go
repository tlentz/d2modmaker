package main

func NoDropZero(d2file *D2File) (*D2File, error) {

	for idx := 0; idx < len(d2file.Records); idx++ {
		val, ok := d2file.Records[idx]["NoDrop"]
		if ok && val != "" {
			d2file.Records[idx]["NoDrop"] = "0"
		}
	}

	return d2file, nil
}
