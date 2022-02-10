package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	fmt.Println("the time now is: ", n)

	t := time.Date(2009, time.November, 10, 21, 0, 0, 0, time.UTC)
	fmt.Println("Go launched at ", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 21:00:00 2009")
	fmt.Printf("type of pased time is %T", parsedTime)
}
