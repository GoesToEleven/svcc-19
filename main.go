package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println("-----")
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", &x)
	fmt.Println("-----")
	fmt.Println(*&x)
}

func foo(n int) {
	n = 43
	fmt.Println("inside foo", n)
}
