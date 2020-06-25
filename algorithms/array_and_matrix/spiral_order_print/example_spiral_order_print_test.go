/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2020-06-24 21:45:13
 * @LastEditors: Elon C
 * @LastEditTime: 2020-06-24 22:56:31
 * @FilePath: \GoPath\src\algorithms\array_and_matrix\spiral_order_print\example_spiral_order_print_test.go
 */
package matrix

import "fmt"

func Example_spiral_order_print() {
	datas := [][][]int{
		{{1}},
		{{1, 2}},
		//{{1}, {2}},
		//{{1, 2}, {3, 4}},
		//{{1, 2}, {3, 4}, {5, 6}},
		//{{1, 2, 3}, {4, 5, 6}},
		//{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
	}
	for i := 0; i < len(datas)-1; i++ {
		spiral_order_print(datas[i])
		fmt.Println()
	}
	spiral_order_print(datas[len(datas)-1])

	//Output:
	//1
	//1 2

}
