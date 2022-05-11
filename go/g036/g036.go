package main

func IsValidSudoku(board [][]byte) bool {
	row := [9][9]bool{{}}
	col := [9][9]bool{{}}

	for i := 0; i < 9; i++ {
		cell := [9]bool{}
		for j := 0; j < 9; j++ {
			vr := board[i][j]
			vc := board[j][i]
			if vr != '.' {
				if row[i][int(vr-'1')] == true {
					return false
				}
				row[i][int(vr-'1')] = true
			}
			if vc != '.' {
				if col[i][int(vc-'1')] == true {
					return false
				}
				col[i][int(vc-'1')] = true
			}
			cellX := 3*(i/3) + j/3
			cellY := 3*(i%3) + j%3
			cellV := board[cellX][cellY]
			if cellV != '.' {
				if cell[int(cellV-'1')] == true {
					return false
				}
				cell[int(cellV-'1')] = true
			}
		}
	}
	return true
}
