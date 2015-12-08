package main

import "testing"

func TestCircuit(t *testing.T) {
	c := NewCircuit()

	// 123 -> x
	c.Execute("123 -> x")
	if v, _ := c.GetWire("x"); v != 123 {
		t.Error("Failure at 123 -> x")
	}

	// 456 -> y
	c.Execute("456 -> y")
	if v, _ := c.GetWire("y"); v != 456 {
		t.Error("Failure at 456 -> y")
	}

	// x AND y -> d
	c.Execute("x AND y -> d")
	if v, _ := c.GetWire("d"); v != 72 {
		t.Error("Failure at x AND y -> d")
	}

	// x OR y -> e
	c.Execute("x OR y -> e")
	if v, _ := c.GetWire("e"); v != 507 {
		t.Error("Failure at x OR y -> e")
	}

	// x LSH	if v,_ :=T 2 -> f
	c.Execute("x LSHIFT 2 -> f")
	if v, _ := c.GetWire("f"); v != 492 {
		t.Error("Failure at x LSH	if v,_ :=T 2 -> e")
	}

	// y RSH	if v,_ :=T 2 -> g
	c.Execute("y RSHIFT 2 -> g")
	if v, _ := c.GetWire("g"); v != 114 {
		t.Error("Failure at y RSH	if v,_ :=T 2 -> g")
	}

	// NOT x -> h
	c.Execute("NOT x -> h")
	if v, _ := c.GetWire("h"); v != 65412 {
		t.Error("Failure at NOT x -> h")
	}

	// NOT y -> i
	c.Execute("NOT y -> i")
	if v, _ := c.GetWire("i"); v != 65079 {
		t.Error("Failure at NOT y -> i")
	}

	c.Execute("1674 -> bc")
	if v, _ := c.GetWire("bc"); v != 1674 {
		t.Error("Two character wire IDs failed")
	}

}
