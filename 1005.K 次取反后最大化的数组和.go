/*
 * @lc app=leetcode.cn id=1005 lang=golang
 * @lcpr version=30104
 *
 * [1005] K 次取反后最大化的数组和
 */

// @lc code=start
func largestSumAfterKNegations(nums []int, k int) int {
	// 1. 排序
	// 2. 遍历数组, 将负数取反, 直到k为0
	// 3. 如果k为奇数, 则将最小的数取反
	// 4. 计算数组和

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	fmt.Println(nums)

	sort.Ints(nums)
	minNums := math.MaxInt32
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k > 0 {
			nums[i] = -nums[i]
			k--
		}
		sum += nums[i]
		minNums = min(minNums, nums[i])
	}

	if k%2 == 1 {
		sum = sum - 2*minNums
	}
	return sum

}

// @lc code=end

/*
// @lcpr case=start
// [4,2,3]\n1\n
// @lcpr case=end

// @lcpr case=start
// [3,-1,0,2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [2,-3,-1,5,-4]\n2\n
// @lcpr case=end

*/

