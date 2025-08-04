/*
 * @lc app=leetcode.cn id=621 lang=golang
 * @lcpr version=30200
 *
 * [621] 任务调度器
 */

// @lc code=start
func leastInterval(tasks []byte, n int) int {
	m := [26]int{}
	for _, v := range tasks {
		m[v-'A']++
	}
	// 找到最大值
	maxFreq := 0
	for i := range 26 {
		maxFreq = max(m[i], maxFreq)
	}
	maxCount := 0
	for i := range 26 {
		if m[i] == maxFreq {
			maxCount++
		}
	}

	return max((n+1)*(maxFreq-1)+maxCount, len(tasks))

}

// @lc code=end

/*
// @lcpr case=start
// ["A","A","A","B","B","B"]\n2\n
// @lcpr case=end

// @lcpr case=start
// ["A","C","A","B","D","B"]\n1\n
// @lcpr case=end

// @lcpr case=start
// ["A","A","A","B","B","B"]\n3\n
// @lcpr case=end

*/

