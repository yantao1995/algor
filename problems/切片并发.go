package problems

import (
	"encoding/json"
	"fmt"
)

/*
 总结：
	切片
	 指针类型可以并发读写
	 非指针类型会 panic 比如  []int
	 但是 append 扩容导致数据不一致
*/

type P struct {
	A0 *[]*[]int //行
	A1 *[][]int  //不行
	A2 *[]int    //行
	A3 *[]int    //不行
}

func TestRace() {
	p := &P{
		A0: &[]*[]int{&[]int{1}},
		A1: &[][]int{[]int{1}},
		A2: &[]int{1},
		A3: &[]int{1},
	}
	bts, _ := json.Marshal(p)
	fmt.Println(string(bts))
	go func() {
		for {

			bts := []byte(`{"A0":[[1,2,3]],"A1":[[1,2,3]],"A2":[1,2,3],"A3":[1,2,3],"m":{"1":1,"2":2}}`)
			if err := json.Unmarshal(bts, p); err != nil {
				panic(err)
			}

			a := *p.A0
			aa := *a[0]
			aa[0] = 2

			bts2 := []byte(`{"A0":[[1,2,3,4,5,6,7]],"A1":[[1,2,3,4,5,6,7]],"A2":[1,2,3,4,5,6],"A3":[1,2,3,4,5,6],"m":{}}`)
			if err := json.Unmarshal(bts2, p); err != nil {
				panic(err)
			}

			a = *p.A0
			aa = *a[0]
			aa[0] = 2
		}
	}()
	for {
		a := *p.A0
		aa := *a[0]
		fmt.Println(aa)
		for k := range aa {
			_ = aa[k]
		}
	}
}
