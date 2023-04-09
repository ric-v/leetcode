package diagonaltraverse

// DiagonalOrder returns the elements of the matrix in diagonal order.
//
// Time complexity: O(n)
// Space complexity: O(n)
func DiagonalOrder(mat [][]int) (result []int) {

	isWalkUp := true                       // isWalkUp indicates whether we are walking up the matrix or not
	height, width := len(mat), len(mat[0]) // height and width of matrix
	result = make([]int, 0, width*height)  // result is the resultant array of integers
	i, j := 0, 0                           // i and j are the row and column indices of the matrix

	for len(result) < height*width { // iterate until we have appended all the elements of the matrix
		result = append(result, mat[i][j]) // append the element to the result array
		// the following switch statement is used to update the row and column indices based on the direction of walk
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
