/*
 * @lc app=leetcode.cn id=455 lang=golang
 * @lcpr version=30104
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	// 贪心
	// g 中有一个序列, s 中一个序列, 需要田忌赛马
	sort.Ints(g)
	sort.Ints(s)
	si := len(s) - 1
	ret := 0
	for i := len(g) - 1; i >= 0; i-- {
		// s[si] 中剩余最大的, 喂给 g[i] 中最大的
		if g[i] <= s[si] {
			si--
			ret++
		}
	}
	return ret
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n[1,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n[1,2,3]\n
// @lcpr case=end

*/

