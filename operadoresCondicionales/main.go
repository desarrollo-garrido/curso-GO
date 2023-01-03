package main

import "fmt"

func main() {
	yearsOld := 58
	fmt.Println("Operadores Condicionales")
	fmt.Print()
	fmt.Println(yearsOld > 18)

	if yearsOld < 18 {
		fmt.Printf("Lo siento, no eres mayor de edad. Tienes %d años", yearsOld)
	} else {
		fmt.Printf("Puedes pasar, tu edad es %d años (pureta de los cojones)", yearsOld)
	}
	fmt.Println()
	var boolVal bool = false
	if !boolVal {
		fmt.Println("Condicion true en negativa")
	} else {
		fmt.Println("Condicion true")
	}
	if value := fn(); value == true {
		fmt.Println("is true")
	}

	// Si es true, la igualdad no hace falta
	if value := fn(); value {
		fmt.Println("is true again !!Gilipollas!!")
	}
	var number int = 33
	if number == 1 {
		fmt.Println("Es 1")
	} else if number == 2 {
		fmt.Println("Es 2")
	} else if number == 3 {
		fmt.Println("es 3")
	} else {
		fmt.Println("Cualquiera sabe que puto numero has puesto !!!Gilipollas!!!")
	}

	switch valor := "fresa"; valor {
	case "fresa":
		fmt.Println("Es correcto, tienes 10 puntos, es ", valor)
	case "manzana":
		fmt.Println("Es correcto, tienes 5 puntos, es ", valor)
	default:
		fmt.Println("No es ninguna fresa / manzana, tienes 0 puntos ", valor)
	}

	numberOption := 21
	switch {
	case numberOption > 18:
		fmt.Println("Eres mayor de edad, puedes pasar 😘")
	case numberOption < 18:
		fmt.Println("Eres menor de edad 😒 adios")
	default:
		fmt.Println("No has puesto nada gilipollas 🤬")

	}
}

func fn() bool {
	return true
}
