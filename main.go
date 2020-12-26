package main

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

	// css := Course{Name: "css desde cero", IsFree: true}

	// js := Course{}
	// js.Name = "Curso de JavaScript"
	// js.UserIDs = []uint{13, 19}

	//fmt.Println(Go.Name)
	//fmt.Printf("%+v\n", css)
	//fmt.Printf("%+v\n", js)

	//ya no es una funcion global como esta: PrintClasses(Go)
	//es una funcion, sino un metodo de Course
	Go.PrintClasses()
	Go.ChangePrice(13.43)
	//fmt.Println(Go.Price)
	//PrintTypesStringToInt()
	//PrintTypesIntToString()
	//justPrint()
	//printInput()
	//PrintBufio()
	whichNumberIsBigger()
}
