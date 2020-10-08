package main

import (
    "fmt"
)

//package level vars

var x int  // before name, after type
var y string 
var z bool 

func main(){

    fmt.Printf("%v\n%v\n%v",x,y,z)
    
}