package main

import "fmt"

func main() {
	fmt.Println("WOW THIS IS SO GREAT!")
	foo()
	fmt.Println("whatever")
	for i:= 0; i < 100; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("i'm in foo")
}

func bar() {
	fmt.Println("we exited")
}
// control flow:
// (1) sequence
// (2) loop; iteration
// (3) conditional