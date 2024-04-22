package main

import (
	"errors"
	"fmt"
	"log"
)

func errorlint() {
	// bad
	err := fmt.Errorf("wrapper for %v", errors.New("st went wrong"))
	log.Println(err)
	// good
	errGood := fmt.Errorf("wrapper for %w", errors.New("st went wrong"))
	log.Println(errGood)
}
