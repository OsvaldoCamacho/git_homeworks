package main

import (
	"fmt"
	"math"
)

func main() {
	var radio float64
	const pi float64 = 3.1416
	fmt.Println("Ingresa el radio del circulo: ")
	fmt.Scanf("%f", &radio)

	res := (pi * math.Pow(radio, 2))
	fmt.Println("El area del circulo es: ", res)

}
