package chapter1

import "fmt"

var c, java, rust, python, haskell bool

var i, j int = 20, 56

func GoVariables() {
	var index int
	fmt.Printf("Languages %v %v %v %v %v %v\n", index, c, java, rust, python, haskell)

	var dog, lion, sheep = "dog", "Lion", "Sheep"

	fmt.Println(i, j, dog, lion, sheep)

	energyType := "Kinetic Energy"
	fmt.Printf("Energy Type %v\n", energyType)

	for _, names := range "OgdenMorrow" {
		fmt.Printf("Names -> %d %c\n", names, names)
	}

	studentGPA := 45.66790
	fmt.Printf("Student GPA : %1.2f\n", studentGPA)

	var a int
	var b float64
	var c bool
	var d string
	fmt.Printf("%v %v %v %v \n", a, b, c, d)
}
