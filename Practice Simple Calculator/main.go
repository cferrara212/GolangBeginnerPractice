package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter a number: ")
	num1Input, _ := reader.ReadString('\n')
	fmt.Print("Please enter another number: ")
	num2Input, _ := reader.ReadString('\n')
	num1Parsed, err := strconv.ParseFloat(strings.TrimSpace(num1Input), 64)
	if err != nil {
		fmt.Println(err)
	}
	num2Parsed, err := strconv.ParseFloat(strings.TrimSpace(num2Input), 64)
	if err != nil {
		fmt.Println(err)
	}
	sumOfInputNumbers := num1Parsed + num2Parsed
	sumOfInputNumbers = math.Round(sumOfInputNumbers*100) / 100
	fmt.Printf("The sum of %v and %v is %v", num1Parsed, num2Parsed, sumOfInputNumbers)
	

}
