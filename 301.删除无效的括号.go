/*
 * @lc app=leetcode.cn id=301 lang=golang
 * @lcpr version=30104
 *
 * [301] 删除无效的括号
 */
package main

// @lc code=start
func removeInvalidParentheses(s string) []string {
	// 判断符号匹配的方法, 统计字符串, (+1)-1, 如果出现过<0, 返回错误,
	isValid := func(s string) bool {
		count := 0
		for _, ch := range s {
			if ch == '(' {
				count++
			} else if ch == ')' {
				count--
			}
			if count < 0 {
				return false
			}
		}
		return count == 0
	}

	visited := map[string]bool{s: true}
	// ✂️ 如果出现过, 不放入队列中了

	queue := []string{}
	queue = append(queue, s)
	var result []string

	// 经典 BFS, 一次取出同一层级的数据
	for len(queue) > 0 {
		found := false
		n := len(queue)

		for i := 0; i < n; i++ {
			top := queue[0]
			queue = queue[1:]

			// 合规, 加入到结果集里面
			if isValid(top) {
				result = append(result, top)
				found = true
				continue
			}

			// 每次去除一次符号, 入队等待检查
			for j := 0; j < len(top); j++ {
				if top[j] != '(' && top[j] != ')' {
					continue
				}
				next := top[:j] + top[j+1:]

				if !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}

			}
		}
		// 这一层发现过, 不进入下一层
		if found {
			break
		}
	}

	return result

}

// @lc code=end

/*
// @lcpr case=start
// "()())()"\n
// @lcpr case=end

// @lcpr case=start
// "(a)())()"\n
// @lcpr case=end

// @lcpr case=start
// ")("\n
// @lcpr case=end

*/
