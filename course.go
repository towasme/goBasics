package main

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

func (c *Course) ChangePrice(price float64){
	c.Price = price
}

// metodos son funciones con un argumento receptor, s especifica entre () entre func y el nombre del metodo
//metodo de declara por fuera de la estructura, no por dentro
func (c Course) PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}