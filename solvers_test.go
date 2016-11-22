package main

import (
	"fmt"
	"reflect"
	"testing"
)

var mbm MarkedPuzzle

// BenchmarkMarkPuzzle tests the speed of the MarkPuzzle function.
func BenchmarkMarkPuzzle(b *testing.B) {
	var mp MarkedPuzzle
	for i := 0; i < b.N; i++ {
		mp = UnsolvedPuzzle.Mark()
	}
	mbm = mp
}

// TestMarkPuzzle tests that the MarkPuzzle function is executing properly.
func TestMarkPuzzle(t *testing.T) {
	t.Skip("Not yet implemented")
}

// TestSolve tests that the Solve function correctly solves a puzzle.
func TestSolve(t *testing.T) {
	p := UnsolvedPuzzle

	p.Solve()

	if !reflect.DeepEqual(p.values, SolvedPuzzle.values) {
		t.Error("Error solving puzzle.")
		fmt.Println("Expected:")
		SolvedPuzzle.Print()
		fmt.Println("Got:")
		p.Print()
	}
}
