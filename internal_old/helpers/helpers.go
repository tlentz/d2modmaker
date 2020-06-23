package helpers

import (
	"encoding/json"
	"fmt"
	"log"
)

// CheckError checks for err, log message/err
func CheckError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

// Check for err, panic if exists
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Max Returns the max between 2 ints
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min Returns the min between 2 ints
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// PP is a pretty printer
func PP(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

// ContainsString checks if a slice contains a given string
func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
