package main

import "fmt"

 var a = 10

func main()  {
	fmt.Println("Hello Init Function")
	fmt.Println(a)
}

func init()  {
	fmt.Println(a)
	fmt.Println("I am the first function that executed first")
	a = 20
	fmt.Println(a)
}