package main

import (
	"fmt"
)

// Puzzle is a two-dimensional array of 9 integers.
type Puzzle struct {
	values [9][9]int
}

// Print shows the puzzle in a human-readable format.
func (p *Puzzle) Print() {
	for i, row := range p.values {
		for j, col := range row {
			fmt.Printf("%d", col)
			switch {
			case (j+1)%3 == 0 && j < 8:
				fmt.Printf(" | ")
			case (j + 1) == 9:
				fmt.Printf("\n")
			default:
				fmt.Printf(" ")
			}
		}
		if (i+1)%3 == 0 && i < 8 {
			fmt.Println("---------------------")
		}
	}
}

// Row returns a single row of the puzzle
func (p *Puzzle) Row(i int) (row [9]int) {
	return p.values[i]
}

// Column returns a single column of the puzzle
func (p *Puzzle) Column(i int) (col [9]int) {
	for j, row := range p.values {
		col[j] = row[i]
	}
	return
}

// Box returns a single box of the puzzle
func (p *Puzzle) Box(i int) (box [9]int) {
	firstRow := (i / 3) * 3
	firstCol := (i % 3) * 3

	var row, col int
	for i := 0; i < 9; i++ {
		row = firstRow + (i / 3)
		col = firstCol + (i % 3)
		box[i] = p.values[row][col]
	}
	return
}

// ApplyMarks updates a Puzzle given the marks from a MarkedPuzzle
func (p *Puzzle) ApplyMarks(mp MarkedPuzzle) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {

			// Skip known values
			if p.values[row][col] != 0 {
				continue
			}

			// When only one possible value, insert
			if len(mp.possibleValues[row][col]) == 1 {
				p.values[row][col] = mp.possibleValues[row][col][0]
			}

		}
	}
}

// PossibleValues returns a slice of possible values for a cell.
func (p *Puzzle) PossibleValues(row, col int) (pv []int) {
	pv = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	box := boxNumber(row, col)

	for i := len(pv) - 1; i >= 0; i-- {
		if inCollection(pv[i], p.Row(row)) ||
			inCollection(pv[i], p.Column(col)) ||
			inCollection(pv[i], p.Box(box)) {
			pv = append(pv[:i], pv[i+1:]...)
		}
	}

	return
}

func boxNumber(row int, col int) (box int) {
	return ((row / 3) * 3) + (col / 3)
}
