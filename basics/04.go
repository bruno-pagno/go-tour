package main

import "fmt"

// Type comes after variable name
func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(123, 456))
}
