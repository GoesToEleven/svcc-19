package main

import (
	"fmt"

	"github.com/GoesToEleven/svcc-19/mymath"
)

func main() {
	fmt.Println("welcome to main")
	z := mymath.Foo(4, 6)
	fmt.Println(z)
}
