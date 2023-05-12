package main

import "fmt"

func main() {
	var x int
	fmt.Print("Digite um valor inteiro: ")
	fmt.Scan(&x)

	var ptr *int = &x

	parimp(ptr)

	fmt.Println("X =", x)
}

func parimp(ptr *int) {
	if *ptr%2 == 0 {
		*ptr = 0
	} else {
		*ptr = 1
	}
}