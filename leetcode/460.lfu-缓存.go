package leetcode

import (
	"container/list"
)

/*
 * @lc app=leetcode.cn id=460 lang=golang
 *
 * [460] LFU 缓存
 */

// @lc code=start
type LFUCache struct {
	MElement map[int]*list.Element //key- 值
	MTime    map[int]int           //key- 时间
	MCount   map[int]int           //key- 次数
	Link     map[int]*list.List    //次数- key链
	Time     int                   //时间序列
	Capacity int
	CountSeq []int // 次数序列
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		MElement: map[int]*list.Element{},
		MTime:    map[int]int{},
		MCount:   map[int]int{},
		Link:     map[int]*list.List{},
		Capacity: capacity,
		CountSeq: []int{},
	}
}

func (this *LFUCache) Get(key int) int {
	if element, ok := this.MElement[key]; !ok {
		return -1
	} else {
		val := element.Value.(int)
		this.Link[this.MCount[key]].Remove(element)
		this.MCount[key]++
		if this.MCount[key] > this.CountSeq[len(this.CountSeq)-1] {
			if this.Link[this.MCount[key]] != nil && this.Link[this.MCount[key]].Len() > 1 {
				this.CountSeq = append(this.CountSeq, this.MCount[key])
			} else {
				if this.Link[this.MCount[key]] == nil {
					this.Link[this.MCount[key]] = list.New()
				}
				this.CountSeq[len(this.CountSeq)-1] = this.MCount[key]
			}
		}
		this.Link[this.MCount[key]].PushFront(val)
		this.MElement[key] = this.Link[this.MCount[key]].Front()
		this.Time++
		this.MTime[key] = this.Time
		return val
	}
}

// type LFUCache struct {
// 	MElement map[int]*list.Element //key- 值
// 	MTime    map[int]int           //key- 时间
// 	MCount   map[int]int           //key- 次数
// 	Link     map[int]*list.List    //次数- key链
// 	Time     int                   //时间序列
// 	Capacity int
// 	CountSeq []int // 次数序列
// }
func (this *LFUCache) Put(key int, value int) {
	if element, ok := this.MElement[key]; !ok { //不存在
		if len(this.MElement) == this.Capacity {
			this.Capacity--
			count := this.CountSeq[0]
			this.CountSeq = this.CountSeq[1:]
			for (this.Link[count] == nil || this.Link[count].Len() == 0) && len(this.CountSeq) > 0 {
				count = this.CountSeq[0]
				this.CountSeq = this.CountSeq[1:]
			}
			elemet := this.Link[count].Back()
			val := elemet.Value.(int)
			delete(this.MElement, val)
			delete(this.MTime, val)
			delete(this.MCount, val)
			this.Link[count].Remove(elemet)
		}
		this.CountSeq = append([]int{1}, this.CountSeq...)
		this.Time++
		this.MTime[key] = this.Time
		this.MCount[key]++
		this.Capacity++
		if this.Link[this.MCount[key]] == nil {
			this.Link[this.MCount[key]] = list.New()
		}
		this.MElement[key] = this.Link[this.MCount[key]].PushFront(value)
	} else {
		this.Link[this.MCount[key]].Remove(element)
		this.MCount[key]++
		if this.MCount[key] > this.CountSeq[len(this.CountSeq)-1] {
			if this.Link[this.MCount[key]].Len() > 1 {
				this.CountSeq = append(this.CountSeq, this.MCount[key])
			} else {
				if this.Link[this.MCount[key]] == nil {
					this.Link[this.MCount[key]] = list.New()
				}
				this.CountSeq[len(this.CountSeq)-1] = this.MCount[key]
			}
		}
		this.Link[this.MCount[key]].PushFront(value)
		this.MElement[key] = this.Link[this.MCount[key]].Front()
		this.Time++
		this.MTime[key] = this.Time
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
