package main

import "errors"

func errcheck() {
	foo := func() error {
		return errors.New("errcheck")
	}
	foo()
	// good
	_ = foo()
}
