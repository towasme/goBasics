package main

import (
	"fmt"
)

func createSlice() {
	matriz := []int{1, 2, 3, 4} // no se define el tamaño
	var newMatriz []int         // cuando n ose declara nada en slice, coloca nil en todo, nil es nulo
	newMatriz = append(newMatriz, 3)
	fmt.Print(matriz)
	fmt.Print(newMatriz)

	// puntero al arreglo
	// longitud del arreglo al q apunta
	// capacidad

	//operacion slicing
	//arreglo := [3]int{1, 2, 3}
	//slice := arreglo[0:]
	//fmt.Println(slice)
	//devuelve {1,2}

}

func sliceCapacity() {
	newSlice := make([]int, 3, 5) //tipo, longitud, capacidad(opcional)
	newSlice = append(newSlice, 4)
	fmt.Println(cap(newSlice)) // devuelve capacidad del slice
}
