package design_mode

import "fmt"

/*
	职责链模式用于分离不同职责，并且动态组合相关职责。
	Golang实现职责链模式时候，因为没有继承的支持，使用链对象包涵职责的方式，即：
	链对象包含当前职责对象以及下一个职责链。
	职责对象提供接口表示是否能处理对应请求。
	职责对象提供处理函数处理相关职责。
	同时可在职责链类中实现职责接口相关函数，使职责链对象可以当做一般职责对象是用。
*/

//接口定义
type Manager interface {
	HaveRight(money int) bool
	HandleFeeRequest(name string, money int) bool
}

//职责链实现类
type RequestChain struct {
	Manager
	successor *RequestChain
}

//设置下一个责任人
func (r *RequestChain) SetSuccessor(m *RequestChain) {
	r.successor = m
}

func (r *RequestChain) HandleFeeRequest(name string, money int) bool {
	if r.Manager.HaveRight(money) {
		return r.Manager.HandleFeeRequest(name, money)
	}
	if r.successor != nil {
		return r.successor.HandleFeeRequest(name, money)
	}
	return false
}

func (r *RequestChain) HaveRight(money int) bool {
	return true
}

//项目管理 链条
type ProjectManager struct{}

func NewProjectManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &ProjectManager{},
	}
}

func (*ProjectManager) HaveRight(money int) bool {
	return money < 500
}

func (*ProjectManager) HandleFeeRequest(name string, money int) bool {
	if name == "Tany" {
		fmt.Println(" project manager permit ", name, " money:", money)
		return true
	}
	fmt.Println(" project manager reject ", name, " money:", money)
	return false
}

//研发管理 链条
type DepManager struct{}

func NewDepManagerChain() *RequestChain {
	return &RequestChain{Manager: &DepManager{}}
}

func (*DepManager) HaveRight(money int) bool {
	return money < 1000
}

func (*DepManager) HandleFeeRequest(name string, money int) bool {
	if name == "anyT" {
		fmt.Println(" dep manager permit ", name, " money:", money)
		return true
	}
	fmt.Println(" dep manager reject ", name, " money:", money)
	return false
}

//常规管理 链条
type GeneralManager struct{}

func NewGeneralManagerChain() *RequestChain {
	return &RequestChain{Manager: &GeneralManager{}}
}

func (*GeneralManager) HaveRight(money int) bool {
	return true
}
func (*GeneralManager) HandleFeeRequest(name string, money int) bool {
	if name == "ayt" {
		fmt.Println(" general manager permit ", name, " money:", money)
		return true
	}
	fmt.Println(" general manager reject ", name, " money:", money)
	return false
}
