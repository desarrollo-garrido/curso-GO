package main

import "fmt"

func main() {
	var mapVar = make(map[uint]string)
	mapVar[101] = "Cervezas"
	mapVar[102] = "Almejas"
	mapVar[108] = "Coquinas"
	mapVar[105] = "Huevos de choco"

	fmt.Println("las tapitas de hoy", mapVar)
}
