package main

import (
	"fmt"
	"math"
)

// Create the tab that'll contain the tetrominoes
func CreateTab(heigth int, supply int) [][]byte {
	LenFinalTab := int(math.Ceil(math.Sqrt(float64(heigth)*4 + float64(supply))))
	finalTab := make([][]byte, LenFinalTab)
	for i := range finalTab {
		finalTab[i] = make([]byte, LenFinalTab)
		for j := range finalTab[i] {
			finalTab[i][j] = 46
		}
	}
	return finalTab
}

// Verify if tetrominoe can be placed on the tab above
func ItCan(finalTab [][]byte, indexs [][]int, x int, y int) bool {
	tmpy := y
	tmpx := x
	for i := range indexs[:len(indexs)-1] {
		y = tmpy
		incre := false
		for j := 0; j < len(indexs[i]); j++ {
			if !incre && i > 0 && indexs[i][j] != indexs[0][0] {
				if indexs[i][j]-1 == indexs[0][0] {
					y = y + 1
				} else if indexs[i][j]-2 == indexs[0][0] {
					y = y + 2
				} else if indexs[i][j]-3 == indexs[0][0] {
					y = y + 3
				} else if indexs[i][j]+1 == indexs[0][0] {
					y = y - 1
				} else if indexs[i][j]+2 == indexs[0][0] {
					y = y - 2
				} else if indexs[i][j]+3 == indexs[0][0] {
					y = y - 3
				}
			}
			if x >= len(finalTab) || y >= len(finalTab[x]) || y < 0 || finalTab[x][y] != 46 {
				return false
			}
			y++
			incre = true
		}
		x++
	}
	x = tmpx
	return true
}

// Place or Delete the tetrominoe on the tab
func PlaceOrDeL(finalTab [][]byte, indexs [][]int, x int, y int, place bool) {
	tmp := y
	tmpx := x
	for i := range indexs[:len(indexs)-1] {
		y = tmp
		incre := false
		for j := 0; j < len(indexs[i]); j++ {
			if !incre && i > 0 && indexs[i][j] != indexs[0][0] {
				if indexs[i][j]-1 == indexs[0][0] {
					y = y + 1
				} else if indexs[i][j]-2 == indexs[0][0] {
					y = y + 2
				} else if indexs[i][j]-3 == indexs[0][0] {
					y = y + 3
				} else if indexs[i][j]+1 == indexs[0][0] {
					y = y - 1
				} else if indexs[i][j]+2 == indexs[0][0] {
					y = y - 2
				} else if indexs[i][j]+3 == indexs[0][0] {
					y = y - 3
				}
			}
			if place {
				finalTab[x][y] = byte(indexs[len(indexs)-1][0])
			} else {
				finalTab[x][y] = 46
			}
			y++
			incre = true
		}
		x++
	}
	y = tmp
	x = tmpx
}

// That function will permit to try possibilities of resolutions
func Backtracking(finalTab [][]byte, Indexs [][][]int, n int) bool {
	if n == len(Indexs) {
		return true
	}
	for i := range finalTab {
		for j := range finalTab[i] {
			if ItCan(finalTab, Indexs[n], i, j) {
				PlaceOrDeL(finalTab, Indexs[n], i, j, true)
				if Backtracking(finalTab, Indexs, n+1) {
					return true
				}
				PlaceOrDeL(finalTab, Indexs[n], i, j, false)
			}
		}
	}
	return false
}

// This function will solve the problem with the function above and will adjust the tab size to tetrominoes
func Solve(Indexs [][][]int, heigth int, supply int) [][]byte {
	finalTab := CreateTab(heigth, supply)
	for !Backtracking(finalTab, Indexs, 0) {
		supply++
		finalTab = CreateTab(heigth, supply)
	}
	return finalTab
}

// Print the Result
func Printer(finalTab [][]byte) {
	for i := 0; i < len(finalTab); i++ {
		for j := 0; j < len(finalTab[i]); j++ {
			fmt.Print(string(finalTab[i][j]))
		}
		fmt.Println()
	}
}
