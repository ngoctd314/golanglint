package main

import "errors"

func nilerr() {
	baz := func() error {
		err := errors.New("nilerr")
		if err != nil {
			return nil // good: return err or comment: "lint:ignore nilerr because biz logic" before return
		}
		if err != nil {
			// lint:ignore nilerr because biz logic
			return nil
		}
		return err
	}
	_ = baz()
}
