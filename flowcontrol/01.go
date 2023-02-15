package main

import "fmt"

func main() {
	sum := 0
	// No parenthesis at if statement
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}
