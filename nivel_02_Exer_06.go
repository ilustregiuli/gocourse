package main

import (
	"fmt"
)

//Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos
	
const(

	x = 2021 + iota
	y 
	z 
	w 
	
)

func main() {
	
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(w)
		
	
	
}
