// Notes
// -----
//   - all go files are a part of a package
//   - if the package is called main, the go compiler knows that it should create
//     a standalone executable if being built
//   - ahared packages can't or well shouldn't use the `main` package name
package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func sayGreetings(n string) {
	fmt.Printf("Good morning, %v\n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye, %v\n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

// Please note: this function was defined differently 
// inside the tutorial and it honestly something
// I may want to think about replicating
func getInitials(n string) (string, string) {
	names := strings.Split(strings.ToUpper(n), " ")
	return names[0][:1], names[1][:1]
}

func getInitials2(n string) (string, string) {
	names := strings.Split(strings.ToUpper(n), " ")
	if(len(names) > 2) {
		return "too many names.", "please try again with just 2"
	} else if(len(names) == 1) {
		return names[0][:1], "_"
	}
	return names[0][:1], names[1][:1]
}

// functions that are called main are triggered automatically
// by the go compiler
// There must only be one `main` function
func main() {
	// -- LESSON #2
	// ----------------------------------------------------------
	fmt.Println("Hello World")

	// -- LESSON #3
	// ----------------------------------------------------------
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

	// -- LESSON #4
	// ----------------------------------------------------------
	age := 12
	fmt.Println()
	fmt.Print("hello, ")   // this doesn't added need line
	fmt.Print("world! \n") // this doesn't added need line

	fmt.Print("without format: My age is ", age, "\n")
	fmt.Printf("with format: My age is %v\n", age)

	var formattedString = fmt.Sprintf("My age is %v", age)
	fmt.Println("The formatted string: ", formattedString)

	// -- LESSON #5
	// ----------------------------------------------------------

	// var _name_of_variables [_arr_length]_type = [_arr_length]int{_values}
	var ages [3]int = [3]int{1, 2, 3}

	// var _name_of_variable = [_arr_length]_type{_values}
	var agesTwo = [3]int{4, 5, 6}

	// _name_of_variable := [_arr_length]_type{_values}
	names := [4]string{"Mary", "Samantha", "Andrew", "Phillipe"}
	names[1] = "Alex" // this is how you can the value

	fmt.Println(ages, agesTwo)
	fmt.Println("Names: ", names, " and No. of names: ", len(names))

	// SLICES (think mutatable arrays)
	var scores = []float64{100.0, 95.3, 12.4}
	scores[2] = 84.3
	scores = append(scores, 4.2)

	fmt.Println(scores)

	rangeOne := names[1:3]
	rangeTwo := names[1:]
	rangeThree := names[:3]

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	// -- LESSON #6
	// ----------------------------------------------------------

	greetings := "hello there friends"
	// PLease note that all the methods in strings package doesn't mutate the original string
	fmt.Println(strings.Contains(greetings, "hello"))
	fmt.Println(strings.ReplaceAll(greetings, "hello", "hey"))
	fmt.Println(strings.ToUpper(greetings))
	fmt.Println(strings.Index(greetings, "ll"))
	fmt.Println(strings.Split(greetings, " "))

	agesArr := []int{45, 20, 21, 64,8, 53, 96, 12}
	sort.Ints(agesArr)
	fmt.Println(agesArr)

	// Please note: when it doesn't find the index, it will return n + 1
	// where n is the number of items in the slice
	
	// There are also string and float version of the methods in sort 
	// package. For example: sort.Strings(_variable_name) & sort.Floats(_variable_name)
	index := sort.SearchInts(agesArr, 21)
	fmt.Println(index)

	// -- LESSON #7
	// ----------------------------------------------------------

	x := 0
	
	// the for keyword is to represent all forms of loops
	// below is an example of while loop
	for x < 5 {
		fmt.Println("The value of x is ", x)
		x++
	}
	fmt.Println()

	for i := 0; i < 5; i++ {
		fmt.Println("The value of i is ", i)
	}
	fmt.Println()

	namesArr := []string{"Mike", "Sam", "Mary", "Leefa", "lovely"}
	for i := 0; i < len(namesArr); i++ {
		fmt.Printf("The name at index %v is %v\n", i, namesArr[i])
	}

	// Please note: if you don't want to use index/value then you have
	// replace it with an underscore (_)

	// also updating the value in the for loop doesn't change the original
	// value in the slice since it creates a local copy of the value in the slice
	for index, value := range namesArr {
		fmt.Printf("The name at index %v is %v\n", index, value)
	}

	// -- LESSON #8
	// ----------------------------------------------------------

	// So for boolean expressions, they are exactly the same as in 
	// other programming languages. Those include the following 
	// examples: >, <, >=, <=, ==, != 

	age = 30 

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	}

	for index, value := range namesArr {
		if index == 1 {
			fmt.Println("continuing at index ", index)
			continue
		}

		if index > 2 {
			fmt.Println("breaking at index ", index)
			break
		}

		fmt.Printf("the value at index %v is %v\n", index, value)
	}

	// -- LESSON #9
	// ----------------------------------------------------------

	sayGreetings("Murphy")
	sayBye("Murphy")
	cycleNames(namesArr, sayGreetings)
	a1 := circleArea(12)
	a2 := circleArea(12.43)

	fmt.Println("The first area is ",a1)
	fmt.Println("The second area is ", a2)

	fmt.Printf("The first area is %0.3f\n", a1)
	fmt.Printf("The second area is %0.2f\n", a2)

	// -- LESSON #10
	// ----------------------------------------------------------
	firstInitial, lastInitial := getInitials("murphy facey")
	fmt.Printf("The initials are %v. %v.\n", firstInitial, lastInitial)
	
	firstInitial, lastInitial = getInitials2("mark wallis")
	fmt.Printf("The initials are %v. %v.\n", firstInitial, lastInitial)

	firstInitial, lastInitial = getInitials2("peter fray junior")
	fmt.Printf("The initials are %v. %v.\n", firstInitial, lastInitial)

	firstInitial, lastInitial = getInitials2("prince")
	fmt.Printf("The initials are %v. %v.\n", firstInitial, lastInitial)

	// -- LESSON #11
	// ----------------------------------------------------------
	sayHello("Murp")
}
