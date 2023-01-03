package main

import (
	"fmt" // funciones sobre string
	"unsafe"
)

func main() {
	var myIntVar int = 22
	myIntVar = -12
	fmt.Println("El valor es ", myIntVar)

	var myUintVar uint
	myUintVar = 32
	fmt.Println("El valor es ", myUintVar)

	var myStringVar string
	myStringVar = "Me llamo Antonio Garrido"
	fmt.Println("El valor es ", myStringVar)

	var myBoolVar bool
	myBoolVar = true
	fmt.Println("El valor es ", myBoolVar)

	fmt.Println("Direccion de memoria es:", &myStringVar)

	myIntVar2 := -312
	fmt.Println("El valor es ", myIntVar2)

	myString2 := "Otro String de los cojones"
	fmt.Println("El valor es ", myString2)

	/* Constantes */

	myConst1 := 1
	fmt.Println("la constante es", myConst1)

	const myConst2 int = 2
	fmt.Println("Valor cd constante2", myConst2)

	const (
		TipoFuente   = "Times New Roman"
		AlturaFuente = 12
		Subrayado    = false
		Negrita      = true
	)

	fmt.Println("El valor de Tipo de Fuente es", TipoFuente)

	var numeroGrande int = 1_419_000_000_000
	fmt.Println("Valor grande", numeroGrande)

	var my8bit int8
	fmt.Printf("Valor de int8: %d\t\n", my8bit)

	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my8bit, unsafe.Sizeof(my8bit), unsafe.Sizeof(my8bit)*8)

	var myVarInt16 uint
	fmt.Printf("type: %T, bytes:%d\n", myVarInt16, unsafe.Sizeof(myVarInt16))

	var myVarInt int
	fmt.Printf("type: %T, bytes:%d\n", myVarInt, unsafe.Sizeof(myVarInt))

	var myFloat32 float32
	fmt.Printf("Default value %f\n", myFloat32)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myFloat32, unsafe.Sizeof(myFloat32), unsafe.Sizeof(myFloat32)*8)

	var myFloat64 float64
	fmt.Printf("Default value %f\n", myFloat64)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myFloat64, unsafe.Sizeof(myFloat64), unsafe.Sizeof(myFloat64)*8)

	fmt.Println()

	var myString3 string
	fmt.Printf("Default value to String %s", myString3)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myString3, unsafe.Sizeof(myString3), unsafe.Sizeof(myString3)*8)

	myLargeString := `En un lugar de la mancha,
	de cuyo nombre no quiero acordarme,
	nacio un caballero de alta alcurnia,
	su puta madre.
	`

	fmt.Printf(myLargeString)

	// Scope
	{
		// fmt.Println("*****************************")
		// floatVar := 33.11
		// fmt.Printf("type: %T, value: %f\n", floatVar, floatVar)
		// // Conversion a tipo.
		// floatStrVar := fmt.Sprintf("%.2f", floatVar)
		// fmt.Printf("type: %T, value: %s\n", floatStrVar, floatStrVar)

		// intVar := 22
		// fmt.Printf("type: %T, value: %d, memory: %d\n", intVar, intVar, &intVar)
		// intStrVar := fmt.Sprintf("%d", intVar)
		// fmt.Printf("type: %T, value: %s\n", intStrVar, intStrVar)

		// intVal1, err := strconv.ParseInt("1233", 0, 64)
		// fmt.Println(err)
		// fmt.Printf("type: %T, value: %d\n", intVal1, intVal1)

		// floatVal1, _ := strconv.ParseFloat("11.22", 64)
		// fmt.Printf("type: %T, value: %f\n", floatVal1, floatVal1)

	}

	{
		// Byte.
		const A byte = 'A'
		const a byte = 'a'
		fmt.Println("*****************")
		fmt.Println(A)
		fmt.Println(a)
		fmt.Println("*****************")
		const R byte = 82
		const t byte = 116
		fmt.Println(string(R))
		fmt.Println(string(t))

		array1 := []byte{106, 107, 108}
		fmt.Println(array1)
		fmt.Println(string(array1))

	}

}
