package main

import (
	"fmt"
	"math"
)

func main() {
	var radio float64

	fmt.Println("Ingresa el radio del circulo: ")
	fmt.Scanf("%f", &radio)

	res := (3.1416 * math.Pow(radio, 2))
	fmt.Println("El area del circulo es: ", res)

}
