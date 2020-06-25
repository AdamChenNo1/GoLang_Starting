<!--
 * @Descripttion: 
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2020-06-23 16:00:14
 * @LastEditors: Elon C
 * @LastEditTime: 2020-06-25 19:54:45
 * @FilePath: \GoPath\src\algorithms\array_and_matrix\rotate_matrix\README.md
--> 

# 将正方形矩阵顺时针转动90°

## 要求：
给定一个N×N 的矩阵matrix，把这个矩阵调整成顺时针转动90°后的形式。

例如：

        1   2   3   4

        5   6   7   8

        9   10   11   12

       13   14   15   16


顺时针转动90°后为：

        13    9    5   1

        14   10   6   2

        15   11   7   3

        16   12   8   4

【要求】额外空间复杂度为O (1)。
## 原理：
使用分圈处理的方式，在矩阵中用左上角的坐标(tR，tC)和右下角的坐标(dR，dC)就可以表示一个子矩阵。
分层旋转矩阵的最外圈矩阵