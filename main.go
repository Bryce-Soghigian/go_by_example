package main

import (
	"fmt"
)

//Hello World
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

func main() {
	//All of the functions will be called in here
	helloWorld()
	values()
}
