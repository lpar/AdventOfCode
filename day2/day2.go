package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// min returns the 2 smallest of 3 ints
func min2(a, b, c int) (int, int) {
	abc := []int{a, b, c}
	sort.Ints(abc)
	return abc[0], abc[1]
}

func WrappingPaper(l, w, h int) int {
	lw := l * w
	wh := w * h
	hl := h * l
	slack, _ := min2(lw, wh, hl)
	tot := 2*lw + 2*wh + 2*hl + slack
	return tot
}

func Ribbon(l, w, h int) int {
	s1, s2 := min2(l, w, h)
	volume := l * w * h
	return s1*2 + s2*2 + volume
}

func StrToInt(x string) int {
	y, err := strconv.ParseInt(x, 10, 0)
	if err != nil {
		panic(fmt.Sprintf("%s not int: %v", x, err))
	}
	return int(y)
}

func Parse(input string) (int, int, int) {
	inp := strings.Split(input, "x")
	if len(inp) != 3 {
		fmt.Println("Parse failed for %s", input)
		os.Exit(3)
	}
	l := StrToInt(inp[0])
	w := StrToInt(inp[1])
	h := StrToInt(inp[2])
	return l, w, h
}

func RibbonNeeded(input string) int {
	return Ribbon(Parse(input))
}

func WrappingPaperNeeded(input string) int {
	return WrappingPaper(Parse(input))
}

func Process(fname string) (paper int, ribbon int) {
	inf, err := os.Open(fname)
	if err != nil {
		fmt.Println("Cannot open input file")
		os.Exit(4)
	}
	defer inf.Close()

	papertotal := 0
	ribbontotal := 0
	scanner := bufio.NewScanner(inf)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		dimens := scanner.Text()
		l, w, h := Parse(dimens)
		paper := WrappingPaper(l, w, h)
		ribbon := Ribbon(l, w, h)
		fmt.Printf("%s = %d wrapping paper, %d ribbon\n", dimens, paper, ribbon)
		papertotal += paper
		ribbontotal += ribbon
	}
	return papertotal, ribbontotal
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

	papertotal, ribbontotal := Process(os.Args[1])
	fmt.Printf("Total wrapping paper needed = %d\n", papertotal)
	fmt.Printf("Total ribbon needed = %d\n", ribbontotal)

}
