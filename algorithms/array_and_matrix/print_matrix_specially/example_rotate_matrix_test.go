/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2020-06-24 21:45:13
 * @LastEditors: Elon C
 * @LastEditTime: 2020-06-25 22:51:06
 * @FilePath: \GoPath\src\algorithms\array_and_matrix\print_matrix_specially\example_rotate_matrix_test.go
 */
package matrix

import "fmt"

func Example_spiral_order_print() {
	datas := [][][]int{
		{{1}},
		{{1, 2}, {3, 4}},
		{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
	}
	for _, e := range datas {
		print_matrix_specially(e)
		fmt.Println()
	}
	//Output:
	//1
	//1 2 3 4
	//1 2 4 7 5 3 6 8 9
	//1 2 5 9 6 3 4 7 10 11 8 12

}

func printMatrix(v [][]int) {
	for _, l := range v {
		for _, e := range l {
			fmt.Printf("%d ", e)
		}
	}
}
