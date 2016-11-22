package main

import (
	"reflect"
)

// MarkedPuzzle is a two-dimensional array containing slices of possible values.
type MarkedPuzzle struct {
	possibleValues [9][9][]int
}

// Row returns a single row of the puzzle
func (mp *MarkedPuzzle) Row(row int) (collection [9]*[]int) {
	for i := 0; i < 9; i++ {
		collection[i] = &mp.possibleValues[row][i]
	}
	return
}

// Column returns a single column of the puzzle
func (mp *MarkedPuzzle) Column(col int) (collection [9]*[]int) {
	for i := 0; i < 9; i++ {
		collection[i] = &mp.possibleValues[i][col]
	}
	return
}

// Box returns a single box of the puzzle
func (mp *MarkedPuzzle) Box(box int) (collection [9]*[]int) {
	firstRow := (box / 3) * 3
	firstCol := (box % 3) * 3

	var row, col int
	for i := 0; i < 9; i++ {
		row = firstRow + (i / 3)
		col = firstCol + (i % 3)
		collection[i] = &mp.possibleValues[row][col]
	}
	return
}

// StrikeMatches removes possible values when they're excluded by a matched pair elsewhere in the row/column/box.
func (mp *MarkedPuzzle) StrikeMatches() {
	for i := 0; i < 9; i++ {
		excludeMatches(mp.Row(i))
	}
	for i := 0; i < 9; i++ {
		excludeMatches(mp.Column(i))
	}
	for i := 0; i < 9; i++ {
		excludeMatches(mp.Box(i))
	}
}

func excludeMatches(collection [9]*[]int) {
	for i := 0; i < 9; i++ {
		if len(*collection[i]) == 1 {
			continue
		}

		matchLength := len(*collection[i])
		matches := []int{i}

		for j := i + 1; j < 9; j++ {
			if j == i {
				continue
			}

			if reflect.DeepEqual(collection[i], collection[j]) {
				matches = append(matches, j)
			}
		}

		if len(matches) != matchLength {
			continue
		}

		for k := 0; k < 9; k++ {
			if contains(k, matches) {
				continue
			}

			badValues := *collection[i]
			for l := 0; l < len(badValues); l++ {
				removeValue(collection[k], badValues[l])
			}
		}
	}
}

func contains(n int, ns []int) bool {
	for i := 0; i < len(ns); i++ {
		if ns[i] == n {
			return true
		}
	}
	return false
}

func removeValue(collection *[]int, bv int) {
	nv := *collection
	for i := 0; i < len(nv); i++ {
		if nv[i] == bv {
			nv = append(nv[:i], nv[i+1:]...)
			break
		}
	}

	*collection = nv
}
