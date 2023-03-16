package main

import (
	"fmt"
)

func findDiagonalOrder(mat [][]int) (result []int) {

	isWalkUp := true
	height, width := len(mat), len(mat[0])
	result = make([]int, 0, width*height)
	i, j := 0, 0

	for len(result) < height*width {
		result = append(result, mat[i][j])

		switch {
		case isWalkUp:
			if i == 0 && j != width-1 {
				isWalkUp = !isWalkUp
				j++
			} else if j == width-1 {
				isWalkUp = !isWalkUp
				i++
			} else {
				i--
				j++
			}

		case !isWalkUp:
			if j == 0 && i != height-1 {
				isWalkUp = !isWalkUp
				i++
			} else if i == height-1 {
				isWalkUp = !isWalkUp
				j++
			} else {
				i++
				j--
			}
		}
	}
	return
}

func main() {

	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
