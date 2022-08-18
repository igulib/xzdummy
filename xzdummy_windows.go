//go:build windows
// +build windows

package xzdummy

import "fmt"

func SayHello() {
	fmt.Println("Hello from WINDOWS version of XZDummy package v0.1.3!")
}
