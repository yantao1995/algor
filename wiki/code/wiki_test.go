package code

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// 单个P的调度逻辑
func TestSchedule(t *testing.T) {
	/*
		4
		0
		1
		2
		3
		如何i超过 257，就会加到 全局 runq 了
		单个P 记录 本地 runq 和 runnext g
		runext记录0, 当1进来的时候， 0就被挤到 本地runq了......
		如果本地 runq 已经放了 256个g，准备放 第257个的时候，就会触发放到全局runq的操作，
		会拿前一半g + 当前待入本地runq 的g 一起进全局runq 。


		调度逻辑 ： 本地队列 p.schedtick%61 ==0 ，每执行61个g，就会从全局队里里面拿一个执行。
			（避免每个本地runq繁忙的时候，导致全局runq迟迟得不到调度）

	*/

	runtime.GOMAXPROCS(1)
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second * 3)
}
