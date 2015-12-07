package main

import "testing"

func TestGrid(t *testing.T) {
	grid := new(Grid)
	grid.Set(0, 0, 999, 999, true)
	if grid.Lit() != 1000*1000 {
		t.Error("Failed to light every light")
	}
	grid.Set(499, 499, 500, 500, false)
	if grid.Lit() != 1000*1000-4 {
		t.Error("Failed to turn off the middle four")
	}
	grid = new(Grid)
	grid.Toggle(0, 0, 999, 0)
	if grid.Lit() != 1000 {
		t.Error("Failed to toggle one row")
	}
}
