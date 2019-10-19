package main

import "fmt"

type person struct {
	first string
	last  string
}

// func (r receiver) identifier(parameters) return(s) { code }

func (p person) speak() {
	fmt.Println("I AM A PERSON NAMED", p.first, p.last)
}

type secretAgent struct {
	person
	ltk bool
}

func (sa secretAgent) speak() {
	fmt.Println("I AM A PERSON NAMED", sa.first, sa.last)
}

type human interface {
	speak()
}

func foo(h human) {
	h.speak()
}

func main() {
	// composite literal
	// p1 := type{values}
	p1 := person{
		first: "Jenny",
		last:  "Moneypenny",
	}
	fmt.Println(p1)

	p2 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	fmt.Println(p2)

	foo(p1)
	foo(p2)

	/*
	in go we really talk about and think about
	a VALUE being of certain TYPES
	*/
}
