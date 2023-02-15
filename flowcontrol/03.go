package main

import "fmt"

func main() {
	sum := 1
	// While loop is a for loop
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

