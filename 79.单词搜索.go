/*
 * @lc app=leetcode.cn id=79 lang=golang
 * @lcpr version=30200
 *
 * [79] 单词搜索
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}

	n := len(board[0])
	if len(word) == 0 {
		return true
	}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var dfs func(i, j, index int) bool
	dfs = func(i, j, index int) bool {
		// 1. 成功找到完整单词！ (Base Case 1)
		if index == len(word) {
			return true
		}
		// 2. 剪枝：越界、已访问、字符不匹配 (Base Case 2 - 失败)
		if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || board[i][j] != word[index] {
			return false
		}

		// 3. 标记当前格子已访问
		visited[i][j] = true

		// 4. 探索四个方向的邻居 (注意：这里的边界检查 i-1>=0 等其实可以省略，
		//    因为下一层 dfs 调用开头会检查。但写上也没错，只是略显重复)
		if dfs(i-1, j, index+1) { // 向上
			// 注意：如果找到路径，这里会直接 return true，
			// 下面的 visited[i][j] = false 不会被执行，这是正确的！
			return true
		}
		if dfs(i+1, j, index+1) { // 向下
			return true
		}
		if dfs(i, j-1, index+1) { // 向左
			return true
		}
		if dfs(i, j+1, index+1) { // 向右
			return true
		}

		// 5. 回溯：如果四个方向都走不通，恢复当前格子的状态
		visited[i][j] = false

		// 6. 返回 false，表示从当前格子出发，无法构成目标单词
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == byte(word[0]) {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"\n
// @lcpr case=end

// @lcpr case=start
// [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"SEE"\n
// @lcpr case=end

// @lcpr case=start
// [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCB"\n
// @lcpr case=end

*/

