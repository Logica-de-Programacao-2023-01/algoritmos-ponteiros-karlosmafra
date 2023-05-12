package main

import "fmt"

func main() {
	x := 10
	var ptr *int = &x

	var n int
	fmt.Print("Digite um número natural: ")
	fmt.Scan(&n)

	atualizar(ptr, n)

	fmt.Println("O valor de x é", x)

}

func atualizar(ptr *int, n int) {
	var s int
	for i := 1; i <= n; i++ {
		s += i
	}
	*ptr = s
}
