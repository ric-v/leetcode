package main

import (
	"fmt"
	"sort"
)

/*
_____________
| 1 | 2 | 3 |
-------------
| 4 | 5 | 6 |
-------------
| 7 | 8 | 9 |
-------------
[ 1     2     4     7     5     3      6     8     9]

	00    01    10    20    11    02     12    21    22
*/
func findDiagonalOrder(mat [][]int) (result []int) {

	if len(mat) == 0 || len(mat[0]) == 0 {
		return
	}

	var N, M = len(mat), len(mat[0])
	var intermediate []int

	for i := 0; i < N+M-1; i++ {

		intermediate = []int{}

		var r, c int

		if i < M {
			r = 0
			c = i
		} else {
			r = i - M + 1
			c = M - 1
		}

		for r < N && c > -1 {
			fmt.Println(i, r, c)
			intermediate = append(intermediate, mat[r][c])
			r++
			c--
		}
		fmt.Println(intermediate)

		if i%2 == 0 {
			sort.Sort(sort.Reverse(sort.IntSlice(intermediate)))
		}

		result = append(result, intermediate...)
	}
	return
}

func main() {

	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
