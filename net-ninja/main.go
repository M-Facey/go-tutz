// Notes
// -----
//   - all go files are a part of a package
//   - if the package is called main, the go compiler knows that it should create
//     a standalone executable if being built
//   - ahared packages can't or well shouldn't use the `main` package name
package main

import "fmt"

// functions that are called main are triggered automatically
// by the go compiler
// There must only be one `main` function
func main() {
	fmt.Println("Hello World")

	// strings
	// 1: var _name_of_variable _type = _value (explicit)
	var nameOne string = "Mario"

	// 2: var _name_of_variable = _value (implicit)
	var nameTwo = "Luigi"

	// 3: var _name_of_variables _type (use default value if none was provided)
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Peach"
	nameThree = "Browser"

	fmt.Println(nameOne, nameTwo, nameThree)

	// 4: _name_of_variable := _value (implicit & magical)
	nameFour := "Yoshi"
	fmt.Println(nameFour)

	// regular ints
	var ageOne int = 27
	var ageTwo = 30
	ageThree := 12

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory
	var numOne int8 = 25
	var numTwo int8 = -128 // negative number included, -128 - 128
	var numThree uint = 10 // no negative number, 0 - 255

	fmt.Println(numOne, numTwo, numThree)

	// you need to specitfy version 32 or 64 for floats
	// the := defaults to using the float64 version
	var scoreOne float32 = 23.4
	fmt.Println(scoreOne)
}
