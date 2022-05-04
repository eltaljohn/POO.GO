package main

import "fmt"

/*type course struct {
	name string
}

func (c course) Print() {
	fmt.Printf("%+v\n", c)
}
*/
// Declaración de Alias
// type myAlias = course

// Definición de tipo
// type newCourse course

type newBool bool

func (b newBool) String() string {
	if b {
		return "VERDADERO"
	}
	return "FALSO"
}

func main() {
	/*c := course{"Go"}
	c.Print()
	fmt.Printf("El Tipo es: %T", c)*/
	var b newBool = false
	fmt.Printf(b.String())
}
