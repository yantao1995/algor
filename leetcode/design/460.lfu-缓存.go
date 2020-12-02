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
	CountSortList     *list.List            //次数 小到大
	KeyCountMap       map[int]int           //key - 次数
	KeyValueMap       map[int]int           //key - value
	KeySortElementMap map[int]*list.Element // sort key - element
	KeyElementMap     map[int]*list.Element //key - element
	CountLinkMap      map[int]*list.List    //次数 - 链
	Capacity          int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		CountSortList:     list.New(),
		KeyCountMap:       map[int]int{},
		KeyValueMap:       map[int]int{},
		KeySortElementMap: map[int]*list.Element{},
		KeyElementMap:     map[int]*list.Element{},
		CountLinkMap:      map[int]*list.List{},
		Capacity:          capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	if v, ok := this.KeyValueMap[key]; this.Capacity < 1 || !ok {
		return -1
	} else {
		this.Visit(key, v)
		return v
	}
}

func (this *LFUCache) Put(key int, value int) {
	if this.Capacity > 0 {
		if _, ok := this.KeyValueMap[key]; ok {
			this.Visit(key, value)
			return
		}
		if this.Capacity == len(this.KeyValueMap) {
			firstSortElement := this.CountSortList.Front()
			count := firstSortElement.Value.(int)
			link := this.CountLinkMap[count]
			linkElement := link.Front()
			keyDelete := linkElement.Value.(int)
			link.Remove(linkElement)
			if link.Len() == 0 {
				this.CountSortList.Remove(firstSortElement)
				delete(this.CountLinkMap, count)
			}
			delete(this.KeyCountMap, keyDelete)
			delete(this.KeyValueMap, keyDelete)
			delete(this.KeySortElementMap, keyDelete)
			delete(this.KeyElementMap, keyDelete)
		}
		//
		firstSortElement := this.CountSortList.Front()
		if firstSortElement == nil || firstSortElement.Value.(int) != 1 {
			this.CountSortList.PushFront(1)
		}
		firstSortElement = this.CountSortList.Front()
		this.KeyCountMap[key] = 1
		this.KeyValueMap[key] = value
		this.KeySortElementMap[key] = firstSortElement
		if this.CountLinkMap[1] == nil {
			this.CountLinkMap[1] = list.New()
		}
		this.KeyElementMap[key] = this.CountLinkMap[1].PushBack(key)
	}
}

func (this *LFUCache) Visit(key int, value int) {
	count := this.KeyCountMap[key]
	element := this.KeyElementMap[key]
	sortElement := this.KeySortElementMap[key]
	this.CountLinkMap[count].Remove(element)
	if nextSortElement := sortElement.Next(); nextSortElement != nil {
		if nextSortElement.Value.(int)-sortElement.Value.(int) > 1 {
			this.CountSortList.InsertAfter(sortElement.Value.(int)+1, sortElement)
		}
	} else {
		this.CountSortList.PushBack(sortElement.Value.(int) + 1)
	}
	nextSortElement := sortElement.Next()
	//
	this.KeyValueMap[key] = value
	this.KeyCountMap[key]++
	this.KeySortElementMap[key] = nextSortElement
	if this.CountLinkMap[count+1] == nil {
		this.CountLinkMap[count+1] = list.New()
	}
	element2 := this.CountLinkMap[count+1].PushBack(key)
	this.KeyElementMap[key] = element2
	if this.CountLinkMap[count] == nil || this.CountLinkMap[count].Len() == 0 {
		this.CountSortList.Remove(sortElement)
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
