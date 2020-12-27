package main

import (
	"fmt"
)

//for normal
func helloWorldXTimes(times int) {
	for i := 0; i < times; i++ {
		fmt.Println("Hello World")
	}
}

//while hecho con for
func semiWhile(times int) {
	i := 0
	for i < times {
		fmt.Println("New Hello Wolrd")
		i++
	}

}

//while hecho con for
func semiWhileWithBreak(times int) {
	i := 0
	for {
		fmt.Println("New Hello Wolrd")
		i++
		if i > times {
			break
		} else if i == 2 {
			continue
		}

	}

}
