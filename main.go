package main

import (
	"fmt"
	"math"
)

// Hello World
func helloWorld() {
	fmt.Println("Hello World")
}

/*

Go has various value types including strings, integers, floats, booleans, etc.
Here are a few basic examples.

*/
func values() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1=", 1+1)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

/*


In Go, variables are explicitly declared
and used by the compiler to e.g. check
 type-correctness of function calls.
*/
func variables() {
	//var declares one or more variables
	var a = "initial"
	fmt.Println(a)
	//You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(c, b)
	//Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)
	//Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
	var e int
	fmt.Println(e)
	//The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.
	random_variable_name := "apple"
	fmt.Println(random_variable_name)

}

/*

Go supports constants of character, string, boolean, and numeric values.

*/
func constants() {
	//Import math at the top
	//A const statement can appear anywhere a var statement can.
	const s string = "I am a constant"
	fmt.Println(s, "string")
	//Constant expressions perform arithmetic with arbitrary precision.
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d, "D")
	//A numeric constant has no type until it’s given one, such as by an explicit conversion.
	fmt.Println(int64(d))
	//A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example,
	// here math.Sin expects a float64.
	fmt.Println(math.Sin(n))
}

/*

for is Go’s only looping construct. Here are three basic types of for loops.

*/
func forLoop() {
	//The most basic type, with a single condition.
	index := 1
	for index <= 3 {
		fmt.Println(index)
		index = index + 1
	}
	//A classic initial/condition/after for loop.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	//for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
	for {
		fmt.Println("Loop")
		break

		//You can also continue to the next iteration of the loop.
	}
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

/*
If else statements are not that complicated
should not have to explain

*/
func ifElseStatements() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	if 8%4 == 0 {
		fmt.Println("basic if statement")
	}
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

func main() {
	//All of the functions will be called in here
	// helloWorld()
	// values()
	// variables()
	// constants()
	// forLoop()
	// ifElseStatements()
}
