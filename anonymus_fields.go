package main

import (
	"fmt"
)

type Human struct {
	name string
}

func (this Human) hablar() string {
	return "hoooola"
}

type Tutor struct {
	Human
}

func crearTutor() {
	tutor := Tutor{Human{"tomas"}}

	fmt.Println(tutor.name)
}
