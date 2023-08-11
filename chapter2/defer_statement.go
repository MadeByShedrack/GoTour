package chapter2

import "fmt"

func DeferFunction() {
	// defer fmt.Println("Hello, defer function")
	// fmt.Println("Hello world")

	fmt.Println("Counting")

	for index := 0; index < 10; index++ {
		defer fmt.Printf("Index %v\n", index)
	}

	fmt.Println("Done")

	defer withdrawFunds(2000)
	depositFunds(12000)
}

func withdrawFunds(amount int) int {

	if amount < 5000 {
		fmt.Println("You can only withdraw 5k and above")
	} else {
		fmt.Println("Trasactions successful")
	}

	return amount
}

func depositFunds(amount int) int {

	if amount > 10000 {
		fmt.Println("The money too much, upgrade your account")
	} else {
		fmt.Println("Trasactions successful")
	}

	return amount
}
