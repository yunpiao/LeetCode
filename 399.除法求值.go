/*
 * @lc app=leetcode.cn id=399 lang=golang
 * @lcpr version=30104
 *
 * [399] 除法求值
 */
package main

// @lc code=start
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 深度优先搜索, 还是图部分
	// 建立图

	type to struct {
		next  string
		value float64
	}
	g := make(map[string][]to, len(equations))
	for i := range equations {
		g[equations[i][0]] = append(g[equations[i][0]], to{
			next:  equations[i][1],
			value: values[i],
		})
		g[equations[i][1]] = append(g[equations[i][1]], to{
			next:  equations[i][0],
			value: 1 / values[i],
		})
	}

	visited := make(map[string]bool, len(g))

	var dfs func(string, string) float64
	dfs = func(start, end string) float64 {
		// 如果找到了, 直接返回 1, 这里题目要求, 图中不出现的返回 -1, 所以需要提前判断
		if _, ok := g[start]; !ok {
			return -1
		}
		// 如果找到了, 直接返回 1
		if start == end {
			return 1
		}

		visited[start] = true
		for _, to := range g[start] {
			if visited[to.next] {
				continue
			}
			ans := dfs(to.next, end)
			if ans != -1 {
				return ans * to.value
			}
		}
		return -1
	}

	ret := make([]float64, 0, len(queries))
	for i := range queries {
		visited = make(map[string]bool, len(g))
		ret = append(ret, dfs(queries[i][0], queries[i][1]))
	}

	return ret
}

// @lc code=end

/*
// @lcpr case=start
// [["a","b"],["b","c"]]\n[2.0,3.0]\n[["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]\n
// @lcpr case=end

// @lcpr case=start
// [["a","b"],["b","c"],["bc","cd"]]\n[1.5,2.5,5.0]\n\n[["a","c"],["c","b"],["bc","cd"],["cd","bc"]]\n
// @lcpr case=end

// @lcpr case=start
// [["a","b"]]\n[0.5]\n[["a","b"],["b","a"],["a","c"],["x","y"]]\n
// @lcpr case=end

*/
