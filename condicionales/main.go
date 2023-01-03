package main

import (
	"fmt"
	"runtime"
)

func main() {
	arch := runtime.GOARCH
	fmt.Println(arch)

	switch arch {
	case "amd64":
		fmt.Println("Sistema operativo Windows")
	}
	switch {
	case arch == "amd64":
		fmt.Println("Sistema operativo Windows gilipollas")
	}

	// No se que hace esta linea
	println()
	diaSemana := "DomingoNOOO"
	switch diaSemana {
	case "Lunes", "Martes", "Miercoles", "Jueves", "Viernes":
		fmt.Println("Es un dia entre-semana", diaSemana)
	case "Sabado", "Domingo":
		fmt.Println("Es fin de semana", diaSemana)
	default:
		fmt.Println("Vete a chuparla, no es un dia de la semana", diaSemana)

	}
}
