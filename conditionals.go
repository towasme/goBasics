package main

import (
	"fmt"
)

func whichNumberIsBigger() {
	var x, y int
	fmt.Println("escribe dos numeros para comprobar el mayor\nEscribe el primer numero")
	fmt.Scanf("%d\n", &x)
	fmt.Println("Escribe el segundo numero")
	fmt.Scanf("%d\n", &y)
	if x < y {
		fmt.Printf("%d no es mayor que %d", x, y)
	} else if x == y {
		fmt.Printf("%d es igual a %d", x, y)
	} else {
		fmt.Printf("%d es mayor que %d", x, y)
	}
}
