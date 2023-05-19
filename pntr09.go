package main

import "fmt"

type Livros struct {
	titulo string
	autor  string
	preco  float64
}

func main() {
	l := Livros{
		titulo: "20 Mil Léguas Submarinas",
		autor:  "Júlio Verne",
		preco:  50.5,
	}

	var ptr *Livros = &l
	desconto(ptr)
	fmt.Println(l)
}

func desconto(ptr *Livros) {

}
