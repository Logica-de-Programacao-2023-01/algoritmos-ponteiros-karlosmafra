package main

import "fmt"

type Conta struct {
	titular string
	saldo   float64
}

func main() {
	c := Conta{
		titular: "Carlos",
		saldo:   100,
	}

	var ptr *Conta = &c
	adcSaldo(ptr)
	fmt.Println(c)
}

func adcSaldo(ptr *Conta) {
	var adc float64
	fmt.Print("Valor para adicionar: ")
	fmt.Scan(&adc)

	ptr.saldo += adc
}
