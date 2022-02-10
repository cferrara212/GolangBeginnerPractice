package main

import (
	"fmt"

)

func main(){
	anInt := 42
	var p = &anInt
	fmt.Println("value of p:", *p)

	value1 := 42.13
	pointer1 := &value1
	fmt.Println("value 1:", *pointer1)

	*pointer1 = *pointer1/31
	fmt.Println("value pointer 1", *pointer1)
	fmt.Println("value 1: ", value1)
}