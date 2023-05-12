package main

import "fmt"

func main() {
	var s string
	fmt.Print("Escreva uma string: ")
	fmt.Scan(&s)

	var ptr *string = &s

	reveter(ptr)

	fmt.Println(s)
}

func reveter(ptr *string) {
	var reverse string
	str := *ptr
	for i := len(str) - 1; i >= 0; i-- {
		reverse += string(str[i])
	}
	*ptr = reverse
}
