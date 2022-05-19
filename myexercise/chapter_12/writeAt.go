package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("writeAt.txt")
	if err != nil {
		panic(err)
	}
	var b bool
	fmt.Println(b)
	defer file.Close()
	file.WriteString("Golang中文社区——这里是多余")
	fmt.Printf("%U", '-')
	fmt.Println(len([]byte("Golang")))
	fmt.Println(len([]byte("Golang中文社区")))
	fmt.Println(len([]byte("Golang中文社区——")))
	n, err := file.WriteAt([]byte("Go语言中文网"), 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}
func isValidSudoku(board [][]byte) bool {
	var rows, cols [9][9]int
	var space [3][3][9]int
	for i, row := range board {
		for j, b := range row {
			if b == '.' {
				continue
			}
			index := b - '1'
			rows[i][index]++
			cols[j][index]++
			space[i/3][j/3][index]++
			if rows[i][index] > 1 || cols[j][index] > 1 || space[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}

func setZeroes(matrix [][]int) {
	rows := make([]bool, len(matrix))
	cols := make([]bool, len(matrix[0]))
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(cols); j++ {
			if rows[i] || cols[j] {
				continue
			}
			if matrix[i][j] == 0 {
				for si := 0; si < len(rows); si++ {
					for sj := 0; sj < len(cols); sj++ {
						matrix[si][sj] = 0
					}
				}
				rows[i] = true
				cols[j] = true
			}
		}
	}
}
