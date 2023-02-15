package main

import "fmt"

// You can omit the type if two or more are from the same type
func add(x, y, z int) int {
	return x + y + z
}

func foo(a, b int, c, d string) int {
	fmt.Println(c)
	fmt.Println(d)
	return a + b
}

func main() {
	fmt.Println(add(1, 2, 3))
	fmt.Println(foo(1, 2, "testing", "this"))
}
