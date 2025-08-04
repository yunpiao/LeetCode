/*
 * @lc app=leetcode.cn id=59 lang=golang
 * @lcpr version=30104
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	nums := make([][]int, n)
	for i := range nums {
		nums[i] = make([]int, n)
	}

	index := 1
	for i := 0; i < n/2; i++ {
		// 左上 固定行为 i, 列从 i 到 n-1-i-1( 小于不等于 n-i-1)
		for j := i; j < n-1-i; j++ {
			nums[i][j] = index
			index++
		}
		// 右上 固定列为 n-1-i, 行从 i 到 n-1-i-1 (小于不等于 n-i-1)
		for j := i; j < n-1-i; j++ {
			nums[j][n-1-i] = index
			index++
		}
		// 右下 固定行为 n-1-i, 列从 n-1-i 到 i+1 (大于不等于 i)
		for j := n - 1 - i; j > i; j-- {
			nums[n-1-i][j] = index
			index++
		}
		// 左下 固定列为 i, 列从 n-i-1 到 i+1 (大于不等于 i)
		for j := n - 1 - i; j > i; j-- {
			nums[j][i] = index
			index++
		}
	}
	if n%2 == 1 {
		nums[(n-1)/2][(n-1)/2] = index
	}
	return nums
}

// @lc code=end

/*
// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/

