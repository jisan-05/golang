package main

import "fmt"

var a = 10
var b = 20

func add(x int, y int) {
	var z = x + y
	fmt.Println(z)
}

func main() {
	p := 5
	q := 6
	add(p,q)
	add(a,b)
	add(a,p)
	// add(b,z)
}
