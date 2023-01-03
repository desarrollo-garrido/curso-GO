package main

import "fmt"

func main() {
	// Several variables at once.

	var name, firsName, lastName string
	name = "Antonio"
	firsName = "Garrido"
	lastName = "Carranza"
	fmt.Println(name, firsName, lastName)

	// Another way.
	brand, model, color := "Volkswagen", "T-Cross", "blue"
	fmt.Printf("The %v %v car is %v\n.", model, brand, color)

	// Constantes.

	const curso string = "Curso profesional de Go" // No se puede cambiar su valor, es inmutable.
	fmt.Printf("curso %v", curso)
}
