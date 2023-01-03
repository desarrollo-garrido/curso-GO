package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("**** Arrays *****")
	println()

	var myIntVar int
	fmt.Printf("Type; %T, bytes: %d, bits: %d\n", myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar)*8)
	var myArrayVar [4]int
	myArrayVar[0] = 22
	myArrayVar[1] = 91
	myArrayVar[2] = 31
	myArrayVar[3] = 19

	fmt.Printf("El valor del array es %d", myArrayVar)
	println()

	// Otra definicion
	myArrayString := [3]string{"Pera", "Manzana", "Sandia"}
	fmt.Println(myArrayString)
	fmt.Println("Size of array ", len(myArrayString)) // len = length
	println()
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myArrayString, unsafe.Sizeof(myArrayString), unsafe.Sizeof(myArrayString)*8)

	var otherArray [3]string
	otherArray[0] = "Antonio Garrido"
	otherArray[1] = "Manuel Gama Estudillo"
	otherArray[2] = "Ernesto Mora Gonzalez"
	fmt.Println(otherArray)
	fmt.Printf("La longitud del array es de %v\n", len(otherArray))

	carBrand := [4]string{"SEAT", "VOLKSWAGEN", "AUDI", "SKODA"}
	fmt.Println(carBrand)

	carBrand2 := [...]string{
		"Toyota",
		"Kia",
		"Lexus",
		"Hyundai",
		"Chrysler",
		"Dodge",
		"Chevrolet",
	}
	fmt.Println(carBrand2)

	numeros := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for i := 0; i < len(numeros); i++ {
		fmt.Printf("Valor de numero, %v, indice %d\n", numeros[i], i)
	}

	for _, v := range numeros {
		fmt.Printf("Valor: %v\n", v)
	}
}
