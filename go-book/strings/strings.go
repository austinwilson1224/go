package main

import (
	"fmt"
)

func main() {
	s := "hello, world"
	fmt.Println(s)

	// build in len function
	fmt.Println(len(s))

	// index operation
	fmt.Println(s[0], s[7])

	// array out of bounds
	// panic: runtime error: index out of range [12] with length 12
	fmt.Println(s[len(s)-2])

	// substrings
	fmt.Println(s[0:5])
	fmt.Println(s[:5])
	fmt.Println(s[7:])

	// str concatenation w/ + 
	fmt.Println("goodbye" + s[5:])

	s2 := "left foot"
	t := s2 
	s2 += ", right foot"
	fmt.Println(t)
	fmt.Println(s2)
	fmt.Println()
}
