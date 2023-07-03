package main

import (
	"fmt"
	"os"
)

// Will put tetrominoes in a byte tab
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

// Verify if we have a correct tetrominoe
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

// Will get indexs of tetrominoes
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

// Put tetrominoes indexs in a tab
func TakeAllIndexs(tetrominoes [][][]byte) [][][]int {
	Indexs := [][][]int{}
	for i, char := 0, 65; i < len(tetrominoes); i, char = i+1, char+1 {
		Indexs = append(Indexs, GetIndexs(tetrominoes[i], char))
		VerifMinoes(Indexs[i])
	}
	return Indexs
}

func Error() {
	fmt.Println("ERROR")
	os.Exit(0)
}
