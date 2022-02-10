package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Number: ")
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {
		fmt.Println(err)
	} else{
		fmt.Println("value of number: ", aFloat)
		fmt.Printf("type is %T: ", aFloat)
	}
}