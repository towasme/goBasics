package main

import (
	"fmt"
)

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

func main() {
	Go := Course{
		Name:    "curso desde cero",
		Price:   30.9,
		IsFree:  false,
		UserIDs: []uint{10, 14, 17},
		Classes: map[uint]string{
			1: "Introduccion",
			2: "Estructuras",
			3: "Maps",
		},
	}

	fmt.Println(Go.Name)
}
