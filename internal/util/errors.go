package util

import (
	"fmt"
)

// CheckError checks for err, log message/err
func CheckError(message string, err error) {
	if err != nil {
		fmt.Println(message, err)
		panic(err)
	}
}

// Check for err, panic if exists
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
