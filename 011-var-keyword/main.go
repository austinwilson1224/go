package main

import (
	"fmt"
)

var name string = "Austin Wilson"

// declare there is a variable withthe identifier "z"
// ASSIGNS the zero value of TYPE INTO to "z" 
var z int 

func main() {
	// short declaration operatro
	x := 42
	fmt.Println(x)

	// using var
	var y = 43
	fmt.Println(y)


	fmt.Println(name)
	foo()
}

func foo() {
	fmt.Println("again:", y)
}