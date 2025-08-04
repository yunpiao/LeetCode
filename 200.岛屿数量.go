/*
 * @lc app=leetcode.cn id=200 lang=golang
 * @lcpr version=30005
 *
 * [200] 岛屿数量
 */

// @lcpr-template-start
package main

// @lcpr-template-end
// @lc code=start
func numIslands(grid [][]byte) int {
	// 深度优先遍历的思想, 如果发现 1, 就标记与其想连接的节点
	rows := len(grid)
	cols := len(grid[0])
	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				count++
				// 深度优先遍历, 将与其相连的 1 都标记为 0
				bfs(grid, i, j)
			}
		}
	}
	return count
}

func bfs(grid [][]byte, i, j int) {
	row := len(grid)
	col := len(grid[0])
	grid[i][j] = '0'
	queue := [][2]int{[2]int{i, j}}
	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]

		if top[0]-1 >= 0 && grid[top[0]-1][top[1]] == '1' {
			grid[top[0]-1][top[1]] = '0'
			queue = append(queue, [2]int{top[0] - 1, top[1]})
		}
		if top[0]+1 < row && grid[top[0]+1][top[1]] == '1' {
			grid[top[0]+1][top[1]] = '0'
			queue = append(queue, [2]int{top[0] + 1, top[1]})
		}
		if top[1]-1 >= 0 && grid[top[0]][top[1]-1] == '1' {
			grid[top[0]][top[1]-1] = '0'
			queue = append(queue, [2]int{top[0], top[1] - 1})
		}
		if top[1]+1 < col && grid[top[0]][top[1]+1] == '1' {
			grid[top[0]][top[1]+1] = '0'
			queue = append(queue, [2]int{top[0], top[1] + 1})
		}
	}

}

func dfs(grid [][]byte, i, j int) {
	rows := len(grid)
	cols := len(grid[0])
	grid[i][j] = '0'
	// 向上
	if i-1 >= 0 && grid[i-1][j] == '1' {
		dfs(grid, i-1, j)
	}
	// 向下
	if i+1 < rows && grid[i+1][j] == '1' {
		dfs(grid, i+1, j)
	}
	// 向左
	if j-1 >= 0 && grid[i][j-1] == '1' {
		dfs(grid, i, j-1)
	}
	// 向右
	if j+1 < cols && grid[i][j+1] == '1' {
		dfs(grid, i, j+1)
	}
}

// @lc code=end

/*
// @lcpr case=start
// [["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]\n
// @lcpr case=end

// @lcpr case=start
// [["1","1","0","0","0"],["1","1","0","0","0"],["0","0","1","0","0"],["0","0","0","1","1"]]\n
// @lcpr case=end

*/
