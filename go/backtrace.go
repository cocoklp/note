package main

import (
	"fmt"
)

/*
解空间树：
	问题的解一般使用解空间树的方式来组织，树的根节点位于第一层，表示搜索的最初状态，依次向下排列
解空间树的动态搜索：
	搜索至树中任一节点时，先判断对应的部分是否满足约束条件，不满足则跳过（剪枝），否则进入子树继续搜索
回溯：
	通过递归返回实现回溯
*/
/*
https://segmentfault.com/a/1190000003733325?utm_source=tag-newest
https://www.jianshu.com/p/dd3c3f3e84c0
https://blog.csdn.net/sinat_27908213/article/details/80599460
https://segmentfault.com/a/1190000018771841
https://www.cnblogs.com/king-lps/p/10748535.html
https://leetcode.wang/leetCode-46-Permutations.html

*/
/*
	N皇后问题
*/

func main() {
	solveNQueens(4)
}
func solveNQueens(n int) [][]string {
	queen := make([][]string, n)
	que := make([]string, n)
	for i := 0; i < n; i++ {
		que[i] = "."
	}
	for i := 0; i < n; i++ {
		queen[i] = que
	}
	backqueen(0, n, queen)
	fmt.Println(queen)
	return queen
}

func backqueen(col, n int, queen [][]string) {
	fmt.Println(col)
	if col == n {
		//fmt.Println(queen)
		return
	}
	for i := 0; i < n; i++ {
		if valid(i, col, n, queen) {
			queen[i][col] = "Q"
			backqueen(col+1, n, queen)
			queen[i][col] = "."
		}
	}
}

func valid(row, col, n int, queen [][]string) bool {
	for i := 0; i < n; i++ {
		if queen[row][i] == "Q" && i != col {
			return false
		}
	}
	for i := 0; i < n; i++ {
		if queen[i][col] == "Q" && i != row {
			return false
		}
	}

	// 左下角
	for i, j := row+1, col-1; i < n && j >= 0; i, j = i+1, j-1 {
		if queen[i][j] == "Q" {
			return false
		}
	}
	// 左上角
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if queen[i][j] == "Q" {
			return false
		}
	}
	// 右上角
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if queen[i][j] == "Q" {
			return false
		}
	}
	// 右下角
	for i, j := row+1, col+1; i < n && j < n; i, j = i+1, j+1 {
		if queen[i][j] == "Q" {
			return false
		}
	}
	return true
}
