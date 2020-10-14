package main

import (
    "fmt"
)

// package level scope
type giuli int
var x giuli
var y int

func main() {
    
    fmt.Printf("%v \n",x)
    fmt.Printf("%T \n",x)
    x = 42
    fmt.Println(x)

    // conversão do valor de "x" para o tipo subjacente
    y = int(x)


    fmt.Println(y)
    fmt.Printf("%T\n",y)
    fmt.Printf("%T",x)	

}
