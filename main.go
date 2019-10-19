package main

import "fmt"

func main() {
	x := 42
	foo(&x)
	fmt.Println("back in main", x)
}

func foo(n *int) {
	*n = 43
	fmt.Println("inside foo", n)
}
