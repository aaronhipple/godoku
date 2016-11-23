package main

import (
	"reflect"
	"testing"
)

// UnsolvedPuzzle is an unsolved puzzle for testing purposes.
var UnsolvedPuzzle = Puzzle{
	[9][9]int{
		[9]int{2, 9, 5, 7, 0, 0, 8, 6, 0},
		[9]int{0, 3, 1, 8, 6, 5, 0, 2, 0},
		[9]int{8, 0, 6, 0, 0, 0, 0, 0, 0},
		[9]int{0, 0, 7, 0, 5, 0, 0, 0, 6},
		[9]int{6, 1, 0, 3, 8, 7, 0, 0, 0},
		[9]int{5, 0, 0, 0, 1, 6, 7, 0, 0},
		[9]int{0, 0, 3, 5, 0, 0, 1, 0, 9},
		[9]int{0, 2, 0, 6, 0, 0, 3, 5, 0},
		[9]int{0, 5, 4, 0, 0, 8, 6, 7, 2},
	},
}

// SolvedPuzzle is a solved puzzle for testing purposes.
var SolvedPuzzle = Puzzle{
	[9][9]int{
		[9]int{2, 9, 5, 7, 4, 3, 8, 6, 1},
		[9]int{4, 3, 1, 8, 6, 5, 9, 2, 7},
		[9]int{8, 7, 6, 1, 9, 2, 5, 4, 3},
		[9]int{3, 8, 7, 4, 5, 9, 2, 1, 6},
		[9]int{6, 1, 2, 3, 8, 7, 4, 9, 5},
		[9]int{5, 4, 9, 2, 1, 6, 7, 3, 8},
		[9]int{7, 6, 3, 5, 2, 4, 1, 8, 9},
		[9]int{9, 2, 8, 6, 7, 1, 3, 5, 4},
		[9]int{1, 5, 4, 9, 3, 8, 6, 7, 2},
	},
}

// TestRow ensures that the Row method returns the correct values.
func TestRow(t *testing.T) {
	expectedRow := [9]int{8, 7, 6, 1, 9, 2, 5, 4, 3}
	if !reflect.DeepEqual(SolvedPuzzle.Row(2), expectedRow) {
		t.Error("Expected", expectedRow, "got", SolvedPuzzle.Row(2))
	}
}

// TestColumn ensures that the Column method returns the correct values.
func TestColumn(t *testing.T) {
	expectedColumn := [9]int{5, 1, 6, 7, 2, 9, 3, 8, 4}
	if !reflect.DeepEqual(SolvedPuzzle.Column(2), expectedColumn) {
		t.Error("Expected", expectedColumn, "got", SolvedPuzzle.Column(2))
	}
}

// TestBox ensures that the Box method returns the correct values.
func TestBox(t *testing.T) {
	expectedBox := [9]int{5, 2, 4, 6, 7, 1, 9, 3, 8}
	if !reflect.DeepEqual(SolvedPuzzle.Box(7), expectedBox) {
		t.Error("Expected", expectedBox, "got", SolvedPuzzle.Box(7))
	}
}
