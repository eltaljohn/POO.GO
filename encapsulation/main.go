package main

import (
	"fmt"
	"github.com/eltaljohn/encapsultaion/course"
)

func main() {
	Go := course.New(
		"Go desde cero",
		0,
		false)
	Go.SetUserIDs([]uint{12, 56, 89})
	Go.SetClasses(map[uint]string{
		1: "Introducción",
		2: "Estructuras",
		3: "Maps",
	})

	Go.SetName("POO con Go")
	fmt.Println(Go.Name())
	Go.PrintClasses()
}
