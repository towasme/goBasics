package main

import (
	"fmt"
)

func printArray() {
	var arreglo [10]string //manera de declarar arreglo con 0 o vacio para cada elemento
	array := [4]int{2, 5}  // tiene 4 pero si no se especifican coloca 0 en el ultimo

	for i := 0; i < len(arreglo); i++ {
		for j := 0; j < len(array); j++ {
			fmt.Print(array[j])
		}
		fmt.Println("")
	}
}
