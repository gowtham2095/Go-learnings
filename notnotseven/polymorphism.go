package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, "says Good Morning James!")
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, "says Shaken not stirred!")
}

func saySomething(h human) {
	h.speak()
}

type human interface {
	speak()
}

func main() {
	p := person{
		"Hari",
		"Haran",
	}

	sa := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}

	saySomething(p)
	saySomething(sa)

	sa.person.speak()
	// p.speak()
	// sa.speak()
	// fmt.Println(p.speak())
	// fmt.Println(sa.speak())
}
