package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		file, err := ioutil.ReadFile(args[0])
		if err != nil {
			Error()
		}
		tetrisFile := string(file)
		tetrominoes := LineBetween(tetrisFile)
		VerifyCar(tetrominoes)
		ValidTetrominoes(tetrominoes)
		tetrominoes = ToAlpha(tetrominoes)
		tetrominoes = DelDotColMinoes(tetrominoes)
		tetrominoes = DelDotLine(tetrominoes)
		// Solve(tetrominoes)
		for i := range tetrominoes {
			for j := range tetrominoes[i] {
				fmt.Println(tetrominoes[i][j])
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Please add an argument and only one.")
	}
}

func Error() {
	fmt.Println("ERROR")
	os.Exit(0)
}

func LineBetween(s string) [][]string {
	DoubleLineSlash := strings.Split(s, "\n\n")
	tabTetrominoes := [][]string{}
	for i := range DoubleLineSlash {
		tetrominoe := strings.Split(DoubleLineSlash[i], "\n")
		if i == len(DoubleLineSlash)-1 {
			tetrominoe = tetrominoe[:len(tetrominoe)-1]
		}
		if len(tetrominoe) != 4 {
			Error()
		}
		tabTetrominoes = append(tabTetrominoes, tetrominoe)
	}
	return tabTetrominoes
}

func VerifyCar(tetrominoes [][]string) {
	for i := range tetrominoes {
		diez := 0
		dot := 0
		for j := range tetrominoes[i] {
			line := tetrominoes[i][j]
			for k := range line {
				if len(line) != 4 {
					Error()
				}
				if string(line[k]) == "#" {
					diez++
				}
				if string(line[k]) == "." {
					dot++
				}
			}
		}
		if diez*3 != dot {
			Error()
		}
	}
}

func ValidTetrominoes(tetrominoes [][]string) bool {
	for i := range tetrominoes {
		ValidTetrominoe(tetrominoes[i])
	}
	return true
}

func ValidTetrominoe(tetrominoe []string) bool {
	for i := range tetrominoe {
		line := tetrominoe[i]
		for j := range line {
			if line == "####" {
				return true
			}
			if line == ".###" || line == "###." {
				indexs := DiezIndex(line)
				TreeDiez(tetrominoe, i, indexs)
				return true
			}
			if line == "##.." || line == ".##." || line == "..##" {
				indexs := DiezIndex(line)
				TwoDiez(tetrominoe, i, indexs)
				return true
			}
			if string(line[j]) == "#" && line == "#..." ||
				line == ".#.." || line == "..#." || line == "...#" {
				OneDiez(tetrominoe, line, j, i)
				return true
			}
		}
	}
	return false
}

func DiezIndex(line string) []int {
	indexTab := []int{}
	for i := range line {
		if string(line[i]) == "#" {
			indexTab = append(indexTab, i)
		}
	}
	return indexTab
}

func TreeDiez(tetrominoe []string, index int, indexs []int) {
	for i := 0; i < len(indexs); i++ {
		if string(tetrominoe[index+1][indexs[i]]) == "#" {
			return
		}
	}
	Error()
}

func TwoDiez(tetrominoe []string, index int, indexs []int) bool {
	nextIndex := DiezIndex(tetrominoe[index+1])
	switch len(nextIndex) {
	case 2:
		if tetrominoe[index+1] == tetrominoe[index] {
			return true
		} else {
			if nextIndex[1] == nextIndex[0]+1 {
				for i := range nextIndex {
					for j := range indexs {
						if nextIndex[i] == indexs[j] {
							return true
						}
					}
				}
			}
		}
	case 1:
		for i := range indexs {
			if indexs[i] == nextIndex[0] && tetrominoe[index+1] == tetrominoe[index+2] {
				return true
			}
		}
	}
	Error()
	return false
}

