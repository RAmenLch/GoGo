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
	for i := range a {
		for j := range a {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
