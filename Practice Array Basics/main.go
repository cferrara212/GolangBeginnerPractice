package main
import (
	"fmt"
)

func main(){
	// arrays are for storing data in Go. You cant easily sort them, or add and remove items at run items. package 
	//in slices for that.
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors)
	fmt.Println(colors[0])

	var numbers = [5]int{5,3,1,2,4}
	fmt.Println("number of numbers", len(numbers))
	fmt.Println("number of colors", len(colors))
}
