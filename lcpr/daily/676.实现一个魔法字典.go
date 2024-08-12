package lcpr

/*
 * @lc app=leetcode.cn id=676 lang=golang
 * @lcpr version=30204
 *
 * [676] 实现一个魔法字典
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
type MagicDictionary struct {
	lenSlice [][]string
}

func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	maxLength := 0
	for _, v := range dictionary {
		maxLength = max(maxLength, len(v))
	}
	this.lenSlice = make([][]string, maxLength+1)
	for _, v := range dictionary {
		this.lenSlice[len(v)] = append(this.lenSlice[len(v)], v)
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	if len(searchWord) >= len(this.lenSlice) {
		return false
	}
	for _, word := range this.lenSlice[len(searchWord)] {
		disCount := 0
		for i := 0; i < len(searchWord); i++ {
			if searchWord[i] != word[i] {
				disCount++
			}
			if disCount > 1 {
				break
			}
		}
		if disCount == 1 {
			return true
		}
	}
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
// @lc code=end