func OneDiez(tetrominoe []string, line string, index int, indTetro int) bool {
	indexs := DiezIndex(tetrominoe[indTetro+1])
	switch len(indexs) {
	case 3:
		for i := range indexs {
			if indexs[i] == index {
				return true
			}
		}
	case 2:
		nextIndexs := DiezIndex(tetrominoe[indTetro+2])
		for i := range indexs {
			if indexs[i] == index {
				for j := range indexs {
					if indexs[j] == nextIndexs[0] {
						return true
					}
				}
			}
		}
	case 1:
		nextIndexs := DiezIndex(tetrominoe[indTetro+2])
		if tetrominoe[indTetro+1] == line {
			if len(nextIndexs) == 2 {
				for i := range nextIndexs {
					if nextIndexs[i] == index {
						return true
					}
				}
			}
			if len(nextIndexs) == 1 && line == tetrominoe[nextIndexs[0]] &&
				line == tetrominoe[indTetro+3] {
				return true
			}
		}
	}
	Error()
	return true
}

func ToAlpha(tab [][]string) [][]string {
	tetrominoes := [][]string{}
	tabTetro := []string{}
	for i, char := 0, 'A'; i < len(tab); i++ {
		tetrominoe := strings.Join(tab[i], "\\n")
		str := ""
		for _, car := range tetrominoe {
			if string(car) == "#" {
				car = char
			}
			str += string(car)
		}
		tabTetro = strings.Split(str, "\\n")
		tetrominoes = append(tetrominoes, tabTetro)
		tabTetro = nil
		char++
		if char > 'Z' {
			Error()
		}
	}
	return tetrominoes
}

func DelDotColMinoes(tetrominoes [][]string) [][]string {
	minoesWithoutCol := [][]string{}
	for i := 0; i < len(tetrominoes); i++ {
		tab := DetectDot(tetrominoes[i])
		minoesWithoutCol = append(minoesWithoutCol, tab)
	}
	return minoesWithoutCol
}

func DetectDot(tetrominoe []string) []string {
	for i := 0; i < len(tetrominoe); i++ {
		for j := 0; j < len(tetrominoe[i]); j++ {
			if string(tetrominoe[0][j]) == "." &&
				string(tetrominoe[1][j]) == "." &&
				string(tetrominoe[2][j]) == "." &&
				string(tetrominoe[3][j]) == "." {
				tetrominoe = DelDotColMinoe(tetrominoe, j)
			}
		}
	}
	return tetrominoe
}

func DelDotColMinoe(tetrominoe []string, index int) []string {
	tabDelDot := []string{}
	for i := 0; i < len(tetrominoe); i++ {
		str := ""
		for j := 0; j < len(tetrominoe[i]); j++ {
			if j != index {
				str += string(tetrominoe[i][j])
			}
		}
		tabDelDot = append(tabDelDot, str)
	}
	return tabDelDot
}

func DelDotLine(tetrominoes [][]string) [][]string {
	finalMinoes := [][]string{}
	for i := 0; i < len(tetrominoes); i++ {
		DelLineDot := []string{}
		for j := 0; j < len(tetrominoes[i]); j++ {
			if string(tetrominoes[i][j]) == ".." ||
				string(tetrominoes[i][j]) == "..." ||
				string(tetrominoes[i][j]) == "...." {
				continue
			}
			DelLineDot = append(DelLineDot, tetrominoes[i][j])
		}
		finalMinoes = append(finalMinoes, DelLineDot)
	}
	return finalMinoes
}

// func Solve(tetrominoes [][]string) string {
// 	finalTab := [4][4]string{
// 		{".", ".", ".", "."},
// 		{".", ".", ".", "."},
// 		{".", ".", ".", "."},
// 		{".", ".", ".", "."}}
// 	for i := 0 ;
// 	fmt.Println(finalTab)
// 	Error()
// 	return ""
// }
// tmp := []string{}
// for i := 0 ; i < len(tetrominoes) ; i++ {

// }
// if len(tetrominoes[i]) == 4 && len(tetrominoes[0]) != 4 {
// 	tmp := tetrominoes[i]
// 	tetrominoes[i] = tetrominoes[0]
// 	tetrominoes[0] = tmp
// }

func SortMinoes(tetrominoes [][]string) [][]string {
	for i := range tetrominoes {
		for j := range tetrominoes[i] {
			if len(tetrominoes[j]) > len(tetrominoes[i]) {
				tetrominoes[i], tetrominoes[j] = tetrominoes[j], tetrominoes[i]
			}
		}
	}
	return tetrominoes
}
