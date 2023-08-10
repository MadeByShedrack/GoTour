package chapter1

import (
	"fmt"
	"math"
)

const ValueOfPi = 3.14

func GoTypeConversions() {

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	fmt.Printf("PI Value %T %v\n", ValueOfPi, ValueOfPi)

	const Truth = true
	fmt.Printf("Type => %T || Value => %v\n", Truth, Truth)
}
