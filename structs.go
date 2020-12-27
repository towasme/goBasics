package main

import (
	"fmt"
)

type User struct {
	edad             int
	nombre, apellido string
}

func newStruct(edad int, nombre string, apellido string) {
	newUser := User{edad, nombre, apellido}

	fmt.Println(newUser)

	newUser.nombre = "cipriano"
	newUser.apellido = "mejia"

	fmt.Println(newUser)
}

func (usuario User) nombre_completo() string {
	return usuario.nombre + " " + usuario.apellido
}

func (this *User) set_name(n string) {
	this.nombre = n
}
