package main

import (
	"log"
)

func checkError(err error) {
	if nil != err {
		log.Fatalf("ERROR: %v\n", err)
	}
}
