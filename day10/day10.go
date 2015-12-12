package main

import (
	"bytes"
	"fmt"
)

func encode(some string) string {
	var result bytes.Buffer
	count := 0
	var curc rune
	for _, c := range some {
		// Hit a new character?
		if curc != c && curc != 0 && count != 0 {
			result.WriteRune(rune(48 + count))
			result.WriteRune(curc)
			count = 0
		}
		count++
		curc = c
	}
	// Fell off the end
	if curc != 0 && count != 0 {
		result.WriteRune(rune(48 + count))
		result.WriteRune(curc)
	}
	return result.String()
}

func main() {

	// Part 1
	//My input	s := "3113322113"
	s := "1321131112"
	for i := 0; i < 40; i++ {
		s = encode(s)
	}
	//	fmt.Println(s)
	fmt.Println(len(s))

	// Part 2
	for i := 0; i < 10; i++ {
		s = encode(s)
	}
	//	fmt.Println(s)
	fmt.Println(len(s))
}
