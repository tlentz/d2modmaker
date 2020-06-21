package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func CheckError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func PP(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
