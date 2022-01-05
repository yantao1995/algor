package leetcode

import "container/list"

/*
 * @lc app=leetcode.cn id=677 lang=golang
 *
 * [677] 键值映射
 */

// @lc code=start
type MapSum struct {
	isWords bool
	val     int
	NextMap map[byte]*MapSum //前缀树
}

func Constructor() MapSum {
	return MapSum{
		isWords: false,
		val:     0,
		NextMap: map[byte]*MapSum{},
	}
}

func (this *MapSum) Insert(key string, val int) {
	that := this
	for i := 0; i < len(key); i++ {
		if _, ok := that.NextMap[key[i]]; !ok {
			that.NextMap[key[i]] = &MapSum{
				isWords: false,
				val:     0,
				NextMap: map[byte]*MapSum{},
			}
		}
		if i == len(key)-1 {
			that.NextMap[key[i]].isWords = true
			that.NextMap[key[i]].val = val
		}
		that = that.NextMap[key[i]]
	}
}

func (this *MapSum) Sum(prefix string) int {
	total := 0
	that := this
	for k := range prefix {
		if _, ok := that.NextMap[prefix[k]]; !ok {
			return 0
		}
		that = that.NextMap[prefix[k]]
	}
	queue := list.New()
	queue.PushBack(that)
	for queue.Len() > 0 {
		element := queue.Front()
		if ms, ok := element.Value.(*MapSum); ok && ms != nil {
			if ms.isWords {
				total += ms.val
			}
			for k := range ms.NextMap {
				queue.PushBack(ms.NextMap[k])
			}
		}
		queue.Remove(element)
	}
	return total
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
// @lc code=end
