package main

import "testing"

func TestCircuit(t *testing.T) {

	if SetRemove("foo bar baz", "foo") != "bar baz" {
		t.Error("SetRemove fail 1")
	}

	if SetRemove("foo bar baz", "bar") != "foo baz" {
		t.Error("SetRemove fail 2")
	}

	if SetRemove("foo bar baz", "baz") != "foo bar" {
		t.Error("SetRemove fail 3")
	}

	if SetRemove("foo", "foo") != "" {
		t.Error("SetRemove fail 4")
	}

}
