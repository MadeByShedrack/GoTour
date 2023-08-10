package chapter1

import (
	"fmt"
	"math"
	"math/rand"
)

func GoPackages() {
	favouriteNumber := rand.Intn(10)
	fmt.Printf("My favourite number is: %v\n", favouriteNumber)

	squareRoot := math.Sqrt(7)

	fmt.Printf("Square root: %v\n", squareRoot)

	fmt.Println(math.Pi)
}
