package main

import (
	"strings"
	"testing"
)

func Part1Process(inst string) int {
	presents := make(map[Coordinates]int)
	Process(presents, inst)
	return Report(presents)
}

func TestProcess(t *testing.T) {
	if Part1Process(">") != 2 {
		t.Error("> should be 2")
	}
	if Part1Process("^>v<") != 4 {
		t.Error("^>v< should be 4")
	}
	if Part1Process("^v^v^v^v^v") != 2 {
		t.Error("^v^v^v^v^v should be 2")
	}
}

func TestSplitInstructions(t *testing.T) {
	odd, even := SplitInstructions("<>^v^<>v^")
	if strings.Compare(odd, "<^^>^") != 0 {
		t.Error("SplitInstructions malfunction, odd wrong: " + odd)
	}
	if even != ">v<v" {
		t.Error("SplitInstructions malfunction, even wrong: " + even)
	}
}
