package main

func lll() {
	fooLLL := func(argumentName1 string, argumentName2 string, argumentName3 string) (nakedReturn1 string, nakedReturn2 string) {
		return
	}
	_ = fooLLL

	fooWithoutLLL := func(argumentName1 string, argumentName2 string, argumentName3 string) (
		nakedReturn1 string, nakedReturn2 string) {
		return
	}
	_ = fooWithoutLLL
}
