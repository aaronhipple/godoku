package godoku

import (
	"fmt"
)

// Puzzle is a two-dimensional array of 9 integers.
type Puzzle struct {
	values [9][9]int
}

// Print shows the puzzle in a human-readable format.
func (p Puzzle) Print() {
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
func (p Puzzle) Row(i int) (row [9]int) {
	return p.values[i]
}

// Column returns a single column of the puzzle
func (p Puzzle) Column(i int) (col [9]int) {
	for j, row := range p.values {
		col[j] = row[i]
	}
	return
}
