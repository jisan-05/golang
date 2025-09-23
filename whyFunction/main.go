package main

import "fmt"

func printWelcomeMessage() {
	fmt.Println("Welcome to the function")
}
func getUserName() string {
	var name string
	fmt.Println("Enter Your Name :")
	fmt.Scan(&name)
	return name
}
func getTwoNum() (int, int) {
	var num1 int
	var num2 int
	fmt.Println("Enter the first Number")
	fmt.Scan(&num1)
	fmt.Println("Enter the Second Number")
	fmt.Scan(&num2)
	return num1, num2
}
func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}
func displayResult(name string, sum int) {
	fmt.Println("Your Name is ", name)
	fmt.Println("The sum is ", sum)
}

func main() {
	printWelcomeMessage()
	name := getUserName()
	num1, num2 := getTwoNum()
	sum := add(num1, num2)
	displayResult(name, sum)

}
