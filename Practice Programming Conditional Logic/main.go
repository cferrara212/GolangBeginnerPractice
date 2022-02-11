package main

import (
	"fmt"
)

func main(){

	// basic if statement syntax
	theAnswer := 42
	var result string
	
	if theAnswer < 0 {
		result = "less than 0"
	} else if theAnswer == 0{
		result = "equal to 0"
	} else{
		result = "greater than 0"
	}

	//declare a variable for use in if statement
	fmt.Println(result)

	if x := -42; x <0 {
		result = "less than 0"
	} else if x == 0 {
		result = "equal to 0"
	} else {
		result = "greater than 0"
	}
	fmt.Println(result)
}