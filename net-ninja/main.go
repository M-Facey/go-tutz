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
}
