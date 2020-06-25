/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2020-06-23 16:11:29
 * @LastEditors: Elon C
 * @LastEditTime: 2020-06-25 21:24:13
 * @FilePath: \GoPath\src\algorithms\array_and_matrix\rotate_matrix\rotate_matrix.go
 */
package matrix

import "fmt"

//rotate_matrix judges whether v is a square matrix and rotates it in clock-wise
func rotate_matrix(v *[][]int) {
	p := *v
	if p == nil {
		panic("v should be a matrix while v is nil")
	}
	if len(p) != len(p[0]) {
		panic("v should be a matrix")
	}
	for i := 0; i < len(p); i++ {
		if p[i] == nil {
			panic(fmt.Sprintf("v should be a matrix while v[%d] is nil", i))
		}
	}
	rotate(v)
}

//rotate rotates a square matrix in clock-wise
func rotate(v *[][]int) {
	length := len(*v)
	width := len((*v)[0])
	tR, tC := 0, 0
	dR, dC := length-1, width-1
	for tR <= dR && tC <= dC {
		rotateEdge(v, tR, tC, dR, dC)
		tR++
		tC++
		dR--
		dC--
	}
}

//rotateEdge rotates the edges of a square matrix
func rotateEdge(v *[][]int, tR, tC, dR, dC int) {
	for i := 0; i < dC-tC; i++ {
		p := *v
		temp := p[tR][tC+i]
		p[tR][tC+i] = p[dR-i][tC]
		p[dR-i][tC] = p[dR][dC-i]
		p[dR][dC-i] = p[tR+i][dC]
		p[tR+i][dC] = temp
	}
}
