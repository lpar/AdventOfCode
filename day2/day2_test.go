package main

import "testing"

func TestWrappingPaper(t *testing.T) {
	if WrappingPaperNeeded("2x3x4") != 58 ||
		WrappingPaperNeeded("2x4x3") != 58 ||
		WrappingPaperNeeded("4x3x2") != 58 ||
		WrappingPaperNeeded("4x2x3") != 58 {
		t.Error("2x3x4 should be 58 sq ft")
	}

	if WrappingPaperNeeded("1x1x10") != 43 {
		t.Error("1x1x10 should be 43 sq ft")
	}

	if RibbonNeeded("2x3x4") != 34 {
		t.Error("2x3x4 should need 34 ft of ribbon")
	}

	if RibbonNeeded("1x1x10") != 14 {
		t.Error("1x1x10 should need 14 ft of ribbon")
	}

}
