package main

import "fmt"

func main() {
	var lado_cuadrado float64
	fmt.Println("Ingresa el tamaño del lado del cuadrado: ")
	fmt.Scanf("%f", &lado_cuadrado)

	res := lado_cuadrado * lado_cuadrado
	fmt.Println("El area del cuadrado es: ", res)

}
