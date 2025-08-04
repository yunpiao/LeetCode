/*
 * @lc app=leetcode.cn id=406 lang=golang
 * @lcpr version=30104
 *
 * [406] 根据身高重建队列
 */

// @lc code=start
func reconstructQueue(people [][]int) [][]int {
	// 关于出现两个维度一起考虑的情况，我们已经做过两道题目了，另一道就是135. 分发糖果 (opens new window)。
	// 其技巧都是确定一边然后贪心另一边，两边一起考虑，就会顾此失彼。
	// 这道题目可以说比135. 分发糖果 (opens new window)难不少，其贪心的策略也是比较巧妙。

	// 先确定k，再确定h
	// 先按照h从大到小排序，如果h相同，则按照k从小到大排序
	// 然后遍历排序后的数组，根据k值插入到结果数组中
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	// fmt.Println(people)
	ret := make([][]int, 0, len(people))
	// for i := range people {
	// 	inserIndex := people[i][1]
	// 	ret = append(ret[:inserIndex], append([][]int{people[i]}, ret[inserIndex:]...)...)
	// }
	for _, p := range people {
		index := p[1]
		ret = append(ret[:index], append([][]int{p}, ret[index:]...)...)
	}
	return ret
}

// @lc code=end

/*
// @lcpr case=start
// [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]\n
// @lcpr case=end

// @lcpr case=start
// [[6,0],[5,0],[4,0],[3,2],[2,2],[1,4]]\n
// @lcpr case=end

*/

