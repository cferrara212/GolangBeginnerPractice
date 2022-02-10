package main
import(
	"fmt"
	"sort"
)

func main(){
	// by removing the specified amount inside array brackets [] array becomes a slice.
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	colors = append(colors, "Purple")
	fmt.Println(colors)

	//1 signifies the number to start on, and len(colors) says go to end of slice len
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	// append removal defaults to zero is no number is before the : so this says keep all starting at zero except for len -1
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	// a slice can also be declared with initial type and size with make()

	numbers := make([]int, 5, 5)
	numbers[0] = 134
	numbers[1] = 72
	numbers[2] = 123
	numbers[3] =42
	numbers[4] =32
	fmt.Println(numbers)

	numbers = append(numbers, 43)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)

}