package main

import "fmt"

func main() {
	var base float64
	var altura float64

	fmt.Println("Ingresa la base del triangulo: ")
	fmt.Scan(&base)

	fmt.Println("Ingresa la altura del triangulo: ")
	fmt.Scan(&altura)

	res := (base * altura) / 2

	fmt.Println("El area del triangulo es: ", res)

}
