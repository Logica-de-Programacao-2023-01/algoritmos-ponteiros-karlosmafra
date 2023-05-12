package main

import "fmt"

func main() {
	var x int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&x)
	var ptr *int = &x

	soma(ptr)

	fmt.Println(x)
}

func soma(ptr *int) {

}
