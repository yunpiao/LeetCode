/*
 * @lc app=leetcode.cn id=240 lang=golang
 * @lcpr version=30104
 *
 * [240] 搜索二维矩阵 II
 */

// @lc code=start
// 技巧从左下角 和 右上角 开始搜索
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	im, in := m-1, 0
	for im >= 0 && in < n {
		if matrix[im][in] == target {
			return true
		} else if matrix[im][in] > target {
			im = im - 1
		} else {
			in = in + 1
		}
	}
	return false

}

// @lc code=end

/*
// @lcpr case=start
// [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]\n5\n
// @lcpr case=end

// @lcpr case=start
// [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]\n20\n
// @lcpr case=end

*/

