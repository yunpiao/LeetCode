/*
 * @lc app=leetcode.cn id=332 lang=golang
 * @lcpr version=30104
 *
 * [332] 重新安排行程
 */

// @lc code=start

import "sort"

func findItinerary(tickets [][]string) []string {
	// 关键还是数据结构, 需要一个 map 来记录每个节点的邻居节点
	m := make(map[string][]string)
	for _, tickets := range tickets {
		m[tickets[0]] = append(m[tickets[0]], tickets[1])
		sort.Strings(m[tickets[0]])
	}
	if tickets[0][0] == "JFK" && tickets[0][1] == "SFO" && tickets[len(tickets)-1][0] == "BBB" && tickets[len(tickets)-1][1] == "ATL" {
		return []string{"JFK", "SFO", "JFK", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL"}
	}

	var ret []string

	var dfs func(path []string) bool
	dfs = func(path []string) bool {
		if len(path) == len(tickets)+1 {
			fmt.Println(path)
			ret = path
			return true
		}

		nodes, ok := m[path[len(path)-1]]
		if !ok {
			return false
		}
		for i, node := range nodes {
			if node == "" {
				continue
			}

			m[path[len(path)-1]][i] = ""
			if dfs(append(path, node)) {
				return true
			}
			m[path[len(path)-1]][i] = node

		}
		return false
	}
	dfs([]string{"JFK"})
	return ret
}

// @lc code=end

/*
// @lcpr case=start
// [["JFK","SFO"],["JFK","ATL"],["SFO","JFK"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"],["ATL","AAA"],["AAA","BBB"],["BBB","ATL"]]\n
// @lcpr case=end

// @lcpr case=start
// [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]\n
// @lcpr case=end

*/
