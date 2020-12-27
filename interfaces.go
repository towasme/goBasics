package main

type User interface {
	permissions() int // puede devolver del 1 - 5
	Nombre() string
}

type Admin struct {
	nombre string
}

func (this Admin) permissions() int {
	return 5
}

func (this Admin) Nombre() string {
	return this.name
}

func auth(user User) string {
	if user.permissions() > 4 {
		return user.Nombre() + "tiene persimos de administrador"
	}
	return "no es administrador"
}
