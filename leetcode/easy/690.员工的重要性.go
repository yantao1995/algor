package easy

import "container/list"

/*
 * @lc app=leetcode.cn id=690 lang=golang
 *
 * [690] 员工的重要性
 */

// @lc code=start
/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * } */ //

func getImportance(employees []*Employee, id int) int {
	ht := map[int]*Employee{}
	for k := range employees {
		ht[employees[k].Id] = employees[k]
	}
	result := ht[id].Importance
	lst := list.New()
	for _, v := range ht[id].Subordinates {
		lst.PushFront(v)
	}
	for lst.Len() > 0 {
		id := lst.Back()
		lst.Remove(id)
		result += ht[id.Value.(int)].Importance
		for _, v := range ht[id.Value.(int)].Subordinates {
			lst.PushFront(v)
		}
	}
	return result
}

// @lc code=end

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}
