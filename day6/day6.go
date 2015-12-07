package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"regexp"
	"strconv"
)

const WIDTH = 1000
const HEIGHT = 1000
const MAXX = WIDTH - 1
const MAXY = HEIGHT - 1

type Grid struct {
	lights [WIDTH * HEIGHT]bool
}

// Factor out the conversion from 2 dimensions to 1; Go will inline this
func idx(x, y int) int {
	return x + y*WIDTH
}

// Map is the classic map higher-order function and applies the provided
// function to all the grid cells in the rectangle (x1,y1) (x2,y2)
func (g *Grid) Map(x1, y1, x2, y2 int, mapfunc func(bool) bool) {
	for x := x1; x <= x2; x += 1 {
		for y := y1; y <= y2; y += 1 {
			i := idx(x, y)
			g.lights[i] = mapfunc(g.lights[i])
		}
	}
}

// Convenience function to Map the whole Grid
func (g *Grid) MapAll(mapfunc func(bool) bool) {
	g.Map(0, 0, MAXX, MAXY, mapfunc)
}

// Set sets the grid cells in the rectangle (x1,y1)..(x2,y2) to the
// specified value.
// TODO: Compare speed with use of Map, to see if this is necessary.
func (g *Grid) Set(x1, y1, x2, y2 int, value bool) {
	for x := x1; x <= x2; x += 1 {
		for y := y1; y <= y2; y += 1 {
			i := idx(x, y)
			g.lights[i] = value
		}
	}
}

func (g *Grid) Toggle(x1, y1, x2, y2 int) {
	g.Map(x1, y1, x2, y2, func(x bool) bool { return !x })
}

// Dump grid as text for debugging purposes
func (g *Grid) Dump() {
	for y := 0; y < MAXY; y += 1 {
		for x := 0; x < MAXX; x += 1 {
			if g.lights[idx(x, y)] {
				fmt.Printf("%d,%d on\n", x, y)
			}
		}
	}
}

func (g *Grid) DumpPNG(filename string) {
	img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))
	for y := 0; y < MAXY; y += 1 {
		for x := 0; x < MAXX; x += 1 {
			if g.lights[idx(x, y)] {
				img.Set(x, MAXY-y, color.RGBA{255, 255, 255, 255})
			} else {
				img.Set(x, MAXY-y, color.RGBA{0, 0, 0, 255})
			}
		}
	}
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}

func (g *Grid) Lit() (total int) {
	g.Map(0, 0, MAXX, MAXY, func(x bool) bool {
		if x {
			total++
		}
		return x
	})
	return total
}

var Command = regexp.MustCompile(`^(turn on|turn off|toggle)\s+(\d+),(\d+)\s+through\s+(\d+),(\d+)$`)

func strtoint(a string) int {
	b, err := strconv.ParseInt(a, 10, 0)
	if err != nil {
		panic("Unable to parse integer [" + a + "]")
	}
	return int(b)
}

func (g *Grid) Perform(command string) {
	m := Command.FindStringSubmatch(command)
	if m == nil {
		panic("Unable to parse instruction [" + command + "]")
	}
	op := m[1]
	x1, y1 := strtoint(m[2]), strtoint(m[3])
	x2, y2 := strtoint(m[4]), strtoint(m[5])
	switch op {
	case "turn on":
		fmt.Printf("%d,%d .. %d,%d on\n", x1, y1, x2, y2)
		g.Set(x1, y1, x2, y2, true)
	case "turn off":
		fmt.Printf("%d,%d .. %d,%d off\n", x1, y1, x2, y2)
		g.Set(x1, y1, x2, y2, false)
	case "toggle":
		fmt.Printf("%d,%d .. %d,%d toggle\n", x1, y1, x2, y2)
		g.Toggle(x1, y1, x2, y2)
	}
}

func Process(filename string) int {
	inf, err := os.Open(filename)
	if err != nil {
		fmt.Println("Cannot open input file")
		os.Exit(4)
	}
	defer inf.Close()
	grid := new(Grid)

	scanner := bufio.NewScanner(inf)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		input := scanner.Text()
		grid.Perform(input)
	}
	return grid.Lit()
}

// First simple test was:
// func poc() {
// 	grid := new(Grid)
// 	grid.Map(3, 3, 6, 7, func(x bool) bool { return !x })
// 	grid.Dump()
// 	fmt.Printf("Number lit = %d\n", grid.Lit())
// 	grid.Map(5, 6, 90, 60, func(x bool) bool { return !x })
// 	grid.DumpPNG("out.png")
// }

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing input file")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		os.Exit(2)
	}
	fmt.Printf("Total lit = %d\n", Process(os.Args[1]))
}
