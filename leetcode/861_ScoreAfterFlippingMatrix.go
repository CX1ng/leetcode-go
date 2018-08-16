package leetcode

func matrixScore(A [][]int) int {
	rowLength := len(A)
	colLength := len(A[0])
	// flipping row
	for i := 0; i < rowLength; i++ { // row
		if A[i][0] == 0 {
			for j := 0; j < colLength; j++ { // col
				A[i][j] = 1 - A[i][j]
			}
		}
	}
	// flipping col
	for i := 1; i < colLength; i++ { // col
		oneCnt := 0
		for j := 0; j < rowLength; j++ { // row
			if A[j][i] == 1 {
				oneCnt++
			}
		}
		if oneCnt <= rowLength/2 {
			for j := 0; j < rowLength; j++ {
				A[j][i] = 1 - A[j][i]
			}
		}
	}
	// calculate sum
	var sum int
	for i := 0; i < rowLength; i++ {
		for j := 0; j < colLength; j++ {
			if A[i][j] == 1 {
				sum += (1 << uint(colLength-j-1))
			}
		}
	}
	return sum
}
