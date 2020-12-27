package main

import (
	"fmt"
)

func copyArray() {
	slice := []int{1, 2, 3, 4, 3, 5}
	copia := make([]int, len(slice), cap(slice)*2) //incrementar la capacidad x 2

	//copy(destino, fuente)
	copy(copia, slice)

	fmt.Println(slice)
	fmt.Println(copia)
}
