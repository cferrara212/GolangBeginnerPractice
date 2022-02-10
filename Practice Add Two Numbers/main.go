package main
import "fmt"

func main(){
	numberOne := 1
	numberTwo := 2
	sum := numberOne + numberTwo
	fmt.Println("the sum of", numberOne, "plus", numberTwo, "is", sum)
	fmt.Printf("the variable type is %T\n", sum)

	//explicit type

	var number int = 1
	fmt.Printf("the type is %T", number)

	//implicit

	numberimplicit := 3
    fmt.Printf("the type is %T", numberimplicit)
}