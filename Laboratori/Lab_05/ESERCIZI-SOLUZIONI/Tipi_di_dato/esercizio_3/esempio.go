package main

import "fmt"

func main() {
	var a int = 4
	var b float64 = 12.5
	var c int
	c = a + int(b)

	var d float64
	d = float64(a / c)

	fmt.Println(a, b, c, d)
}
