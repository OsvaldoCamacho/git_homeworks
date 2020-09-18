package main

import "fmt"

func main() {
	var lado float64
	fmt.Println("Ingresa el tamaño del lado del cuadrado: ")
	fmt.Scanf("%f", &lado)

	res := lado * lado
	fmt.Println("El area del cuadrado es: ", res)

}
