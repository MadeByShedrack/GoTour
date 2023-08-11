package chapter2

import (
	"fmt"
	"math"
	"runtime"
)

func GoConditionals() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "windows":
		fmt.Println("Windows")
	case "linux":
		fmt.Println("Linux")
	case "freebsd":
		fmt.Println("BSD")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func sqrt(x float64) string {

	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {

	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}
