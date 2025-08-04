/*
 * @lc app=leetcode.cn id=763 lang=golang
 * @lcpr version=30104
 *
 * [763] 划分字母区间
 */
package main

// @lc code=start
func partitionLabels(s string) []int {
	// 统计当前的所有字符的最后出现 idnex 都小于 i
	dict := [26]int{}
	for i, v := range s {
		dict[v-'a'] = i
	}

	ret := []int{}
	left, right := -1, -1
	for i, v := range s {
		right = max(right, dict[v-'a'])
		if i == right {
			ret = append(ret, right-left)
			left = right
		}
	}
	return ret

	// // 用重叠区间来计算
	// ret := []int{}
	// if len(s) == 0 {
	// 	return ret
	// }

	// // 得到字典
	// m := make(map[rune][]int)
	// for i, v := range s {
	// 	if m[v] == nil {
	// 		m[v] = []int{len(s), -1}
	// 	}
	// 	m[v][0] = min(m[v][0], i)
	// 	m[v][1] = max(m[v][1], i)
	// }

	// currMax := 0
	// count := 0
	// for _, v := range s {
	// 	if m[v][0] > currMax {
	// 		ret = append(ret, count)
	// 		count = 0
	// 	}
	// 	currMax = max(currMax, m[v][1])
	// 	count++
	// }
	// ret = append(ret, count)

	// return ret

}

// @lc code=end

/*
// @lcpr case=start
// "abcad"\n
// @lcpr case=end

// @lcpr case=start
// "eccbbbbdec"\n
// @lcpr case=end

*/
