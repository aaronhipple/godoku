package main

import (
	"sync"
)

// Mark examines a Puzzle and returns a MarkedPuzzle -- that is, a Puzzle with markings indicating possible values.
func (p *Puzzle) Mark() (mp MarkedPuzzle) {
	var row, col int
	var wg sync.WaitGroup

	for row = 0; row < 9; row++ {
		for col = 0; col < 9; col++ {
			wg.Add(1)

			go func(p *Puzzle, mp *MarkedPuzzle, row, col int) {
				defer wg.Done()
				if p.values[row][col] != 0 {
					mp.possibleValues[row][col] = []int{p.values[row][col]}
					return
				}

				mp.possibleValues[row][col] = p.PossibleValues(row, col)
			}(p, &mp, row, col)
		}
	}

	wg.Wait()

	mp.StrikeMatches()

	return
}

func inCollection(n int, c [9]int) bool {
	for i := 0; i < 9; i++ {
		if c[i] == n {
			return true
		}
	}
	return false
}

// Solve solves a Puzzle.
func (p *Puzzle) Solve() {
	for i := 0; i < 200; i++ {
		mp := p.Mark()
		p.ApplyMarks(mp)
	}
}
