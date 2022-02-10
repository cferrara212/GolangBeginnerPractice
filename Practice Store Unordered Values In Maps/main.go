package main

import(
	"fmt"
	"sort"
)
// these key value pairs are unordered lists, the output is not guarenteed to be read in the same order every time.
func main(){
	states := make(map[string]string)
	fmt.Println(states)
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)
	delete(states, "OR")
	states["NY"] = "New York"
	fmt.Println(states)

	//loop
	for k, v := range states {
		fmt.Printf("%v, %v\n", k, v)
	}

	//create a slice for order

	keys := make([]string, len(states))
	i:= 0
	for k := range states{
		keys[i] =k 
		i++
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}