//go:generate echo 'go package 2'
package main

type Astruct struct {
	Name string
}

func (a Astruct) MyName() string {
	return a.Name
}

func NewA(n string) Astruct {
	a := Astruct{n}
	return a
}
