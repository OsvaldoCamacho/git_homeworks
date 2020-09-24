package func_circles

import (
	"fmt"
	"math"
)

func func_circles(radio float64) {
	const pi float64 = 3.1416

	res := (pi * math.Pow(radio, 2))
	fmt.Println("El area del circulo es: ", res)

}
