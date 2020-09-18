package main

import "fmt"

func main() {
	var faren float64
	fmt.Println("Ingresa los grados Fahrenheit a convertir: ")
	fmt.Scan(&faren)

	res := (faren - 32) * 5 / 9
	fmt.Println("Los grados Fahrenheit a Celsius son : ", res)

}
