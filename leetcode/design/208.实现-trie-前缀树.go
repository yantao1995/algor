package leetcode

/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
// type Trie struct {
// 	isWords  bool
// 	NextTrie map[byte]*Trie //前缀树
// }

// func Constructor() Trie {
// 	return Trie{
// 		isWords:  false,
// 		NextTrie: map[byte]*Trie{},
// 	}
// }

// func (this *Trie) Insert(word string) {
// 	for i := 0; i < len(word); i++ {
// 		if _, ok := this.NextTrie[word[i]]; !ok {
// 			this.NextTrie[word[i]] = &Trie{
// 				isWords:  false,
// 				NextTrie: map[byte]*Trie{},
// 			}
// 		}
// 		if i == len(word)-1 {
// 			this.NextTrie[word[i]].isWords = true
// 		}
// 		this = this.NextTrie[word[i]]
// 	}
// }

// func (this *Trie) Search(word string) bool {
// 	for i := 0; i < len(word); i++ {
// 		if _, ok := this.NextTrie[word[i]]; !ok {
// 			return false
// 		}
// 		if i == len(word)-1 {
// 			return this.NextTrie[word[i]].isWords
// 		}
// 		this = this.NextTrie[word[i]]
// 	}
// 	return false
// }

// func (this *Trie) StartsWith(prefix string) bool {
// 	for i := 0; i < len(prefix); i++ {
// 		if _, ok := this.NextTrie[prefix[i]]; !ok {
// 			return false
// 		}
// 		this = this.NextTrie[prefix[i]]
// 	}
// 	return true
// }

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
