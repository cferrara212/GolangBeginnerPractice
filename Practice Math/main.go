package main

import (
	"fmt"
	"math"
)

func main() {
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("intsumis", intSum)
	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("floatsumis", floatSum)

	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("Rounded float sum: ", floatSum)

	cirRadius := 15.5
	circumference := cirRadius * 2 * math.Pi
	fmt.Printf("circumference %.2f\n", circumference)
}
