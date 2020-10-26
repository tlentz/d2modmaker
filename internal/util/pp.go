package util

import (
	"encoding/json"
	"fmt"
)

// PP is a pretty printer
func PP(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
