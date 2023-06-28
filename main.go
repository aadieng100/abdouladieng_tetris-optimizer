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

func PlaceOrDeL(finalTab [][]byte, indexs [][]int, x int, y int, place bool) {
	for i := range indexs[:len(indexs)-1] {
		for j := range indexs[i] {
			if place {
				finalTab[x][indexs[i][j]] = byte(indexs[len(indexs)-1][0])
			} else {
				finalTab[x][indexs[i][j]] = 46
			}
		}
		x++
	}
}

func Backtracking(finalTab [][]byte, Indexs [][][]int, n int) bool {
	if n == len(Indexs) {
		return true
	}
	for i := range finalTab {
		for j := range finalTab[i] {
			indexs, itCan := FixIndexs(finalTab, Indexs[n], i, j, 65+n)
			if itCan {
				PlaceOrDeL(finalTab, indexs, i, j, true)
				if Backtracking(finalTab, Indexs, n+1) {
					return true
				}
				PlaceOrDeL(finalTab, indexs, i, j, false)
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

func FixIndexs(finalTab [][]byte, Indexs [][]int, indexi int, indexj int, char int) ([][]int, bool) {
	tmp := make([][]int, len(Indexs))
	for i := range Indexs {
		tmp[i] = make([]int, len(Indexs[i]))
		copy(tmp[i], Indexs[i])
	}
	itCan := true
	for i := range tmp[:len(tmp)-1] {
		for j := range tmp[i] {
			offset := tmp[i][j] - Indexs[0][0]
			tmp[i][j] = indexj + offset
			if tmp[i][j] < 0 || tmp[i][j] > len(finalTab)-1 {
				return Indexs, false
			}
			if indexi >= len(finalTab) || finalTab[indexi][tmp[i][j]] != 46 {
				return Indexs, false
			}
		}
		indexi++
	}
	return tmp, itCan
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
