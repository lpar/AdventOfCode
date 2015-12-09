package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var Hex = regexp.MustCompile(`\\x[[:xdigit:]]{2}`)
var Escaped = regexp.MustCompile(`\\.`)

func decodeHex(x string) string {
	y, err := hex.DecodeString(x[2:])
	if err != nil {
		panic("Invalid hex escape \\x" + x)
	}
	return string(y)
}

func nobs(x string) string {
	fmt.Println("[" + x + "]")
	return x[1:]
}

func Decode(input string) (string, error) {
	// Sanity check the outer quotes are there before we remove them
	if !strings.HasPrefix(input, `"`) || !strings.HasSuffix(input, `"`) {
		return input, fmt.Errorf("string `%s` not enclosed in quotes", input)
	}
	// Remove outer quotes and decode hex escapes
	input = Hex.ReplaceAllStringFunc(input[1:len(input)-1], decodeHex)
	// Decode all other backslash escapes by stripping the backslash
	input = Escaped.ReplaceAllStringFunc(input,
		func(x string) string { return x[1:] })
	return input, nil
}

func Encode(input string) string {
	result := make([]rune, 0, 2*len(input))
	result = append(result, '"')
	for _, c := range input {
		if c == '"' || c == '\\' {
			result = append(result, '\\')
		}
		result = append(result, c)
	}
	result = append(result, '"')
	return string(result)
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
	inf, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Cannot open input file")
		os.Exit(4)
	}
	defer inf.Close()

	datasize := 0
	codesize := 0
	encsize := 0
	scanner := bufio.NewScanner(inf)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		input := scanner.Text()
		value, err := Decode(input)
		if err != nil {
			panic(err)
		}
		encval := Encode(input)
		datasize += len(value)
		codesize += len(input)
		encsize += len(encval)
	}
	fmt.Println(codesize - datasize)
	fmt.Println(encsize - codesize)
}
