package leetcode

/*
- 867.Transpose Matrix
- Description:
Given a matrix A, return the transpose of A.

The transpose of a matrix is the matrix flipped over it's main diagonal, switching the row and column indices of the matrix.

Example 1:
Input: [[1,2,3],[4,5,6],[7,8,9]]
Output: [[1,4,7],[2,5,8],[3,6,9]]

Example 2:
Input: [[1,2,3],[4,5,6]]
Output: [[1,4],[2,5],[3,6]]
*/
func transpose(A [][]int) [][]int {
	rowLen := len(A)
	colLen := len(A[0])
	matrix := make([][]int, colLen)
	for i := 0; i < colLen; i++ {
		matrix[i] = make([]int, rowLen)
		for j := 0; j < rowLen; j++ {
			matrix[i][j] = A[j][i]
		}
	}

	return matrix
}
