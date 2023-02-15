package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const seed int64 = 123
	rand.Seed(seed)
	fmt.Println("Current seed:", seed)
	fmt.Println("Random generated number:", rand.Intn(9999))
}
