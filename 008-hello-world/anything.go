package main

import "fmt"

func main() {
	n, _ := fmt.Println("WOW THIS IS SO GREAT!",42, true)
	fmt.Println(n)
	// fmt.Println(err)
	foo()
	fmt.Println("whatever")
	for i:= 0; i < 100; i++ {
		if i % 2 == 0 {
			// fmt.Println(i)
		}
	}
	bar()
	fmt.Println("a" + "b")
}

func foo() {
	fmt.Println("i'm in foo")
}

func bar() {
	fmt.Println("we exited")
}

func throwaway() {

	// the underscore is for and unused return value 
	numberOfBytes, _ := fmt.Println("this is a test", true, 42)
	fmt.Println(numberOfBytes)
}
// control flow:
// (1) sequence
// (2) loop; iteration
// (3) conditional