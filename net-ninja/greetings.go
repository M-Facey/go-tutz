// so the way you use functions or variables from
// for other files is that you first ensure that
// the package names are the same. Then you need
// to run both files together.
package main

import "fmt"

var points = []int{20,90,30,41, 53, 12,45}

func sayHello(n string) {
	fmt.Println("Hello", n)
}