package main

import "fmt"

// standard and name function
// func add(a int, b int)  {
// 	fmt.Println(a + b)
// }

func main() {
	// anonymous Function
	// Immediately Invoked Function Expression
	// IIFE
	// Invoke Mean -> call/execute
	func(a int, b int) {
		c := a + b
		fmt.Println(c)
	}(3, 4)

	// Function expression or assign function in variable

	sub := func(a int, b int) {
		fmt.Println(a - b)
	}

	sub(4,2)

}

func init() {
	fmt.Println("I will be called first")
}
