package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const InstructionCodes = "^v<>"

const debug = false

type Coordinates struct {
	X, Y int
}

func (c Coordinates) String() string {
	return fmt.Sprintf("(%d,%d)", c.X, c.Y)
}

var _ fmt.Stringer = Coordinates{}

func Process(presents map[Coordinates]int, instructions string) {
	santa := Coordinates{0, 0}
	presents[santa]++
	if debug {
		fmt.Println("Delivery phase")
		fmt.Printf("  %s has %d presents\n", santa, presents[santa])
	}
	for _, c := range instructions {
		if strings.ContainsRune(InstructionCodes, c) {
			switch c {
			case '<':
				santa.X--
			case '>':
				santa.X++
			case 'v':
				santa.Y--
			case '^':
				santa.Y++
			}
			presents[santa]++
			if debug {
				fmt.Printf("  %s has %d presents\n", santa, presents[santa])
			}
		}
	}
}

func Report(presents map[Coordinates]int) int {
	result := 0
	if debug {
		fmt.Println("Totaling phase")
	}
	for coords, prescount := range presents {
		if debug {
			fmt.Printf("  %s has %d presents\n", coords, prescount)
		}
		result++
	}
	fmt.Printf("%d houses have at least one present\n\n", result)
	return result
}

func SplitInstructions(x string) (string, string) {
	isodd := true
	odd := make([]rune, 0, len(x)/2+1)
	even := make([]rune, 0, len(x)/2+1)
	for _, c := range x {
		if strings.ContainsRune(InstructionCodes, c) {
			if isodd {
				odd = append(odd, c)
			} else {
				even = append(even, c)
			}
			isodd = !isodd
		}
	}
	return string(odd[0:]), string(even[0:])
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

	instructions := string(bytes)

	fmt.Println("Part 1")
	presents := make(map[Coordinates]int)
	Process(presents, instructions)
	Report(presents)

	fmt.Println("Part 2")
	santa, robosanta := SplitInstructions(instructions)

	presents = make(map[Coordinates]int)
	Process(presents, santa)
	Process(presents, robosanta)
	Report(presents)

}
