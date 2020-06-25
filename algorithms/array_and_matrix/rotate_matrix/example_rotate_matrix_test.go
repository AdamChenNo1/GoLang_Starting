/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2020-06-24 21:45:13
 * @LastEditors: Elon C
 * @LastEditTime: 2020-06-25 21:10:39
 * @FilePath: \GoPath\src\algorithms\array_and_matrix\rotate_matrix\example_rotate_matrix_test.go
 */
package matrix

import "fmt"

func Example_spiral_order_print() {
	datas := [][][]int{
		{{1}},
		{{1, 2}, {3, 4}},
		{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
	}
	for _, e := range datas {
		rotate_matrix(&e)
		printMatrix(e)
		fmt.Println()
	}
	//Output:
	//1
	//3 1 4 2
	//7 4 1 8 5 2 9 6 3

}

func printMatrix(v [][]int) {
	for _, l := range v {
		for _, e := range l {
			fmt.Printf("%d ", e)
		}
	}
}
