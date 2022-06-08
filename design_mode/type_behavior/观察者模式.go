package design_mode

import "fmt"

/*
	观察者模式用于触发联动。
	一个对象的改变会触发其它观察者的相关动作，而此对象无需关心连动对象的具体实现。
*/

// 被观察对象
type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{name}
}

//实现观察者的通知接口
func (r *Reader) Update(s *Subject) {
	fmt.Println("reader update :", r.name, s.context)
}

// 观察者接口
type Observer interface {
	Update(*Subject)
}

//观察者对象
type Subject struct {
	observers []Observer
	context   string
}

func NewSubject() *Subject {
	return &Subject{observers: []Observer{}}
}

//被观察者对象注册通知
func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

//通知被观察者
func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

//更新内容后通知
func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}
