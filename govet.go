package main

import "log"

func govet() {
	// bad
	log.Printf("your name is %d", "admin")
	// correct
	log.Printf("your name is %s", "admin")
}
