/*
 * @lc app=leetcode.cn id=207 lang=golang
 * @lcpr version=30104
 *
 * [207] 课程表
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	m := make(map[int][]int, numCourses)
	for _, per := range prerequisites {
		m[per[0]] = append(m[per[0]], per[1])
	}

	visited := make(map[int]int, numCourses)

	var dfs func(i int) bool
	dfs = func(i int) bool {
		for _, v := range m[i] {

			if visited[v] == 1 {
				return false
			}
			if visited[v] == 2 {
				continue
			}

			visited[v] = 1

			if !dfs(v) {
				return false
			}
			visited[v] = 2
		}
		return true
	}

	for _, per := range m {
		for _, v := range per {
			if !dfs(v) {
				return false
			}
		}
	}
	return true

}

// @lc code=end

/*
// @lcpr case=start
// 2\n[[1,0]]\n
// @lcpr case=end

// @lcpr case=start
// 2\n[[1,0],[0,1]]\n
// @lcpr case=end

*/
