package main

import (
	"fmt"
	"strconv"
) 

func PrintTypesIntToString() {
	edad := 22

	edad_str := strconv.Itoa(edad)

	fmt.Println("Mi edad es " + edad_str)
}

func PrintTypesStringToInt() {
	NuevaEdad := "10"

	edad_int,_ := strconv.Atoi(NuevaEdad)

	fmt.Println(edad_int + 10)
}