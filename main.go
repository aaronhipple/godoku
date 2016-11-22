package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("One file argument must be passed")
		return
	}

	path := args[0]
	puzzle, err := importPuzzle(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Given Puzzle:")
	puzzle.Print()

	puzzle.Solve()
	fmt.Println("Solved Puzzle:")
	puzzle.Print()
}

func importPuzzle(path string) (puzzle Puzzle, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','

	input, err := reader.ReadAll()
	if err != nil {
		return
	}

	for i, row := range input {
		for j, col := range row {
			puzzle.values[i][j], err = strconv.Atoi(col)
			if err != nil {
				return
			}
		}
	}

	return
}
