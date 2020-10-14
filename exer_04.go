package main

import (
    "fmt"
)

type giuli int
var x giuli

func main(){
    
    fmt.Printf("%v \n",x)
    fmt.Printf("%T \n",x)
    x = 42
    fmt.Println(x)

}