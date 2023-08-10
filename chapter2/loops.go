package chapter2

import "fmt"

func GoLoops() {

	sum := 0

	for index := 0; index < 10; index++ {
		sum += index
	}

	fmt.Printf("Sum -> %v\n", sum)
	iterations()
}

func iterations() {

	sum := 1

	for sum < 250 {

		sum += sum
	}

	fmt.Printf("Sum -> %v\n", sum)
}
