<!--
 * @Descripttion: 
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2020-06-23 16:00:14
 * @LastEditors: Elon C
 * @LastEditTime: 2020-06-25 21:26:57
 * @FilePath: \GoPath\src\algorithms\array_and_matrix\print_matrix_specially\README.md
--> 

# “之”字形打印矩阵

## 要求：
给定一个矩阵matrix，按照“之”字形的方式打印这个矩阵，例如：

        1    2    3    4

        5    6    7    8

        9   10   11   12

“之”字形打印的结果为：1，2，5，9，6，3，4，7，10，11，8，12
【要求】额外空间复杂度为O (1)。
## 原理：
使用分圈处理的方式，在矩阵中用左上角的坐标(tR，tC)和右下角的坐标(dR，dC)就可以表示一个子矩阵。
分层旋转矩阵的最外圈矩阵