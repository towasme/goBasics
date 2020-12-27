package main

import "fmt"

func pointerAllTheWay() {
	var x, y *int
	entero := 5

	x = &entero
	y = &entero

	fmt.Println(x)
	fmt.Println(y)
}
