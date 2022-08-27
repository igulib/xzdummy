package one

import "fmt"

func positiveOrNegative(val int) {
	if val > 0 {
		fmt.Println("val is positive")
	} else if val == 0 {
		fmt.Println("val is zero")
	} else {
		fmt.Println("val is negative")
	}
}

func PackageOneHello() {
	fmt.Println("Hello from package one!")
	positiveOrNegative(5)
}
