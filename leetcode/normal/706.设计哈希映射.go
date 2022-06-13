package leetcode

/*
 * @lc app=leetcode.cn id=706 lang=golang
 *
 * [706] 设计哈希映射
 */

// @lc code=start
// type MyHashMap struct {
// 	Slice []int
// }

// /** Initialize your data structure here. */
// func Constructor() MyHashMap {
// 	hm := MyHashMap{
// 		Slice: make([]int, 1000000),
// 	}
// 	for k := range hm.Slice {
// 		hm.Slice[k] = -1
// 	}
// 	return hm
// }

// /** value will always be non-negative. */
// func (this *MyHashMap) Put(key int, value int) {
// 	this.Slice[key] = value
// }

// /** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
// func (this *MyHashMap) Get(key int) int {
// 	return this.Slice[key]
// }

// /** Removes the mapping of the specified value key if this map contains a mapping for the key */
// func (this *MyHashMap) Remove(key int) {
// 	this.Slice[key] = -1
// }

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
// @lc code=end
