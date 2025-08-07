/*
 * @lc app=leetcode.cn id=349 lang=golang
 * @lcpr version=30104
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	set := make(map[int]struct{}, len(nums2))
	ret := []int{}
	for _, v := range nums2 {
		set[v] = struct{}{}
	}	

	for _, v := range nums1 {
		if _, exist := set[v]; exist {
			delete(set, v)
			ret = append(ret, v)
		}
	}

	return ret
}

// @lc code=end

/*
// @lcpr case=start
// [1000,2,2,1000]\n[1000,1000]\n
// @lcpr case=end

// @lcpr case=start
// [4,9,5]\n[9,4,9,8,4]\n
// @lcpr case=end

*/

