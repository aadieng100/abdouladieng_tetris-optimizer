package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

func Printer(finalTab [][]byte) {
	for i := 0; i < len(finalTab); i++ {
		for j := 0; j < len(finalTab[i]); j++ {
			fmt.Print(string(finalTab[i][j]))
		}
		fmt.Println()
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		file, err := ioutil.ReadFile(args[0])
		if err != nil {
			Error()
		}
		tetrominoes := TabMinoes(file)
		for i := 0; i < len(tetrominoes); i++ {
			indexs := GetIndexs(tetrominoes[i], 0)
			VerifMinoes(indexs)
			//ToAlfa(tetrominoes[i], char)
		}
		finalTab := Solve(tetrominoes, len(tetrominoes), 0)
		Printer(finalTab)
		return
	}
	fmt.Println("Only one argument is allowed.")
}

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

func PlaceOrDeL(finalTab [][]byte, indexs [][]int, x int, y int, place bool) {
	for i := 0; i < len(indexs)-1; i++ {
		for j := 0; j < len(indexs[i]); j++ {
			if place {
				a := indexs[len(indexs)-1][0]
				finalTab[x][indexs[i][j]] = byte(a)
				continue
			}
			finalTab[x][indexs[i][j]] = byte(46)
		}
		x++
	}
}

func Backtraking(finalTab [][]byte, tetrominoes [][][]byte, n int) bool {
	if n == len(tetrominoes) {
		return true
	}
	for i := 0; i < len(finalTab); i++ {
		for j := 0; j < len(finalTab[i]); j++ {
			indexs, itCan := FixIndexs(finalTab, tetrominoes[n], i, j, 65+n)
			if itCan {
				PlaceOrDeL(finalTab, indexs, i, j, true)
				Printer(finalTab)
				fmt.Println()
				if Backtraking(finalTab, tetrominoes, n+1) {
					return true
				}
				PlaceOrDeL(finalTab, indexs, i, j, false)
				Printer(finalTab)
				fmt.Println()
			}
		}
	}
	return false
}

func Solve(tetrominoes [][][]byte, heigth int, supply int) [][]byte {
	finalTab := CreateTab(heigth, supply)
	for !Backtraking(finalTab, tetrominoes, 0) {
		supply++
		finalTab = CreateTab(heigth, supply)
	}
	return finalTab
}

func FixIndexs(finalTab [][]byte, tetrominoe [][]byte, indexi int, indexj int, char int) ([][]int, bool) {
	x := indexi
	tab := GetIndexs(tetrominoe, char)
	tmp := [][]int{}
	for k := range tab {
		tabTmp := []int{}
		for l := 0; l < len(tab[k]); l++ {
			tabTmp = append(tabTmp, tab[k][l])
		}
		tmp = append(tmp, tabTmp)
	}
	itCan := true
	for i := 0; i < len(tab)-1; i++ {
		for j := 0; j < len(tab[i]); j++ {
			if tab[i][j] == tmp[0][0] {
				tab[i][j] = indexj
			} else if tab[i][j]-1 == tmp[0][0] {
				tab[i][j] = indexj + 1
			} else if tab[i][j]-2 == tmp[0][0] {
				tab[i][j] = indexj + 2
			} else if tab[i][j]-3 == tmp[0][0] {
				tab[i][j] = indexj + 3
			} else if tab[i][j]+1 == tmp[0][0] {
				tab[i][j] = indexj - 1
			} else if tab[i][j]+2 == tmp[0][0] {
				tab[i][j] = indexj - 2
			} else if tab[i][j]+3 == tmp[0][0] {
				tab[i][j] = indexj - 3
			}
			if tab[i][j] < 0 || tab[i][j] > len(finalTab)-1 {
				tab = tmp
				return tab, false
			}
			current := tab[i][j]
			if indexi+len(tab) > len(finalTab) ||
				indexj+len(tab[i]) > len(finalTab[x]) ||
				finalTab[x][current] != byte(46) {
				tab = tmp
				return tab, false
			}
		}
		x++
	}
	return tab, itCan
}

func VerifMinoes(tab [][]int) {
	for i := 0; i < len(tab)-1; i++ {
		if i+1 < len(tab)-1 {
			ok := false
			for j := 0; j < len(tab[i]); j++ {
				if j+1 < len(tab[i]) && tab[i][j+1] != tab[i][j]+1 {
					Error()
				}
				for k := 0; k < len(tab[i+1]); k++ {
					if tab[i][j] == tab[i+1][k] {
						ok = true
					}
				}
			}
			if !ok {
				Error()
			}
		}
	}
}

func GetIndexs(doubleTab [][]byte, char int) [][]int {
	tab := []int{}
	twoTab := [][]int{}
	diez := 0
	dot := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if doubleTab[i][j] == 35 {
				diez++
				tab = append(tab, j)
			}
			if doubleTab[i][j] == 46 {
				dot++
			}
		}
		if diez > 0 && diez < 4 && len(tab) == 0 {
			Error()
		}
		if len(tab) != 0 {
			twoTab = append(twoTab, tab)
			tab = nil
		}
	}
	if diez != 4 || diez*3 != dot {
		Error()
	}
	tab = append(tab, char)
	twoTab = append(twoTab, tab)
	return twoTab
}

// func ToAlfa(doubleTab [][]byte, char int) {
// 	for i := 0; i < len(doubleTab); i++ {
// 		for j := 0; j < len(doubleTab[i]); j++ {
// 			if doubleTab[i][j] == 35 {
// 				doubleTab[i][j] = byte(char)
// 			}
// 		}
// 	}
// }

func TabMinoes(file []byte) [][][]byte {
	tab := []byte{}
	doubleTab := [][]byte{}
	tripleTab := [][][]byte{}
	for i := 0; i < len(file); i++ {
		if i == len(file)-1 {
			if file[i] != 10 {
				tab = append(tab, file[i])
			}
			if len(tab) == 4 {
				doubleTab = append(doubleTab, tab)
				tripleTab = append(tripleTab, doubleTab)
				continue
			}
		}
		if i+1 < len(file) && file[i] == 10 {
			if file[i+1] != 10 {
				if len(tab) == 4 {
					doubleTab = append(doubleTab, tab)
					tab = nil
					continue
				}
				Error()
			} else {
				if len(tab) == 4 {
					doubleTab = append(doubleTab, tab)
					tab = nil
				}
				if len(doubleTab) == 4 {
					tripleTab = append(tripleTab, doubleTab)
					doubleTab = nil
					i++
					continue
				}
				Error()
			}
		}
		tab = append(tab, file[i])
	}
	return tripleTab
}

func Error() {
	fmt.Println("ERROR")
	os.Exit(0)
}
