package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const debug = false

const Vowels = "aeiou"

var Prohibited = []string{"ab", "cd", "pq", "xy"}

func Nice1(some string) bool {
	// Check for the prohibited strings first
	for _, bad := range Prohibited {
		if strings.Contains(some, bad) {
			return false
		}
	}
	// If we pass that, count pairs and vowels
	vowels := 0
	doubles := 0
	lc := '\000'
	for _, c := range some {
		if strings.ContainsRune(Vowels, c) {
			vowels++
		}
		if c == lc {
			doubles++
		}
		lc = c
	}
	return vowels >= 3 && doubles >= 1
}

func Nice2(some string) bool {
	// See if a pair occurs in two non-overlapping places.
	gotpair := false
	for i, _ := range some {
		if i < len(some)-1 {
			// Sneaky trick: Only find the first and last occurrences and check
			// those. If there are even more occurrences between them, we don't care.
			pair := some[i : i+2]
			i1 := strings.Index(some, pair)
			i2 := strings.LastIndex(some, pair)
			if i1 >= 0 && i2 >= 0 && (i2-i1) > 1 {
				gotpair = true
				if debug {
					fmt.Printf("Pair %s at [%d] and [%d] in %s\n", pair, i1, i2, some)
				}
				break
			}
		}
	}
	if !gotpair {
		return false
	}
	// Now check for a triplet -- two character the same with exactly one
	// between them
	c2 := '\000'
	c3 := '\000'
	for i, c1 := range some {
		if i < len(some) {
			if c1 == c3 {
				if debug {
					fmt.Printf("Triplet %c%c%c at [%d] in %s\n", c1, c2, c3, i, some)
				}
				return true
			}
			c3 = c2
			c2 = c1
		}
	}
	return false
}

func Process(fname string, round int) (nicecount int) {
	inf, err := os.Open(fname)
	if err != nil {
		fmt.Println("Cannot open input file")
		os.Exit(4)
	}
	defer inf.Close()

	// Work out which Nice function to use
	nice := Nice1
	if round == 2 {
		nice = Nice2
	}

	scanner := bufio.NewScanner(inf)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		input := scanner.Text()
		if nice(input) {
			fmt.Printf("%s nice\n", input)
			nicecount++
		} else {
			fmt.Printf("%s naughty\n", input)
		}
	}
	return nicecount
}

func main() {

	Nice2("abcdefg")

	if len(os.Args) < 2 {
		fmt.Println("Missing input file")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		os.Exit(2)
	}

	for i := 1; i <= 2; i++ {
		fmt.Printf("Round %d:\n  Total nice = %d\n", i, Process(os.Args[1], i))
	}

}
