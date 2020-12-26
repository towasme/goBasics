package main

import (
	"bufio"
	"fmt"
	"os"
)

//imprimir con inputs de consola
func printInput() {
	var name string
	fmt.Println("Cual es tu nombre?")
	fmt.Scanf("%s\n", &name)
	fmt.Printf("Hola %s\n", name)
}

//imprimir normal datos normales
func justPrint() {
	edad := 24
	fmt.Printf("Mi edad es %d\n", edad)
}

// imprimir utilizando bufio. bufio retorna dos vlores el texto y error
func PrintBufio() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre: ")
	nombre, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Hola %s, esto se imprime con la libreria bufio", nombre)
	}
}
