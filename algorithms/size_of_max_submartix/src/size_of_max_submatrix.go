package MaxSubmatrixSize

func Size_of_max_submatrix(number [][]int) int {
	for i := 0; i < len(number); i++ {
		for j := 0; j < len(number[0]); j++ {
			if number[i][j] < 0 || number[i][j] > 1 {
				panic("element of array must be 0 or 1")
			}
		}
	}
	return getMaxSubmatrixSize(number)
}

func getMaxSubmatrixSize(number [][]int) int {
	width := len(number[0])
	height := len(number)
	row := make([]int, width, width)   //每行为底往上的连续1的个数
	col := make([]int, height, height) //每列为底，往右的连续1的个数
	res := make([]int, width, width)
	res[0] = 0
	for i := 0; i < height; i++ {
		col[0] = 0
		for j := 0; j < width; j++ {
			if j == 0 {
				if number[i][j] == 1 {
					row[j]++
					col[i]++
					if res[j] < row[j] {
						res[j] = row[j]
					}
				} else {
					row[j] = 0
					col[i] = 0
				}
			} else if i == 0 {
				if number[i][j] == 1 {
					row[j]++
					col[i]++
					if res[j] < col[i] {
						res[j] = col[i]
					}
				} else {
					row[i] = 0
					col[i] = 0
				}
			} else {
				if number[i][j] == 1 {
					row[j]++
					col[i]++
					if number[i][j-1] == 0 {
						if row[j] > res[j] {
							res[j] = row[j]
						} else if res[j-1] > res[j] {
							res[j] = res[j-1]
						}
					} else if number[i-1][j] == 0 {
						if col[i] > res[j] {
							res[j] = col[i]
						} else if res[j-1] > res[j] {
							res[j] = res[j-1]
						}
					} else {
						if col[i] > res[j] {
							res[j] = col[i]
						} else if res[j-1] > res[j] {
							res[j] = res[j-1]
						} else if row[j] > res[j] {
							res[j] = row[j]
						}
					}
				} else {
					row[j] = 0
					col[i] = 0
					if res[j-1] > res[j] {
						res[j] = res[j-1]
					}
				}
			}

		}
	}
	return res[width-1]
}
