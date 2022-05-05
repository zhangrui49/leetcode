package g118

/**

给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和

*/

func generate(numRows int) [][]int {
	r := make([][]int, numRows)
	if numRows == 1 {
		r[0] = []int{1}
		return r
	}
	if numRows == 2 {
		r[0] = []int{1}
		r[1] = []int{1, 1}
		return r
	}
	r[0] = []int{1}
	r[1] = []int{1, 1}
	for i := 2; i < numRows; i++ {
		count := i + 1
		r[i] = make([]int, count)
		r[i][0] = 1
		r[i][i] = 1
		for j := 1; j < count-1; j++ {
			r[i][j] = r[i-1][j-1] + r[i-1][j]
		}

	}

	return r
}
