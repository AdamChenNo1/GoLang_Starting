/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2020-06-23 16:11:29
 * @LastEditors: Elon C
 * @LastEditTime: 2020-06-24 22:41:59
 * @FilePath: \GoPath\src\algorithms\array_and_matrix\spiral_order_print\spiral_order_print.go
 */
package matrix

import "fmt"

//spiral_order_print judges whether v is a int matrix and prints it spirally if it really is
func spiral_order_print(v [][]int) {
	if v == nil {
		panic("v should be a matrix while v is nil")
	}
	for i := 0; i < len(v); i++ {
		if v[i] == nil {
			panic(fmt.Sprintf("v should be a matrix while v[%d] is nil", i))
		}
	}
	spiralOrderPrint(v)
}

//spiralOrderPrint prints an int matrix spirally
func spiralOrderPrint(v [][]int) {
	length := len(v)
	width := len(v[0])
	tR, tC := 0, 0
	dR, dC := length-1, width-1
	for tR <= dR && tC <= dC {
		printEdge(v, tR, tC, dR, dC)
		tR++
		tC++
		dR--
		dC--
	}
}

//printEdge prints the edge of an int matrix
func printEdge(v [][]int, tR, tC, dR, dC int) {
	fmt.Printf("%d ", v[tR][tC])
	//1*1矩阵
	if dR == tR && tC == dC {
		return
	}

	for i := tC + 1; i <= dC; i++ {
		fmt.Printf("%d ", v[tR][i])
	}

	//矩阵v仅一行
	if tR == dR {
		return
	}

	for i := tR + 1; i <= dR; i++ {
		fmt.Printf("%d ", v[i][dC])
	}

	//矩阵v仅一列
	if tC == dC {
		return
	}

	for i := dC - 1; i >= tC; i-- {
		fmt.Printf("%d ", v[dR][i])
	}
	for i := dR - 1; i > tR; i-- {
		fmt.Printf("%d ", v[i][tC])
	}
}
