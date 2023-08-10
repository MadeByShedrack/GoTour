package chapter1

import "fmt"

func GoFunctions() {

	x := 42
	y := 13

	result := addTwo(x, y)

	fmt.Printf("Result -> %v\n", result)

	first, second := swapStrings("Hello", "World")
	fmt.Printf("Swapped Strings -> %v %v\n", first, second)

	fmt.Println(split(25))
}

func addTwo(x, y int) int {
	return x + y
}

func swapStrings(first, second string) (string, string) {
	return first, second
}

func split(sum int) (num1, num2 int) {
	num1 = sum * 4 / 9
	num2 = sum - num1
	return
}
