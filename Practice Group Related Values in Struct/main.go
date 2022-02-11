package main

import (
	"fmt"
)

func main(){
	// make poodle, and instanciate it as a Dog with Breed and Weight fields
	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)
	// use %+v to give the 
	fmt.Printf("%+v\n", poodle)
	// use string interpoloation with a placeholder %v in order to access specific fields of poodle Dog struct
	fmt.Printf("Breed %v\nWeight %v\n", poodle.Breed, poodle.Weight)
	//access and change value of struct
	poodle.Weight = 9
	fmt.Printf("Breed %v\nWeight %v\n", poodle.Breed, poodle.Weight)

}

//dog is a struct
type Dog struct {
	Breed string
	Weight int
}