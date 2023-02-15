package main

import "fmt"

// A return withou value returns the named return variables
// This is know as naked return
func split(sum int) (x, y int) {
	x = 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
