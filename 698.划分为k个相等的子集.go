/*
 * @lc app=leetcode.cn id=698 lang=golang
 * @lcpr version=30100
 *
 * [698] 划分为k个相等的子集
 */

// @lc code=start
func canPartitionKSubsets(nums []int, k int) bool {
	// 异常值处理
	if k < 0 {
		return false
	}

	if len(nums) == 0 {
		return false
	}
	// 计算总和
	sum := 0
	for _, v := range nums {
		sum += v
	}

	// 计算每个桶的目标值
	target := sum / k

	// 剪枝 处理异常, 如果最小值比 target 大, 不可能有正确划分
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })

	if nums[0] > target {
		return false
	}

	// 不能整➗, 也有问题
	if sum%k != 0 {
		return false
	}

	// 长度不够分, 也有问题
	if len(nums) < k {
		return false
	}

	// 将 nums 均匀分布成 k 份, 每一份中数量为 n, 这种需要使用回溯
	bluks := make([]int, k)

	var dfs func(depth int) bool
	dfs = func(depth int) bool {
		// 能到最后一层, 改变之前的限制都已经满足了, 每个桶都已经满了
		if depth == len(nums) {
			return true
		}

		for i := 0; i < k; i++ {
			if bluks[i]+nums[depth] > target {
				continue
			}
			bluks[i] += nums[depth]
			if dfs(depth + 1) {
				return true
			}
			bluks[i] -= nums[depth]
			if bluks[i] == 0 {
				break
			}
		}
		return false
	}

	return dfs(0)
}

// @lc code=end

/*
// @lcpr case=start
// [10,1,10,9,6,1,9,5,9,10,7,8,5,2,10,8]\n11\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4]\n3\n
// @lcpr case=end

*/

