package main
import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')  //char in single quotes
	fmt.Println("you entered: ", input)
}