package queue

import "fmt"

const popNilVaule = "nil"

type queueType string

// [head  tail)
type Queue struct {
	QueueSlice []queueType
	Head       int //队首
	Tail       int //队尾
	Size       int
}

//左入
func (q *Queue) LPush(str queueType) *Queue {
	if q.Head == 0 {
		if q.Size == len(q.QueueSlice) {
			s := []queueType{str}
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
func (q *Queue) RPush(str queueType) *Queue {
	if q.Size < len(q.QueueSlice) {
		if q.Head != 0 {
			if q.Head != q.Tail || q.Size != 0 {
				fmt.Println(q.Size)
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
func (q *Queue) LPop() *Queue {
	if q.Size > 0 {
		fmt.Printf("%v\t", q.QueueSlice[q.Head])
		q.QueueSlice[q.Head] = popNilVaule
		if q.Head != q.Tail {
			q.Head++
		}
		q.Size--
		println()
	} else {
		println(popNilVaule)
	}
	return q
}

//左全出
func (q *Queue) LPopAll() {
	for q.Size > 0 {
		fmt.Printf("%v\t", q.QueueSlice[q.Head])
		q.QueueSlice[q.Head] = popNilVaule
		q.Size--
		if q.Head != q.Tail {
			q.Head++
		}
	}
	println()
}

//右出
func (q *Queue) RPop() *Queue {
	if q.Size > 0 {
		fmt.Printf("%v\t", q.QueueSlice[q.Tail])
		q.QueueSlice[q.Tail] = popNilVaule
		if q.Head != q.Tail {
			q.Tail--
		}
		q.Size--
		println()
	} else {
		println(popNilVaule)
	}
	return q
}

//右全出
func (q *Queue) RPopAll() {
	for q.Size > 0 {
		fmt.Printf("%v\t", q.QueueSlice[q.Tail])
		q.QueueSlice[q.Tail] = popNilVaule
		q.Size--
		if q.Head != q.Tail {
			q.Tail--
		}
	}
	println()
}

//TestStack
func TestQueue() {
	// queue3 := new(Queue)
	// queue3.LPush("1").LPush("2").LPush("3").LPush("4")
	// queue3.RPop()
	// queue3.LPop()
	// fmt.Println(queue3)
	// queue3.LPush("5")
	// fmt.Println(queue3)
	// queue3.LPush("6")
	// fmt.Println(queue3)
	// queue3.LPush("7")
	// fmt.Println(queue3)
	// queue3.LPush("8")
	// fmt.Println(queue3)
	// queue := new(Queue)
	// queue.LPush("1").LPush("2").LPush("3")
	// fmt.Println(queue)
	// queue.LPopAll()
	// fmt.Println(queue)
	// queue.LPop()
	// queue.LPush("4").LPush("5")
	// fmt.Println(queue)
	// queue.LPop()
	// fmt.Println(queue)
	// queue.LPop()
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
	// fmt.Println(queue2)
	// queue2.RPop()
	// fmt.Println(queue2)
	// queue2.RPop()
	// fmt.Println(queue2)
	// println("----------")
	// queue2.RPush("3")
	// fmt.Println(queue2)
	// queue2.RPush("4")
	// fmt.Println(queue2)
	// queue2.RPopAll()
	// fmt.Println(queue2)
	queue4 := new(Queue)
	queue4.LPush("1").RPush("2").LPush("3")
	fmt.Println(queue4)
	queue4.RPop().LPop()
	fmt.Println(queue4)
	queue4.LPush("1").RPush("2").LPush("3")
	fmt.Println(queue4)
	queue4.LPopAll()
	fmt.Println(queue4)
}
