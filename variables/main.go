package main

import "fmt"

func main() {
	aa := 12
	bb := "abc"
	cc := 3.14
	dd := true

	var a int
	var b string
	var c float64
	var d bool

	var a_long int = 1
	var b_long string = "cba"
	var c_long float64 = 3.15
	var d_long bool = false

	fmt.Print(a, b, c, d)
	fmt.Print("\n")
	fmt.Print(aa, bb, cc, dd)
	fmt.Print("\n")
	fmt.Print(a_long, b_long, c_long, d_long)
}
