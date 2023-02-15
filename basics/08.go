package main

import "fmt"

// var used on package level
// if no initializer is present the variable will initialize with equivalent to null/zero
var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
