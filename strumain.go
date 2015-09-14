//go:generate echo 'Example of using package and vendor in GO version 1.5'
package main

import (
	// First, you have to open vendoring in GO1.5, use: `$ export GO15VENDOREXPERIMENT=1`
	// GO will find packages in "$PWD/vendor/" if "$GOPATH/src/" has no packages you want.
	// Then you can use package struven now!
	sv "github.com/iamken1204/struven"
)

func main() {
	println("Here is file0 in package main.")

	// You can directely call other file's method under the same package name.
	SaySomething("I am using file1's method!")

	a := NewA("Struct from file2")
	println(a.MyName())

	ven := sv.NewRandomVen()
	println(ven.ID)
}
