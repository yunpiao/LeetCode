/*
 * @lc app=leetcode.cn id=349 lang=golang
 * @lcpr version=30104
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	// 减少内存分配, nums1 数据量较小, 统计较小的数据放入 map 中
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	// 统计数据
	numsSet := make(map[int]struct{})
	for _, nums := range nums1 {
		numsSet[nums] = struct{}{}
	}

	// 根据 set 判断有没有出现过, 出现过就放在结果表中
	result := make([]int, 0, len(nums1))
	for _, nums := range nums2 {
		if _, exist := numsSet[nums]; exist {
			result = append(result, nums)
			delete(numsSet, nums)
		}
	}
	return result
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,2,1]\n[2,2]\n
// @lcpr case=end

// @lcpr case=start
// [4,9,5]\n[9,4,9,8,4]\n
// @lcpr case=end

*/

