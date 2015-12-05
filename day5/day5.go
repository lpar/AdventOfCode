package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/mathpl/golang-pkg-pcre/src/pkg/pcre"
)

const debug = false

// The problem specifically says that only the unaccented vowels are allowed
const Vowels = "aeiou"

var Prohibited = []string{"ab", "cd", "pq", "xy"}

// Nice1 implements the naughty-or-nice algorithm for part 1 of day 5
func Nice1(some string) bool {
	// Count pairs and vowels in a single linear pass
	vowels := 0
	doubles := 0
	lc := '\000'
	for _, c := range some {
		if strings.ContainsRune(Vowels, c) {
			vowels++
		}
		if c == lc && unicode.IsLetter(c) {
			doubles++
		}
		lc = c
	}
	if vowels < 3 || doubles < 1 {
		return false
	}
	// Check for the prohibited strings
	for _, bad := range Prohibited {
		if strings.Contains(some, bad) {
			return false
		}
	}
	return true
}

// Precompile the regular expressions for best performance
var Nice1Pair = pcre.MustCompile(`(*UTF8)([a-z])\1`, 0)
var Nice1Vowels = pcre.MustCompile(`(*UTF8)(?:.*[aeiou]){3}`, 0)

// Nice1RegExp also implements the naughty-or-nice algorithm for part 1 of
// day 5, but uses Perl Compatible Regular Expressions (PCRE)
func Nice1RegExp(some string) bool {
	for _, bad := range Prohibited {
		if strings.Contains(some, bad) {
			return false
		}
	}
	return Nice1Pair.MatcherString(some, 0).Matches() &&
		Nice1Vowels.MatcherString(some, 0).Matches()
}

func Nice1RegExpB(some string) bool {
	if !Nice1Pair.MatcherString(some, 0).Matches() ||
		!Nice1Vowels.MatcherString(some, 0).Matches() {
		return false
	}
	for _, bad := range Prohibited {
		if strings.Contains(some, bad) {
			return false
		}
	}
	return true
}

// Nice2 implements the naughty-or-nice algorithm for part 2 of day 5
func Nice2(some string) bool {
	gotpair := false
	gottriplet := false
	c2 := '\000'
	c3 := '\000'
	for i, c1 := range some {

		// Got two adjacent letters?
		if unicode.IsLetter(c1) {
			// If so, we can look for a pair that's also repeated somewhere else
			// but not overlapping
			if unicode.IsLetter(c2) {
				// Only pair-check if we need to
				if !gotpair {
					// Assemble our pair as a string
					pair := string([]rune{c2, c1})
					// Trick: Find the first and last occurrences and see if they are
					// non-overlapping. There may be others in the middle, but that
					// doesn't matter.
					i1 := strings.Index(some, pair)
					i2 := strings.LastIndex(some, pair)
					if i1 >= 0 && i2 >= 0 && (i2-i1) > 1 {
						gotpair = true
						if debug {
							fmt.Printf("Pair %s at [%d] and [%d] in %s\n", pair, i1, i2, some)
						}
					}
				}

				// Can only triplet detect if we have three adjacent letters
				if !gottriplet && unicode.IsLetter(c3) {
					// Triplet detection
					if c1 == c3 {
						gottriplet = true
						if debug {
							fmt.Printf("Triplet %c%c%c at [%d] in %s\n", c1, c2, c3, i, some)
						}
					}
				}
			}
		}

		// Exit as soon as we fulfil both conditions
		if gotpair && gottriplet {
			return true
		}

		c3 = c2
		c2 = c1
	}
	return false
}

// Again, precompile the regular expressions
var Nice2Pair = pcre.MustCompile(`(*UTF8)([\p{L}]{2}).*\1`, 0)
var Nice2Triple = pcre.MustCompile(`(*UTF8)([\p{L}])[\p{L}]\1`, 0)

// Nice2RegExp is a PCRE implementation of Nice2
func Nice2RegExp(some string) bool {
	return Nice2Pair.MatcherString(some, 0).Matches() &&
		Nice2Triple.MatcherString(some, 0).Matches()
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
