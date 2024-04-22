package main

import "log"

func staticcheck() {
	a := 1
	// bad
	if a == 1 {
	}
	// good
	if a == 1 {
		log.Println(a)
	}
}
