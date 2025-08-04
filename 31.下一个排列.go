/*
 * @lc app=leetcode.cn id=31 lang=golang
 * @lcpr version=30005
 *
 * [31] 下一个排列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func nextPermutation(nums []int) {
	// 1. 先找到拐点, 从后往前找, 一个下降的拐点位置
	// 2. 如果有拐点, 继续从后往前, 找到第一个大于拐点位置的点
	// 3. 交换拐点和找到的点
	// 4. 拐点后的数据逆序下
	n := len(nums)
	if n < 1 {
		return
	}
	p1, p2 := -1, -1
	for i := n - 1; i > 0; i-- {
		// 找到比后一个小的值
		if nums[i] > nums[i-1] {
			p1 = i - 1
			break
		}
	}
	if p1 != -1 {
		for i := n - 1; i > p1; i-- {
			if nums[p1] < nums[i] {
				p2 = i
				break
			}
		}
		nums[p1], nums[p2] = nums[p2], nums[p1]
	}
	reverse(nums[p1+1:])
}

func reverse(nums []int) {
	n := len(nums)
	// 逆序数组
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [3,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,5]\n
// @lcpr case=end

*/

