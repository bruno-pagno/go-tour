package main

import "fmt"

func main() {
	// Init and post statements are optional
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

