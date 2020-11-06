package main

import (
	"fmt"
)

var diferente bool
var maior bool

//	Escreva expressões utilizando os operadores, e atribua seus valores a variáveis.
//	Demonstre estes valores.

func main() {
	
	x := 0
	y := 1
	if (x == y){
	  diferente = false	
	} else if (x != y) {
	  diferente = true	
	}
	
	if(x > y) {
	  maior = true
	} else if (x < y){
	  maior = false
	}
	
	fmt.Printf("X é diferente de Y: %v\n",diferente)
	fmt.Printf("X é Maior que Y: %v\n",maior)
	
	fmt.Println("--------------------------------------------")
	// solução da professora
	a := (10 == 100)
	b := (10 != 100)
	c := (10 <= 100)
	d := (10 < 100)
	e := (10 >= 100)
	f := (10 > 100)
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e, f)
	
	
	
}

