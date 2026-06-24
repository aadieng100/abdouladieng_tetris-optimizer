package main

import (
	"fmt"
	"math"
)

// Point represents a 2D coordinate offset.
type Point struct {
	X int
	Y int
}

// Solver contains the state of the backtracking search.
type Solver struct {
	board       [][]byte   // N x N grid of bytes
	boardSize   int        // board dimension (N)
	tetros      [][]Point  // normalized coordinates for each tetromino
	labels      []byte     // letter labels ('A', 'B', ...) for each tetromino
	next        []int      // index of next identical tetromino shape (symmetry breaking)
	placed      []bool     // tracks which tetrominoes have been placed
	sawIt       []bool     // symmetry breaking flag
	numTetros   int        // total number of tetrominoes to place
	placedCount int        // count of currently placed tetrominoes
}

// CreateTab initializes a square board of the given size filled with empty spaces ('.').
func CreateTab(size int) [][]byte {
	finalTab := make([][]byte, size)
	for i := range finalTab {
		finalTab[i] = make([]byte, size)
		for j := range finalTab[i] {
			finalTab[i][j] = '.'
		}
	}
	return finalTab
}

// canPlace checks if the tetromino z can be placed with its reference block at (sX, sY).
func (s *Solver) canPlace(z int, sX, sY int) bool {
	for _, p := range s.tetros[z] {
		x := sX + p.X
		y := sY + p.Y
		if x < 0 || x >= s.boardSize || y < 0 || y >= s.boardSize || s.board[x][y] != '.' {
			return false
		}
	}
	return true
}

// place sets or clears the board cells for tetromino z and manages placement/symmetry tracking.
func (s *Solver) place(z int, sX, sY int, put bool) {
	for _, p := range s.tetros[z] {
		x := sX + p.X
		y := sY + p.Y
		if put {
			s.board[x][y] = s.labels[z]
		} else {
			s.board[x][y] = '.'
		}
	}
	if put {
		s.placed[z] = true
		s.placedCount++
		if s.next[z] != -1 {
			s.sawIt[s.next[z]] = false
		}
	} else {
		s.placed[z] = false
		s.placedCount--
		if s.next[z] != -1 {
			s.sawIt[s.next[z]] = true
		}
	}
}

// backtrack recursively searches for a valid packing layout.
// It iterates through the board cells in raster order.
func (s *Solver) backtrack(start int, freeCells int) bool {
	// Base case: all tetrominoes have been successfully placed
	if s.placedCount == s.numTetros {
		return true
	}

	// Calculate the maximum allowed empty spaces on this board size
	maxFree := s.boardSize*s.boardSize - s.numTetros*4
	
	// If we've reached the end of the board or skipped too many cells, prune this branch
	if start >= s.boardSize*s.boardSize || freeCells > maxFree {
		return false
	}

	sX := start / s.boardSize
	sY := start % s.boardSize

	// If the current cell is already occupied, skip it
	if s.board[sX][sY] != '.' {
		return s.backtrack(start+1, freeCells)
	}

	// Option 1: Try to place any available tetromino starting at the current cell (sX, sY)
	for z := 0; z < s.numTetros; z++ {
		if !s.placed[z] && !s.sawIt[z] && s.canPlace(z, sX, sY) {
			s.place(z, sX, sY, true)
			if s.backtrack(start+1, freeCells) {
				return true
			}
			s.place(z, sX, sY, false) // Backtrack
		}
	}

	// Option 2: Leave the current cell empty.
	// We increment freeCells because we chose not to cover this empty cell.
	return s.backtrack(start+1, freeCells+1)
}

// Solve finds the smallest square board that can accommodate all parsed tetrominoes.
// It converts parsed coordinates into normalized Point structures and initializes the Solver.
func Solve(allCoords [][][]int, height int, supply int) [][]byte {
	var tetros [][]Point
	var labels []byte

	for _, coords := range allCoords {
		// The last element contains the single value with the character label (e.g. 65 for 'A')
		label := byte(coords[len(coords)-1][0])
		labels = append(labels, label)

		// Convert standard row/column indices to relative coordinates relative to the first block
		var points []Point
		found := false
		var firstX, firstY int

		for i := 0; i < len(coords)-1; i++ {
			row := coords[i]
			for _, col := range row {
				if !found {
					firstX = i
					firstY = col
					found = true
				}
				points = append(points, Point{X: i - firstX, Y: col - firstY})
			}
		}
		tetros = append(tetros, points)
	}

	numTetros := len(tetros)
	next := make([]int, numTetros)
	sawIt := make([]bool, numTetros)
	for i := 0; i < numTetros; i++ {
		next[i] = -1
		sawIt[i] = false
	}

	// Detect identical tetrominoes to break symmetries during search
	for i := 0; i < numTetros; i++ {
		for j := i + 1; j < numTetros; j++ {
			identical := true
			for k := 0; k < 4; k++ {
				if tetros[i][k].X != tetros[j][k].X || tetros[i][k].Y != tetros[j][k].Y {
					identical = false
					break
				}
			}
			if identical {
				next[i] = j
				sawIt[j] = true
				break
			}
		}
	}

	// Start with the minimum mathematically possible board size
	size := int(math.Ceil(math.Sqrt(float64(height)*4 + float64(supply))))

	for {
		solver := &Solver{
			board:       CreateTab(size),
			boardSize:   size,
			tetros:      tetros,
			labels:      labels,
			next:        next,
			placed:      make([]bool, numTetros),
			sawIt:       make([]bool, numTetros),
			numTetros:   numTetros,
			placedCount: 0,
		}
		copy(solver.sawIt, sawIt)

		if solver.backtrack(0, 0) {
			return solver.board
		}
		size++ // Increment board size directly on failure
	}
}

// Printer prints the board layout to standard output.
func Printer(finalTab [][]byte) {
	for i := 0; i < len(finalTab); i++ {
		for j := 0; j < len(finalTab[i]); j++ {
			fmt.Print(string(finalTab[i][j]))
		}
		fmt.Println()
	}
}
