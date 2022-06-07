package design_mode

import "fmt"

/*
	桥接模式分离抽象部分和实现部分。使得两部分独立扩展。
	桥接模式类似于策略模式，区别在于策略模式封装一系列算法使得算法可以互相替换。
	策略模式使抽象部分和实现部分分离，可以独立变化。
*/

//消息发送抽象接口
type AbstractMessage interface {
	SendMessage(text, to string)
}

//消息发送接口
type MessageInter interface {
	Send(text, to string)
}

//经过 SMS
type MessageSMS struct{}

func ViaSMS() MessageInter {
	return &MessageSMS{}
}

func (*MessageSMS) Send(text, to string) {
	fmt.Println(text, " ---> ", to)
}

//经过 Email
type MessageEmail struct{}

func ViaEmail() MessageInter {
	return &MessageEmail{}
}

func (*MessageEmail) Send(text, to string) {
	fmt.Println(text, " ---> ", to)
}

//消息发送 通常 SMS
type CommonMessage struct {
	method MessageInter
}

func NewCommonMessage(method MessageInter) AbstractMessage {
	return &CommonMessage{method}
}

func (m *CommonMessage) SendMessage(text, to string) {
	m.method.Send(text, to)
}

//消息发送 紧急|异常 Email
type UrgencyMessage struct {
	method MessageInter
}

func NewUrgencyMessage(method MessageInter) AbstractMessage {
	return &UrgencyMessage{method}
}

func (m *UrgencyMessage) SendMessage(text, to string) {
	m.method.Send("Urgency: "+text, to)
}
