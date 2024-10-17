package problems

import (
	"encoding/json"
	"fmt"
	"testing"
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

// 并发
func TestSliceRace(t *testing.T) {
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

// 切片拷贝
func TestSliceCopy(t *testing.T) {
	//深拷贝
	asz := [][1]int{
		{1},
		{4},
	}

	bsz := append([][1]int(nil), asz...)
	btssz := `[[11],[44]]`
	json.Unmarshal([]byte(btssz), &asz)
	fmt.Println(bsz)

	fmt.Println("-----------------------------")

	//浅拷贝
	a := [][]int{
		{1},
		{4},
	}

	b := append([][]int(nil), a...)
	bts := `[[11],[44]]`
	json.Unmarshal([]byte(bts), &a)
	fmt.Println(b)

	fmt.Println("-----------------------------")

	//单数组深拷贝
	a1 := []int{
		1,
	}

	b1 := append([]int(nil), a1...)
	a1[0] = 11
	// bts1 := `[11]`
	// json.Unmarshal([]byte(bts1), &a1)
	fmt.Println(b1)

	fmt.Println("-----------------------------")

	//单数组深拷贝
	a2 := []int{
		1,
	}
	b2 := make([]int, 0, 10)
	b2 = append(b2, a2...)
	a2[0] = 11
	bts2 := `[111]`
	json.Unmarshal([]byte(bts2), &a2)
	fmt.Println(a2, b2)
}
