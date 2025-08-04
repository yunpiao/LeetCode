/*
 * @lc app=leetcode.cn id=208 lang=golang
 * @lcpr version=30104
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
// 前缀树有两个关键点
// 1. 到下一层怎么索引, 使用 26 数组
// 2. 怎么确定一个 str 是否在树中, 用一个 isEnd, 如果遍历结束,之后有  end 就是代表存在
type Trie struct {
	Children [26]*Trie
	IsEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	curr := this
	for _, v := range word {
		if curr.Children[v-'a'] == nil {
			curr.Children[v-'a'] = &Constructor()
		}
		curr = curr.Children[v-'a']
	}
	curr.IsEnd = true
}

func (this *Trie) Search(word string) bool {
	curr := this
	for _, v := range word {
		if curr.Children[v-'a'] == nil {
			return false
		}
		curr = curr.Children[v-'a']
	}
	return curr.IsEnd

}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for _, v := range prefix {
		if curr.Children[v-'a'] == nil {
			return false
		}
		curr = curr.Children[v-'a']
	}
	return true

}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end



