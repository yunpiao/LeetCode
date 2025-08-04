/*
 * @lc app=leetcode.cn id=860 lang=golang
 * @lcpr version=30104
 *
 * [860] 柠檬水找零
 */

// @lc code=start
func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	// 贪心算法, 优先使用10元, 如果10元不够, 则使用5元
	// 因为5元可以找零给10元, 而10元不能找零给20元
	// 所以5元需要尽可能多

	// 遍历bills, 如果当前是5元, 则five++
	// 如果当前是10元, 则ten++, five--
	// 如果当前是20元, 则ten--, five-- 或者 five -= 3
	// 如果five < 0, 则返回false
	for i := range bills {
		if bills[i] == 5 {
			five++
		}
		if bills[i] == 10 {
			ten++
			five--
		}
		if bills[i] == 20 {
			if ten > 0 {
				ten--
				five--
			} else {
				five -= 3
			}
		}
		if five < 0 {
			return false
		}
	}
	return true

}

// @lc code=end

/*
// @lcpr case=start
// [5,5,5,10,20]\n
// @lcpr case=end

// @lcpr case=start
// [5,5,10,10,20]\n
// @lcpr case=end

*/

