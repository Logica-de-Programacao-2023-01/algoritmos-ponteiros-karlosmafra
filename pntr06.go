package main

import "fmt"

type Livro struct {
	titulo string
	autor  string
}

func main() {
	l := Livro{
		titulo: "Nome",
		autor:  "Anônimo",
	}

	var ptr *Livro = &l
	autor(ptr)
	fmt.Println(l)
}

func autor(ptr *Livro) {
	if ptr.autor == "Anônimo" {
		ptr.titulo = "Desconhecido"
	}
}