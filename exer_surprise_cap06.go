package main

import (
	"fmt"
)
var letter string;

func main() {

	for i:=33; i <=122; i++ {
		letter = string(i)
		fmt.Printf("%d - %v \n",i,letter)
	}
}

