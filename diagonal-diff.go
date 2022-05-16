package main

func diagonalDifference(a [][]int32) int32 {
	// Write your code here
	var r int32
	var diag1 int32
	var diag2 int32

	l := len(a)

	for i := 0; i < l; i++ {
		diag1 += a[i][i]
		diag2 += a[l-(i+1)][i]
	}

	if diag1 > diag2 {
		r = diag1 - diag2
	} else {
		r = diag2 - diag1
	}

	return r
}
