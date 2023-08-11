package chapter2

import (
	"fmt"
	"runtime"
	"time"
)

func GoSwitch() {

	checkOS := runtime.GOOS

	switch checkOS {
	case "darwin":
		fmt.Printf("Go runs on %v\n", checkOS)
	case "linux":
		fmt.Printf("Go runs on %v\n", checkOS)
	default:
		fmt.Printf("%s\n", checkOS)
	}

	fmt.Println("When's Saturday")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away")
	}

	// for range time.Tick(time.Second * 1) {
	// 	fmt.Printf("Ticking in %v\n", time.Second)
	// }

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
