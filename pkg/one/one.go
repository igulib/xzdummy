package one

import "fmt"

func isPositive(val int) {
	if val > 0 {
		fmt.Println("val is positive or zero")
	} else {
		fmt.Println("val is negative")
	}
}

func PackageOneHello() {
	fmt.Println("Hello from package one!")
	isPositive(5)
}
