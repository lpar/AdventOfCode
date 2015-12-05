package main

import (
	"bufio"
	"os"
	"testing"
)

func TestNice1(t *testing.T) {
	var naughty = []string{"jchzalrnumimnmhp", "haegwjzuvuyypxyu",
		"dvszwmarrgswjxmb"}
	var nice = []string{"ugknbfddgicrmopn", "aaa", "eéeeéé", "e🎅🎅aa🎅i"}

	for _, s := range nice {
		if !Nice1(s) {
			t.Error(s + " should be nice")
		}
		if !Nice1RegExp(s) {
			t.Error(s + " should be nice (regexp)")
		}
	}
	for _, s := range naughty {
		if Nice1(s) {
			t.Error(s + " should be naughty")
		}
		if Nice1RegExp(s) {
			t.Error(s + " should be naughty (regexp)")
		}
	}
}

func TestNice2(t *testing.T) {
	var naughty = []string{"uurcxstgmygtbstg", "ieodomkazucvgmuy"}
	var nice = []string{"qjhvhtzxzqqjkmpb", "xxyxx", "ááüáá", "🎅yz🎅☃🎅yzy🎅"}

	for _, s := range nice {
		if !Nice2(s) {
			t.Error(s + " should be nice")
		}
		if !Nice2RegExp(s) {
			t.Error(s + " should be nice (regexp)")
		}
	}
	for _, s := range naughty {
		if Nice2(s) {
			t.Error(s + " should be naughty")
		}
		if Nice2RegExp(s) {
			t.Error(s + " should be naughty (regexp)")
		}
	}
}

func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func BenchmarkNice2(b *testing.B) {
	testData, _ := ReadFile("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, line := range testData {
			_ = Nice2(line)
		}
	}
}

func BenchmarkNice2RegExp(b *testing.B) {
	testData, _ := ReadFile("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, line := range testData {
			_ = Nice2RegExp(line)
		}
	}
}

func BenchmarkNice1(b *testing.B) {
	testData, _ := ReadFile("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, line := range testData {
			_ = Nice1(line)
		}
	}
}

func BenchmarkNice1RegExp(b *testing.B) {
	testData, _ := ReadFile("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, line := range testData {
			_ = Nice1RegExp(line)
		}
	}
}
