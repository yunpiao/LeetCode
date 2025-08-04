/*
 * @lc app=leetcode.cn id=134 lang=golang
 * @lcpr version=30104
 *
 * [134] 加油站
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	// 贪心解法
	// 1. 如果总油量小于总消耗, 则无法到达
	// 2. 如果从i出发, 油量小于消耗, 则i+1出发
	// 因为如果有答案, 是唯一, 所以可以直接推出 i 是答案
	gasSum := 0
	costSum := 0
	for i := range gas {
		gasSum += gas[i]
		costSum += cost[i]
	}
	if costSum > gasSum {
		return -1
	}
	currGas := 0
	ret := 0
	for i := range gas {
		if currGas < 0 {
			currGas = 0
			ret = i
		}
		currGas = currGas + (gas[i] - cost[i])
	}
	return ret

	// 暴力解法
	for i := range gas {
		totalGas := 0
		for j := i; j < len(gas)+i; j++ {
			totalGas += gas[j%len(gas)]
			totalGas -= cost[j%len(gas)]
			if totalGas < 0 {
				break
			}
		}
		if totalGas >= 0 {
			return i
		}
	}
	return -1

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n[3,4,5,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [2,3,4]\n[3,4,3]\n
// @lcpr case=end

*/

