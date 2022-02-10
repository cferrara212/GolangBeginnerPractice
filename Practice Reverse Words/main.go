package main

import (
	"fmt"
	"strings"
)

func reverse_string (s string) string{
	chars := []rune (s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func reverse_words (s string) string {
	words := strings.Fields (s) 
	for i, j := 0, len(words)-1; i< j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main(){
	fmt.Println(reverse_words("brown cat"))
	fmt.Printf("%v\n", reverse_string("rad rad rad rad"))
}