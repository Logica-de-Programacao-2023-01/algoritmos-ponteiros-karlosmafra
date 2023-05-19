package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Print("Digite um número: ")
	fmt.Scan(&x)

	var ptr *float64 = &x

	media(ptr)

	fmt.Println(x)
}

func media(ptr *float64) {
	*ptr = (*ptr + math.Pi) / 2
}
