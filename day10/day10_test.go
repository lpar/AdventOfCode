package main

import "testing"

func TestCircuit(t *testing.T) {

	if encode("1") != "11" {
		t.Error("Encode fail 1")
	}
	if encode("11") != "21" {
		t.Error("Encode fail 11")
	}
	if encode("21") != "1211" {
		t.Error("Encode fail 21")
	}
	if encode("1211") != "111221" {
		t.Error("Encode fail 1211")
	}
	if encode("111221") != "312211" {
		t.Error("Encode fail 111221")
	}

}
