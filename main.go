package main

import "fmt"

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	fname := "Todd McLeod"
	fmt.Println("Hello world. I am", fname)
	fmt.Printf("%T\n", y)
	fmt.Println(y)
}
