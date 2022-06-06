package design_mode

import "sync"

/*
 使用懒惰模式的单例模式 once 保证只执行一次
*/

type Singleton struct { // 单例模式
	name string
}

var once sync.Once
var singletonInstance *Singleton // 单例对象

func GetSingletonInstance() *Singleton {
	once.Do(func() {
		singletonInstance = &Singleton{"once"}
	})
	return singletonInstance
}
