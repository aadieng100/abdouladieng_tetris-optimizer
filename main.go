package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	
	// The application expects exactly one argument: the path to the input text file.
	if len(args) == 1 {
		// Read the entire file content.
		file, err := os.ReadFile(args[0])
		if err != nil {
			Error()
		}
		
		// Parse the raw byte file into a slice of 4x4 boards (tetrominoes).
		tetrominoes := TabMinoes(file)
		
		// Extract the coordinates of the '#' blocks for each tetromino and validate their shapes.
		indexs := TakeAllIndexs(tetrominoes)
		
		// Solve the packing problem to find the smallest square board.
		finalTab := Solve(indexs, len(indexs), 0)
		
		// Print the final board layout to stdout.
		Printer(finalTab)
		return
	}
	
	fmt.Println("Only one argument is allowed.")
}
