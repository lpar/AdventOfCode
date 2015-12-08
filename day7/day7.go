package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const debug = false

type Circuit struct {
	wires     map[string]uint16
	boverride uint16
}

var Whitespace = regexp.MustCompile(`\s+`)

const Arrow = "->"

func NewCircuit() *Circuit {
	return &Circuit{wires: make(map[string]uint16)}
}

func (c *Circuit) GetWire(id string) (uint16, error) {
	if id == "b" && c.boverride != 0 {
		return c.boverride, nil
	}
	value, ok := c.wires[id]
	if !ok {
		return 0, fmt.Errorf("%s has no value yet", id)
	}
	return value, nil
}

func (c *Circuit) OverrideB(value uint16) {
	c.boverride = value
}

func (c *Circuit) SetWire(id string, value uint16) {
	old, ok := c.wires[id]
	if ok {
		fmt.Printf("%s redefined from %d to %d\n", id, old, value)
	}
	c.wires[id] = value
}

func (c *Circuit) Execute(command string) error {
	// Parse it ourselves
	words := Whitespace.Split(command, -1)
	var err error
	switch {
	case words[1] == Arrow:
		// foo -> bar
		err = c.wire(words[0], words[2])
	case words[2] == Arrow:
		// NOT foo -> bar
		err = c.not(words[1], words[3])
	case words[3] == Arrow:
		// foo OP bar -> baz
		err = c.op(words[1], words[0], words[2], words[4])
	default:
		panic("Could not execute instruction " + command)
	}
	return err
}

func (c *Circuit) wire(src, dest string) error {
	x, err := c.eval(src)
	if err != nil {
		return err
	}
	c.SetWire(dest, x)
	if debug {
		fmt.Printf("%s -> %s\n", src, dest)
	}
	return nil
}

func (c *Circuit) not(src, dest string) error {
	x, err := c.eval(src)
	if err != nil {
		return err
	}
	c.SetWire(dest, ^x)
	if debug {
		fmt.Printf("^%s -> %s\n", src, dest)
	}
	return nil
}

// eval(expr) evaluates a one-word expression which is either an integer
// or the name of a wire.
func (c *Circuit) eval(expr string) (uint16, error) {
	a, err := strconv.ParseInt(expr, 10, 16)
	if err != nil {
		return c.GetWire(expr)
	}
	return uint16(a), nil
}

func (c *Circuit) op(opcode, foo, bar, dest string) error {
	fooval, err := c.eval(foo)
	if err != nil {
		return err
	}
	barval, err := c.eval(bar)
	if err != nil {
		return err
	}
	var destval uint16
	success := false
	switch opcode {
	case "RSHIFT":
		destval = fooval >> barval
		success = true
	case "LSHIFT":
		destval = fooval << barval
		success = true
	case "AND":
		destval = fooval & barval
		success = true
	case "OR":
		destval = fooval | barval
		success = true
	}
	if !success {
		panic("Unable to interpret opcode " + opcode + " " + foo + " " + bar + " -> " + dest)
	}
	if debug {
		fmt.Printf("%s %s %s -> %s\n", opcode, foo, bar, dest)
	}
	c.SetWire(dest, destval)
	return nil
}

func (c *Circuit) Reset() {
	fmt.Printf("Reset ")
	for key, _ := range c.wires {
		delete(c.wires, key)
		fmt.Printf("%s ", key)
	}
	fmt.Println()
}

func (c *Circuit) Run(program []string) uint16 {
	for {
		for _, instr := range program {
			if c.Execute(instr) == nil {
				if debug {
					fmt.Println(instr)
				}
			}
		}
		a, err := c.GetWire("a")
		if err == nil {
			fmt.Printf("a = %d\n", a)
			return a
		}
	}
}

func Process(filename string) {
	inf, err := os.Open(filename)
	if err != nil {
		fmt.Println("Cannot open input file")
		os.Exit(4)
	}
	defer inf.Close()

	// Read the circuit code
	var code []string
	scanner := bufio.NewScanner(inf)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		input := scanner.Text()
		code = append(code, input)
	}

	circuit := NewCircuit()
	a1 := circuit.Run(code)

	circuit.OverrideB(a1)
	circuit.Reset()
	a2 := circuit.Run(code)

	fmt.Printf("Part 1: a = %d\n", a1)
	fmt.Printf("Part 2: a = %d\n", a2)
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
	Process(os.Args[1])
}
