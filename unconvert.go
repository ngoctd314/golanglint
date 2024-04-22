package main

func unconvert() {
	var a int64
	foo := func(x int64) {}
	foo(int64(a))
	// good
	foo(a)
}
