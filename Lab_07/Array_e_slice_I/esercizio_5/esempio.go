package main

import "fmt"

func main() {
	var a [6]int

	//fmt.Println(a)
	
	for i := range a {
		a[i] = i
	}

	for _, v := range a {
		v *= 2
	}

	fmt.Println(a)
}
