package main

import (
	"fmt"
)


//	Atribua um valor int a uma variável
//	Demonstre esse valor em binário e hexa
//	Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável		
//	Demonstre esta outra variável em decimal, binário e hexadecimal	
	

func main() {
	
	variavel := 70
	fmt.Printf("%d\n%b\n%#x\n", variavel, variavel, variavel)
	varDireita := variavel << 1
	fmt.Println("----------------------------------------------")
	fmt.Printf("%d\n%b\n%#x", varDireita, varDireita, varDireita)
		
	
	
}
