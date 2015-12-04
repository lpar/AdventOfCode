package main

import "testing"

// Return just the floor number return value from Process
func floor(x int, y int) int {
	return x
}

// Return just the basement instruction number return value from Process
func basement(x int, y int) int {
	return y
}

func TestProcess(t *testing.T) {
	if floor(Process("(())")) != 0 || floor(Process("()()")) != 0 {
		t.Error("(()) and ()() should both result in floor 0.")
	}

	if floor(Process("(((")) != 3 || floor(Process("(()(()(")) != 3 {
		t.Error("((( and (()(()( should both result in floor 3.")
	}
	// ))((((( also results in floor 3.

	if floor(Process("))(((((")) != 3 {
		t.Error("))((((( should also result in floor 3.")
	}

	if floor(Process("())")) != -1 || floor(Process("))(")) != -1 {
		t.Error("()) and ))( should both result in floor -1 (the first basement level).")
	}

	if floor(Process(")))")) != -3 || floor(Process(")())())")) != -3 {
		t.Error("))) and )())()) should both result in floor -3.")
	}

	if basement(Process(")")) != 1 {
		t.Error(") should cause him to enter the basement at character position 1.")
	}

	if basement(Process("()())")) != 5 {
		t.Error("()()) should cause him to enter the basement at character position 5.")
	}
}
