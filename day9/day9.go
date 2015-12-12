package main

/* Ugly solution using goroutines to collect answers, and strings for sets.
 * But it works, is fast, and collects both answers in one go.
 */

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Route struct {
	From string
	To   string
}

type Result struct {
	RouteTaken string
	Cost       int
}

var debug bool = false

// Channels to pass all possible routes and costs to filtering goroutines
var Results1 = make(chan Result)
var Results2 = make(chan Result)

// Channels to pass back the filtered answers
var Shortest = make(chan Result)
var Longest = make(chan Result)

// Fast storage for distance from A->B
var Costs map[Route]int

var InputLine = regexp.MustCompile(`^(\w+) to (\w+) = (\d+)$`)

// Read the input data, storing the possible routes and their distances in
// Costs, and return a string naming all the locations mentioned.
func ReadInput(filename string) string {
	inf, err := os.Open(filename)
	if err != nil {
		fmt.Println("Cannot open input file")
		os.Exit(4)
	}
	defer inf.Close()
	scanner := bufio.NewScanner(inf)
	scanner.Split(bufio.ScanLines)
	locationset := make(map[string]bool)
	Costs = make(map[Route]int)
	for scanner.Scan() {
		input := scanner.Text()
		m := InputLine.FindStringSubmatch(input)
		cost, err := strconv.Atoi(m[3])
		if err != nil {
			panic(err)
		}
		locationset[m[1]] = true
		locationset[m[2]] = true
		Costs[Route{m[1], m[2]}] = cost
	}
	var locset string
	for k := range locationset {
		if locset == "" {
			locset = locset + k
		} else {
			locset = locset + " " + k
		}
	}
	return locset
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing input file")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		os.Exit(2)
	}
	locations := ReadInput(os.Args[1])

	// Start goroutines to find the longest and shortest routes
	go FindShortest()
	go FindLongest()

	// Recursively compute all possible routes and their distances, and feed
	// that to both goroutines
	compute("", locations)

	// Signal that we've finished coming up with routes
	close(Results1)
	close(Results2)

	// Read back and display the answers the goroutines found
	answer1 := <-Shortest
	fmt.Printf("Least expensive route:\n%s = %d\n", answer1.RouteTaken, answer1.Cost)

	answer2 := <-Longest
	fmt.Printf("Most expensive route:\n%s = %d\n", answer2.RouteTaken, answer2.Cost)
}

func FindShortest() {
	mincost := 0
	minroute := ""
	for {
		result, ok := <-Results1
		if !ok {
			Shortest <- Result{minroute, mincost}
			return
		}
		if debug {
			fmt.Printf("%s = %d\n", result.RouteTaken, result.Cost)
		}
		if result.Cost < mincost || mincost == 0 {
			mincost = result.Cost
			minroute = result.RouteTaken
		}
	}
}

func FindLongest() {
	maxcost := 0
	maxroute := ""
	for {
		result, ok := <-Results2
		if !ok {
			Longest <- Result{maxroute, maxcost}
			return
		}
		if debug {
			fmt.Printf("%s = %d\n", result.RouteTaken, result.Cost)
		}
		if maxcost < result.Cost {
			maxcost = result.Cost
			maxroute = result.RouteTaken
		}
	}
}

func compute(route string, tovisit string) {
	if tovisit == "" {
		cost, err := CostForRoute(route)
		if err == nil {
			Results1 <- Result{route, cost}
			Results2 <- Result{route, cost}
		} else {
			fmt.Println(err)
		}
		return
	}
	for _, x := range SetValues(tovisit) {
		compute(SetAppend(route, x), SetRemove(tovisit, x))
	}
}

func CostForRoute(route string) (int, error) {
	stops := strings.Split(route, " ")
	cost := 0
	for i := 0; i < len(stops)-1; i++ {
		a := stops[i]
		b := stops[i+1]
		c1 := Costs[Route{a, b}]
		c2 := Costs[Route{b, a}]
		if c1 == 0 && c2 == 0 {
			// Impossible route
			return 0, fmt.Errorf("Can't go from %s to %s", a, b)
		}
		if c1 == 0 {
			cost += c2
		} else {
			cost += c1
		}
	}
	if cost == 0 {
		panic("Bad cost for route [" + route + "]")
	}
	return cost, nil
}

// Below here are simple string-based set operations

func SetRemove(list string, value string) string {
	if list == value {
		return ""
	}
	if strings.HasPrefix(list, value) {
		return strings.Replace(list, value+" ", "", 1)
	}
	return strings.Replace(list, " "+value, "", 1)
}

func SetAppend(list string, value string) string {
	if list == "" {
		return value
	}
	return list + " " + value
}

func SetValues(list string) []string {
	if list == "" {
		panic("Get all values from empty list?")
	}
	return strings.Split(list, " ")
}
