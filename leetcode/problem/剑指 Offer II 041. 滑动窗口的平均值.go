package leetcode

/*

给定一个整数数据流和一个窗口大小，根据该滑动窗口的大小，计算滑动窗口里所有数字的平均值。
实现 MovingAverage 类：
MovingAverage(int size) 用窗口大小 size 初始化对象。
double next(int val) 成员函数 next 每次调用的时候都会往滑动窗口增加一个整数，请计算并返回数据流中最后 size 个值的移动平均值，即滑动窗口里所有数字的平均值。

*/
type MovingAverage struct {
	totalValue float64
	queue      []int
	ptr        int
}

/** Initialize your data structure here. */
// func Constructor(size int) MovingAverage {
// 	return MovingAverage{
// 		totalValue: 0,
// 		queue:      make([]int, 0, size),
// 		ptr:        0,
// 	}
// }

func (this *MovingAverage) Next(val int) float64 {
	this.totalValue += float64(val)
	if len(this.queue) < cap(this.queue) {
		this.queue = append(this.queue, val)
		return this.totalValue / float64(len(this.queue))
	}
	this.totalValue -= float64(this.queue[this.ptr])
	this.queue[this.ptr] = val
	this.ptr++
	if this.ptr == len(this.queue) {
		this.ptr = 0
	}
	return this.totalValue / float64(cap(this.queue))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */

/*
	队列queue存放历史数据
	ptr指针保存当前的初始位置
	cap(queue)保存size
	len(queue)为当前的长度

	长度不够就直接保存在末尾，然后求均值
	长度够了，那就减去当前 ptr位置的，然后加上新添加的数，然后移动指针，求均值，返回。
*/
