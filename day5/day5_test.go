package main

import "testing"

func TestNice1(t *testing.T) {
	if !Nice1("ugknbfddgicrmopn") {
		t.Error("ugknbfddgicrmopn should be nice")
	}
	if !Nice1("aaa") {
		t.Error("aaa should be nice")
	}
	if Nice1("jchzalrnumimnmhp") {
		t.Error("jchzalrnumimnmhp should be naughty")
	}
	if Nice1("haegwjzuvuyypxyu") {
		t.Error("haegwjzuvuyypxyu should be naughty")
	}
	if Nice1("dvszwmarrgswjxmb") {
		t.Error("dvszwmarrgswjxmb should be naughty")
	}
}

func TestNice2(t *testing.T) {
	if !Nice2("qjhvhtzxzqqjkmpb") {
		t.Error("qjhvhtzxzqqjkmpb should be nice")
	}
	if !Nice2("xxyxx") {
		t.Error("xxyxx should be nice")
	}
	if Nice2("uurcxstgmygtbstg") {
		t.Error("uurcxstgmygtbstg should be naughty")
	}
	if Nice2("ieodomkazucvgmuy") {
		t.Error("ieodomkazucvgmuy should be naughty")
	}
}
