package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		file, err := ioutil.ReadFile(args[0])
		if err != nil {
			Error()
		}
		tetrominoes := TabMinoes(file)
		Indexs := TakeAllIndexs(tetrominoes)
		finalTab := Solve(Indexs, len(Indexs), 0)
		Printer(finalTab)
		return
	}
	fmt.Println("Only one argument is allowed.")
}
