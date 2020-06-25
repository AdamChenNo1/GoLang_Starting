/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2020-06-23 16:11:29
 * @LastEditors: Elon C
 * @LastEditTime: 2020-06-25 22:53:38
 * @FilePath: \GoPath\src\algorithms\array_and_matrix\print_matrix_specially\print_matrix_specially.go
 */
package matrix

import "fmt"

//print_matrix_specially judges whether v is a matrix and prints it specially
func print_matrix_specially(v [][]int) {
	if v == nil {
		panic("v should be a matrix while v is nil")
	}

	for i := 0; i < len(v); i++ {
		if v[i] == nil {
			panic(fmt.Sprintf("v should be a matrix while v[%d] is nil", i))
		}
	}
	printMatrixSpecially(v)
}

//printMatrixSpecially prints a matrix specially
func printMatrixSpecially(v [][]int) {
	lengthM := len(v) - 1
	widthM := len(v[0]) - 1
	for s := 0; s <= lengthM+widthM; s++ {
		end := s
		if s > widthM {
			end = widthM
		}
		start := 0
		if s-start > lengthM {
			start = s - lengthM
		}
		if s%2 == 0 {
			for i := start; i <= end; i++ {
				fmt.Printf("%d ", v[s-i][i])
			}
		} else {
			for i := end; i >= start; i-- {
				fmt.Printf("%d ", v[s-i][i])
			}
		}

	}
}
