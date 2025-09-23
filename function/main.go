package main

import "fmt"

// import "fmt"

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}
func getNumbers(num1 int, num2 int) (int,int)  {
	sum := num1 + num2
	mul := num1 * num2
	return  sum , mul
}

func printSomething()  {
	println("This is Just print something")
}
func sayHello(name string)  {
	fmt.Println("Welcome to the go lang course ,",name)
}

func main() {
	// a := 10
	// b := 20
	// p,q := getNumbers(a,b)
	// fmt.Println(p)
	// fmt.Println(q)
	printSomething()
	sayHello("Jisan")
}