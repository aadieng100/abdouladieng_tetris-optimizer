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
		finalTab := [][]byte{
			{46, 46, 46, 46},
			{46, 46, 46, 46},
			{46, 46, 46, 46},
			{46, 46, 46, 46}}
		file, err := ioutil.ReadFile(args[0])
		if err != nil {
			Error()
		}
		tetrisFile := string(file)
		tetrominoes := LineBetween(tetrisFile)
		// VerifyCar(tetrominoes)
		VerifyMinoes(tetrominoes)
		tetrominoes = ToAlpha(tetrominoes)
		tetrominoes = DelDotColMinoes(tetrominoes)
		tetrominoes = DelDotLine(tetrominoes)
		// finalTab = Solve(tetrominoes, finalTab)
		for i := range finalTab {
			fmt.Println(finalTab[i])
		}
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

// func VerifyCar(tetrominoes [][]string) {
// 	for i := range tetrominoes {
// 		diez := 0
// 		dot := 0
// 		for j := range tetrominoes[i] {
// 			line := tetrominoes[i][j]
// 			for k := range line {
// 				if len(line) != 4 {
// 					Error()
// 				}
// 				if string(line[k]) == "#" {
// 					diez++
// 				}
// 				if string(line[k]) == "." {
// 					dot++
// 				}
// 			}
// 		}
// 		if diez*3 != dot {
// 			Error()
// 		}
// 	}
// }

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

func Solve(tetrominoes [][]string, finalTab [][]byte) [][]byte {
	for i := 0; i < len(tetrominoes); i++ {
		finalTab = TryMinoe(tetrominoes[i], finalTab)
	}
	return finalTab
}

func TryMinoe(tetrominoe []string, finalTab [][]byte) [][]byte {
	fmt.Println(MinoDiezIndex(tetrominoe))
	for i := 0; i < len(finalTab); i++ {
		k := 0
		for j := 0; j < len(finalTab[i]); j++ {
			if finalTab[i][j] == 46 && i < len(tetrominoe) && k < len(tetrominoe[i]) {
				finalTab[i][j] = tetrominoe[i][k]
				k++
			}
		}
	}
	return finalTab
}

func MinoDiezIndex(tetrominoe []string) [][]int {
	doubleTabIndex := [][]int{}
	numberOfDiez := 0
	numberOfDot := 0
	for i := 0; i < len(tetrominoe); i++ {
		tabIndex := []int{}
		for j := 0; j < len(tetrominoe[i]); j++ {
			if tetrominoe[i][j] == 35 {
				numberOfDiez++
				tabIndex = append(tabIndex, j)
			}
			if tetrominoe[i][j] == 46 {
				numberOfDot++
			}
		}
		if len(tabIndex) == 0 && numberOfDiez > 0 && numberOfDiez != 4 {
			Error()
		}
		if len(tabIndex) != 0 {
			doubleTabIndex = append(doubleTabIndex, tabIndex)
		}
	}
	if numberOfDot != (numberOfDiez * 3) {
		Error()
	}
	return doubleTabIndex
}

func VerifyMinoes(tetrominoes [][]string) {
	for i := range tetrominoes {
		Indexs := MinoDiezIndex(tetrominoes[i])
		VerifyIndexs(Indexs)
	}
}

func VerifyIndexs(Indexs [][]int) {
	if len(Indexs) == 4 {
		for i := range Indexs {
			for j := range Indexs[i] {
				if Indexs[i][j] != Indexs[0][0] {
					Error()
				}
			}
		}
	} else if len(Indexs) == 3 {
		count := 0
		for i := range Indexs[0] {
			for j := range Indexs[1] {
				if Indexs[0][i] == Indexs[1][j] {
					count++
				}
			}
		}
		if count == 1 {
			for i := range Indexs[1] {
				for j := range Indexs[2] {
					if Indexs[1][i] == Indexs[2][j] {
						count++
					}
				}
			}
		}
		if count != 2 {
			Error()
		}
	} else if len(Indexs) == 2 {
		if len(Indexs[0]) == len(Indexs[1]) && (Indexs[0][0] != Indexs[1][0] || Indexs[0][1] != Indexs[1][1]) {
			count := 0
			for i := 0; i < len(Indexs[0]); i++ {
				for j := 0; j < len(Indexs[1]); j++ {
					if Indexs[0][i] == Indexs[1][j] {
						count++
					}
				}
			}
			if count != 1 {
				Error()
			}
		}
	}
}
