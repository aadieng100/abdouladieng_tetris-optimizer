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
		Indexs := [][][]int{}
		for i, char := 0, 65; i < len(tetrominoes); i, char = i+1, char+1 {
			indexs := GetIndexs(tetrominoes[i], char)
			Indexs = append(Indexs, indexs)
			VerifMinoes(Indexs[i])
		}
		finalTab := Solve(Indexs, len(Indexs), 0)
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

func Backtracking(finalTab [][]byte, Indexs [][][]int, n int) bool {
	if n == len(Indexs) {
		return true
	}
	for i := range finalTab {
		for j := range finalTab[i] {
			//indexs, itCan := FixIndexs(finalTab, Indexs[n], i, j, 65+n)
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

func Solve(Indexs [][][]int, heigth int, supply int) [][]byte {
	finalTab := CreateTab(heigth, supply)
	for !Backtracking(finalTab, Indexs, 0) {
		supply++
		finalTab = CreateTab(heigth, supply)
	}
	return finalTab
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
	for i := 0; i < len(doubleTab); i++ {
		for j := 0; j < len(doubleTab[i]); j++ {
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
