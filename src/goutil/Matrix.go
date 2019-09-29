package goutil

func MakeIntMatrix(shape [2]int) [][]int {
	var mat = make([][]int, shape[0])
	for i := 0; i < shape[0]; i++ {
		mat[i] = make([]int, shape[1])
	}
	return mat
}
func CopyMatrix(src [][]int) [][]int {
	var mat = make([][]int, len(src))
	for i := 0; i < len(src); i++ {
		mat[i] = make([]int, len(src))
		copy(mat[i], src[i])
	}
	return mat
}

func Equal(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) > 0 {
		if len(a[0]) != len(b[0]) {
			return false
		}
	}

	hashA := 0
	hashB := 0
	for i := 0; i < len(a); i++ {
		step := uint(0)
		hashA = 0
		hashB = 0
		for j := 0; j < len(a[i]); j++ {
			hashA += (a[i][j] + 1) << step
			hashB += (b[i][j] + 1) << step
			step += 2
		}
		if hashA != hashB {
			return false
		}
	}
	return true
}
