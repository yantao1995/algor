package leetcode

import (
	"math/rand"
	"time"
)

/*
 * @lc app=leetcode.cn id=380 lang=golang
 *
 * [380] 常数时间插入、删除和获取随机元素
 */

// @lc code=start
type RandomizedSet struct {
	HashMap map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		HashMap: map[int]int{},
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.HashMap[val]; ok {
		return false
	}
	this.HashMap[val] = val
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.HashMap[val]; !ok {
		return false
	}
	delete(this.HashMap, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	rd := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(this.HashMap))
	index := 0
	for k := range this.HashMap {
		if index == rd {
			return this.HashMap[k]
		}
		index++
	}
	return 0
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
