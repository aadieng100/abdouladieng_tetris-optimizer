package main

import (
	"fmt"
	"os"
)

// TabMinoes parses the raw file bytes into a list of 4x4 tetrominoes.
// It enforces that each tetromino is exactly 4x4 characters, and separated by a single newline.
func TabMinoes(file []byte) [][][]byte {
	row := []byte{}
	tetromino := [][]byte{}
	tetrominoes := [][][]byte{}
	
	for i := 0; i < len(file); i++ {
		// Handle EOF (last character of the file)
		if i == len(file)-1 {
			if file[i] != '\n' {
				row = append(row, file[i])
			}
			if len(row) == 4 {
				tetromino = append(tetromino, row)
				tetrominoes = append(tetrominoes, tetromino)
				continue
			}
		}
		
		// Handle newline characters which separate lines
		if i+1 < len(file) && file[i] == '\n' {
			if file[i+1] != '\n' {
				// Single newline: marks the end of a row in the current tetromino grid
				if len(row) == 4 {
					tetromino = append(tetromino, row)
					row = nil
					continue
				}
				Error()
			} else {
				// Double newline: marks the boundary between two tetromino grids
				if len(row) == 4 {
					tetromino = append(tetromino, row)
					row = nil
				}
				if len(tetromino) == 4 {
					tetrominoes = append(tetrominoes, tetromino)
					tetromino = nil
					i++ // Skip the second newline
					continue
				}
				Error()
			}
		}
		
		row = append(row, file[i])
	}
	return tetrominoes
}

// VerifMinoes verifies that a tetromino is orthogonally connected.
// It checks that each row's '#' blocks are contiguous and that there is at least one
// vertical connection between adjacent rows of the tetromino.
func VerifMinoes(coords [][]int) {
	for i := 0; i < len(coords)-1; i++ {
		if i+1 < len(coords)-1 {
			connected := false
			// Check that blocks on row i are contiguous
			for j := 0; j < len(coords[i]); j++ {
				if j+1 < len(coords[i]) && coords[i][j+1] != coords[i][j]+1 {
					Error()
				}
				// Check for at least one orthogonal vertical connection with row i+1
				for k := 0; k < len(coords[i+1]); k++ {
					if coords[i][j] == coords[i+1][k] {
						connected = true
					}
				}
			}
			if !connected {
				Error()
			}
		}
	}
}

// GetIndexs extracts the 0-indexed column coordinates for each row containing a '#' character.
// It verifies that there are exactly 4 '#' characters and 12 '.' characters per tetromino.
// It appends the character label identifier (e.g. 'A', 'B') as the final element.
func GetIndexs(tetromino [][]byte, char int) [][]int {
	rowCoords := []int{}
	coords := [][]int{}
	hashCount := 0
	dotCount := 0
	
	for i := 0; i < len(tetromino); i++ {
		for j := 0; j < len(tetromino[i]); j++ {
			if tetromino[i][j] == '#' {
				hashCount++
				rowCoords = append(rowCoords, j)
			}
			if tetromino[i][j] == '.' {
				dotCount++
			}
		}
		// A row check to fail early if unexpected structures are found
		if hashCount > 0 && hashCount < 4 && len(rowCoords) == 0 {
			Error()
		}
		if len(rowCoords) != 0 {
			coords = append(coords, rowCoords)
			rowCoords = nil
		}
	}
	
	// A valid tetromino must have exactly 4 '#' blocks and 12 '.' empty spaces
	if hashCount != 4 || hashCount*3 != dotCount {
		Error()
	}
	
	// Store the tetromino character label at the end of the coordinate slice
	labelInfo := []int{char}
	coords = append(coords, labelInfo)
	return coords
}

// TakeAllIndexs extracts coordinates for all tetrominoes and runs shape verification.
func TakeAllIndexs(tetrominoes [][][]byte) [][][]int {
	allCoords := [][][]int{}
	for i, char := 0, 65; i < len(tetrominoes); i, char = i+1, char+1 {
		allCoords = append(allCoords, GetIndexs(tetrominoes[i], char))
		VerifMinoes(allCoords[i])
	}
	return allCoords
}

// Error prints "ERROR" to stdout and terminates the program with exit code 0 as required.
func Error() {
	fmt.Println("ERROR")
	os.Exit(0)
}
