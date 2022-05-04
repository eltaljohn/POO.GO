package main

import "fmt"

func main() {
	Go := Course{
		"Go desde cero",
		12.34,
		false,
		[]uint{12, 56, 89},
		map[uint]string{
			1: "Intraducción",
			2: "Estructuras",
			3: "Maps",
		},
	}

	Go.PrintClasses()
	Go.ChangePrice(12)
	fmt.Printf("%+v", Go)
}
