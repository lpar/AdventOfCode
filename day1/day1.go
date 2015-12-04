package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Process processes a set of instructions, and returns the resulting floor
// number and the instruction position which causes Santa to enter the basement
func Process(instructions string) (floor int, basement int) {
	// Count the number of instructions we process, for part 2
	instnum := 0
	for _, c := range instructions {
		switch c {
		case '(':
			floor++
			instnum++
		case ')':
			floor--
			instnum++
		}
		if floor == -1 && basement == 0 {
			basement = instnum
		}
	}
	return
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Missing instructions file")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		os.Exit(2)
	}

	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(3)
	}

	floor, instnum := Process(string(bytes))
	fmt.Printf("Floor %d, instruction %d\n", floor, instnum)
}
