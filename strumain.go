//go:generate echo 'Example of using package and vendor in GO version 1.5'
package main

func main() {
	println("Here is go package main.")

	SaySomething("I am using package1!")

	a := NewA("Struct from package2")
	println(a.MyName())
}
