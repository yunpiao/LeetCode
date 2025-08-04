/*
 * @lc app=leetcode.cn id=74 lang=golang
 * @lcpr version=30104
 *
 * [74] 搜索二维矩阵
 */
package main

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	// 先从第一列二分查找目标行
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if target < matrix[0][0] || target > matrix[m-1][n-1] {
		return false
	}
	// 再行目标行找到确定值
	left, right := 0, m-1
	mid := 0
	for left <= right {
		mid = (right-left)/2 + left
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	findIndex := right

	// mid 为目标行
	left, right = 0, n-1
	for left <= right {
		mid := (right-left)>>1 + left
		if matrix[findIndex][mid] == target {
			return true
		} else if matrix[findIndex][mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false

}

// @lc code=end

/*
// @lcpr case=start
// [[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n3\n
// @lcpr case=end

// @lcpr case=start
// [[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n13\n
// @lcpr case=end

*/
