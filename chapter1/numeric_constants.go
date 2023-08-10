package chapter1

import "fmt"

const (
	BigNumbers   = 1 << 100
	SmallNumbers = BigNumbers >> 99
)

func NumericConstants() {
	fmt.Println(needInt(SmallNumbers))
	fmt.Println(needFloat(SmallNumbers))
	fmt.Println(needFloat(BigNumbers))
}

func needInt(numbers int) int {
	return numbers*10 + 1
}

func needFloat(numbers float64) float64 {
	return numbers * 0.1
}
