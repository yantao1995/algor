package leetcode

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 */

// @lc code=start
// type LRUCache struct {
// 	Lst      *list.List //队列
// 	Ht       map[int]int
// 	Capacity int
// }

// func Constructor(capacity int) LRUCache {
// 	return LRUCache{
// 		Lst:      list.New(),
// 		Ht:       map[int]int{},
// 		Capacity: capacity,
// 	}
// }

// func (this *LRUCache) Get(key int) int {
// 	if _, ok := this.Ht[key]; !ok {
// 		return -1
// 	}
// 	for node := this.Lst.Front(); node != nil; node = node.Next() {
// 		if node.Value.(int) == key {
// 			this.Lst.MoveToFront(node)
// 		}
// 	}
// 	return this.Ht[key]
// }

// func (this *LRUCache) Put(key int, value int) {
// 	if _, ok := this.Ht[key]; ok { //存在 更新
// 		for node := this.Lst.Front(); node != nil; node = node.Next() {
// 			if node.Value.(int) == key {
// 				this.Lst.MoveToFront(node)
// 			}
// 		}
// 	} else {
// 		if this.Capacity == this.Lst.Len() {
// 			val := this.Lst.Back()
// 			delete(this.Ht, val.Value.(int))
// 			this.Lst.Remove(val)
// 		}
// 		this.Lst.PushFront(key)
// 	}
// 	this.Ht[key] = value
// }

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
