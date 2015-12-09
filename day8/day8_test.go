package main

import (
	"fmt"
	"testing"
)

func TestCircuit(t *testing.T) {

	if x, _ := Decode(`"abc"`); x != `abc` || len(x) != 3 {
		t.Error("Quote removal failed")
	}

	if x, _ := Decode(`"\x27"`); x != `'` || len(x) != 1 {
		t.Error("Hex decode 2 failed")
	}

	if x, _ := Decode(`""`); x != "" || len(x) != 0 {
		t.Error("Empty string decode failed")
	}

	if x, _ := Decode(`"aaa\"aaa"`); x != `aaa"aaa` || len(x) != 7 {
		t.Error("String with escaped quotes decode failed")
	}

	if x, _ := Decode(`"\x48\x65\x78"`); x != `Hex` {
		t.Error("Hex decode 2 failed")
	}

	if x := Encode(`""`); x != `"\"\""` || len(x) != 6 {
		fmt.Println(x)
		fmt.Println(len(x))
		t.Error("Quotes encode failed")
	}

	if x := Encode(`"abc"`); x != `"\"abc\""` || len(x) != 9 {
		t.Error("Quoted abc encode failed")
	}

	if x := Encode(`"aaa\"aaa"`); x != `"\"aaa\\\"aaa\""` || len(x) != 16 {
		t.Error("Quoted long encode failed")
	}

	if x := Encode(`"\x27"`); x != `"\"\\x27\""` || len(x) != 11 {
		t.Error("Quoted hex escape encode failed")
	}

}
