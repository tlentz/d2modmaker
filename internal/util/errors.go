package util

import "log"

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
