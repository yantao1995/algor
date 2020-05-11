package queue

import (
	"algor/vals"
	"fmt"
)

// [head  tail)
type Queue struct {
	QueueSlice []vals.AlgorType
	Head       int //队首
	Tail       int //队尾
	Size       int
}

//左入
func (q *Queue) LPush(str vals.AlgorType) *Queue {
	if q.Head == 0 {
		if q.Size == len(q.QueueSlice) {
			s := []vals.AlgorType{str}
			s = append(s, q.QueueSlice...)
			q.QueueSlice = s
		} else { //右移
			for i := 0; i < q.Size; i++ {
				q.QueueSlice[q.Tail-i+1] = q.QueueSlice[q.Tail-i]
			}
		}
		if q.Size != 0 {
			q.Tail++
		}
	} else {
		if q.Size != 0 {
			q.Head--
		}
	}
	q.QueueSlice[q.Head] = str
	q.Size++
	return q
}

//右入
func (q *Queue) RPush(str vals.AlgorType) *Queue {
	if q.Size < len(q.QueueSlice) {
		if q.Head != 0 {
			if q.Head != q.Tail || q.Size != 0 {
				for i := 0; i < q.Size; i++ { //左移
					q.QueueSlice[q.Head+i-1] = q.QueueSlice[q.Head+i]
				}
			}
			q.QueueSlice[q.Tail] = str
			if q.Size != 0 {
				q.Head--
			}
		} else {
			if q.Size != 0 {
				q.Tail++
			}
			q.QueueSlice[q.Tail] = str
		}
	} else {
		q.QueueSlice = append(q.QueueSlice, str)
		if q.Size != 0 {
			q.Tail++
		}
	}
	q.Size++
	return q
}

//左出
func (q *Queue) LPop() vals.AlgorType {
	var rtnVal vals.AlgorType
	if q.Size > 0 {
		//fmt.Printf("%v\t", q.QueueSlice[q.Head])
		rtnVal = q.QueueSlice[q.Head]
		q.QueueSlice[q.Head] = vals.AlgorNilVaule
		if q.Head != q.Tail {
			q.Head++
		}
		q.Size--
		return rtnVal
	}
	return vals.AlgorNilVaule
}

//左全出
func (q *Queue) LPopAll() {
	for q.Size > 0 {
		fmt.Printf("%v\t", q.QueueSlice[q.Head])
		q.QueueSlice[q.Head] = vals.AlgorNilVaule
		q.Size--
		if q.Head != q.Tail {
			q.Head++
		}
	}
	println()
}

//右出
func (q *Queue) RPop() vals.AlgorType {
	var rtnVal vals.AlgorType
	if q.Size > 0 {
		//fmt.Printf("%v\t", q.QueueSlice[q.Tail])
		rtnVal = q.QueueSlice[q.Tail]
		q.QueueSlice[q.Tail] = vals.AlgorNilVaule
		if q.Head != q.Tail {
			q.Tail--
		}
		q.Size--
		return rtnVal
	}
	return vals.AlgorNilVaule
}

//右全出
func (q *Queue) RPopAll() {
	for q.Size > 0 {
		fmt.Printf("%v\t", q.QueueSlice[q.Tail])
		q.QueueSlice[q.Tail] = vals.AlgorNilVaule
		q.Size--
		if q.Head != q.Tail {
			q.Tail--
		}
	}
	println()
}

//TestStack
func TestQueue() {
	q := new(Queue)
	q.LPush("1")
	fmt.Println("beforePopsize", q.Size)
	fmt.Println("data", q.LPop(), "size", q.Size)
	fmt.Println("data", q.RPop(), "size", q.Size)
	q.LPush("1").LPush("2").LPush("3")
	fmt.Println("beforePopsize", q.Size)
	fmt.Println("data", q.RPop(), "size", q.Size)
	fmt.Println("data", q.RPop(), "size", q.Size)
	fmt.Println("data", q.RPop(), "size", q.Size)
	fmt.Println("data", q.RPop(), "size", q.Size)
	// queue3 := new(Queue)
	// queue3.LPush("1").LPush("2").LPush("3").LPush("4")
	// println(queue3.RPop())
	// println(queue3.LPop())
	// fmt.Println(queue3)
	// queue3.LPush("5")
	// fmt.Println(queue3)
	// queue3.LPush("6")
	// fmt.Println(queue3)
	// queue3.LPush("7")
	// fmt.Println(queue3)
	// queue3.LPush("8")
	// fmt.Println(queue3)
	// println("----------")
	// queue := new(Queue)
	// queue.LPush("1").LPush("2").LPush("3")
	// fmt.Println(queue)
	// queue.LPopAll()
	// fmt.Println(queue)
	// queue.LPop()
	// println()
	// queue.LPush("4").LPush("5")
	// fmt.Println(queue)
	// queue.LPop()
	// println()
	// fmt.Println(queue)
	// queue.LPop()
	// println()
	// fmt.Println(queue)
	// println("----------")
	// queue.RPush("4")
	// fmt.Println(queue)
	// queue.RPush("5")
	// fmt.Println(queue)
	// queue.RPush("7")
	// fmt.Println(queue)
	// queue.RPush("8")
	// fmt.Println(queue)
	// println("----------")
	// queue2 := new(Queue)
	// queue2.RPush("1").RPush("2").RPush("3")
	// fmt.Println(queue2)
	// queue2.RPop()
	// println()
	// fmt.Println(queue2)
	// queue2.RPop()
	// println()
	// fmt.Println(queue2)
	// queue2.RPop()
	// println()
	// fmt.Println(queue2)
	// println("----------")
	// queue2.RPush("3")
	// fmt.Println(queue2)
	// queue2.RPush("4")
	// fmt.Println(queue2)
	// queue2.RPopAll()
	// fmt.Println(queue2)
	// println("----------")
	// queue4 := new(Queue)
	// queue4.LPush("1").RPush("2").LPush("3")
	// fmt.Println(queue4)
	// queue4.RPop()
	// println()
	// queue4.LPop()
	// println()
	// fmt.Println(queue4)
	// queue4.LPush("1").RPush("2").LPush("3")
	// fmt.Println(queue4)
	// queue4.LPopAll()
	// fmt.Println(queue4)
}
