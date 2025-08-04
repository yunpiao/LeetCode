/*
 * @lc app=leetcode.cn id=213 lang=golang
 * @lcpr version=30005
 *
 * [213] 打家劫舍 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func rob(nums []int) int {
	// 这个在新的数据上增加了一个收尾的限制, 第一个和最后一个不能同时被偷
	// 但是不能通过标记, 因为第一个的状态会影响后续的状态
	// 所以需要通过计算 rob(nums[1:n-1]) 绝对没有第一个, + rob(nums[0:n-2]) 绝对没有最后一个, 两个加在一起, 将环形依赖,变为两个线性依赖
	// 注意特殊情况 最后一个和第一个可能是同一个, 需要特殊处理
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return max(rob1(nums[0:n-1]), rob1(nums[1:n]))
}

func rob1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp0, dp1 := 0, nums[0]
	for _, v := range nums[1:] {
		dp1, dp0 = max(dp0+v, dp1), dp1
	}
	return max(dp0, dp1)
}

// @lc code=end

/*
// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

*/

