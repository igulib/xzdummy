//go:build !windows
// +build !windows

package xzdummy

import (
	"fmt"

	"github.com/igulib/xzdummy/pkg/one"
)

// PackageOneHello is a re-exported function from the internal package
// github.com/igulib/xzdummy/pkg/one
var PackageOneHello = one.PackageOneHello

func SayHello() {
	fmt.Println("Hello from UNIX version of XZDummy package v0.1.5!")
}
