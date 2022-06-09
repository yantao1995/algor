package design_mode

import (
	design_behavior "algor/design_mode/type_behavior"
	design_create "algor/design_mode/type_create"
	design_struct "algor/design_mode/type_struct"
	"fmt"
	"testing"
)

func TestTypeBehavior(t *testing.T) {
	return
	//模板方法 (父类提供并实现模版方法,子类选择性实现模板方法,继承父类方法)
	var downloaderHttp design_behavior.Downloader = design_behavior.NewHTTPDownloader()
	downloaderHttp.Download("http source")
	var downloaderFtp design_behavior.Downloader = design_behavior.NewFTPDownloader()
	downloaderFtp.Download("ftp source")

	//迭代器 (迭代目标对象的索引范围值，实现对任意对象的迭代)
	var aggregate design_behavior.Aggregate
	aggregate = design_behavior.NewNumber(1, 10)
	design_behavior.IteratorPrint(aggregate.Iterator())

	//命令 (将目标对象的方法封装到对象中，来进行调用)
	mb := &design_behavior.MainBoard{}
	startCommand := design_behavior.NewStartCommand(mb)
	rebootCommand := design_behavior.NewRebootCommand(mb)
	box1 := design_behavior.NewBox(startCommand, rebootCommand)
	box1.PressButton1()
	box1.PressButton2()
	box2 := design_behavior.NewBox(rebootCommand, startCommand)
	box2.PressButton1()
	box2.PressButton2()

	//观察者   (被观察者注册到观察者持有，然后观察者更新时，挨个通知被观察者)
	reader1 := design_behavior.NewReader("reader1")
	reader2 := design_behavior.NewReader("reader2")
	subject := design_behavior.NewSubject()
	subject.Attach(reader1)
	subject.Attach(reader2)
	subject.UpdateContext("通知")

	//中介者   (封装了对象直接的交互关系)
	mediator := design_behavior.GetMediatorInstance()
	mediator.CD = &design_behavior.CDDriver{}
	mediator.CPU = &design_behavior.CPU{}
	mediator.CD.ReadData()
	fmt.Println(mediator.CD.Data, mediator.CPU.Sound, mediator.CPU.Video)
}
func TestTypeStruct(t *testing.T) {
	//桥  (抽象类和实现类的组合)
	m := design_struct.NewCommonMessage(design_struct.ViaSMS())
	m.SendMessage("test via sms", "zhang san")
	m2 := design_struct.NewCommonMessage(design_struct.ViaEmail())
	m2.SendMessage("test via email", "li si")
	m3 := design_struct.NewUrgencyMessage(design_struct.ViaSMS())
	m3.SendMessage("test urgency via sms", "zhang si")
	m4 := design_struct.NewUrgencyMessage(design_struct.ViaEmail())
	m4.SendMessage("test urgency via email", "li san")

	//装饰  (继承，多态,不会调用父类方法)
	var decorator design_struct.ComponentDecorator = &design_struct.ComponentDecoratorImpl{}
	fmt.Println(decorator.Calc())
	decorator = design_struct.NewAddDecoratorImpl(decorator, 10)
	fmt.Println(decorator.Calc())
	decorator = design_struct.NewMulDecoratorImpl(decorator, 8)
	fmt.Println(decorator.Calc())
	decorator = design_struct.NewAddDecoratorImpl(decorator, 8)
	fmt.Println(decorator.Calc())

	//享元	(共享对象而不复制对象)
	viewer := design_struct.NewImageViewer("imageViewer")
	viewer.Display()
	viewer2 := design_struct.NewImageViewer("imageViewer")
	fmt.Println(viewer.ImageFlyweight == viewer2.ImageFlyweight)

	//组合  (使用相同的接口来调用对象集)
	root := design_struct.NewComponent(design_struct.CompositeNode, "root")
	n1 := design_struct.NewComponent(design_struct.CompositeNode, "n1")
	n2 := design_struct.NewComponent(design_struct.CompositeNode, "n2")
	n3 := design_struct.NewComponent(design_struct.CompositeNode, "n3")
	l1 := design_struct.NewComponent(design_struct.LeafNode, "l1")
	l2 := design_struct.NewComponent(design_struct.LeafNode, "l2")
	l3 := design_struct.NewComponent(design_struct.LeafNode, "l3")
	root.AddChild(n1)
	root.AddChild(n2)
	n1.AddChild(n3)
	n1.AddChild(l1)
	n2.AddChild(l2)
	n2.AddChild(l3)
	root.Print("prev ")

	//代理  (调用对象方法前后进行处理)
	var sub design_struct.Subject = &design_struct.Proxy{}
	fmt.Println(sub.Do())

	//适配器   (实现接口调用对应需要适配的接口)
	adaptee := design_struct.NewAdaptee()
	target := design_struct.NewTarget(adaptee)
	fmt.Println(target.Request())

	//外观   (对外提供一个的接口，包含多个接口的方法的处理结果)
	api := design_struct.NewApi()
	fmt.Println(api.Test())

	//选项  （初始化时为对象的每个属性赋值）
	struct_op := design_struct.NewOption(design_struct.IdOpFunc(1), design_struct.NameOpFunc("name"))
	fmt.Println(struct_op)
}

func TestTypeCreate(t *testing.T) {
	//简单工厂	(根据传参生成不同对象)
	simpleFactoryApi1 := design_create.NewSimpleFactoryApi(1)
	simpleFactoryApi2 := design_create.NewSimpleFactoryApi(2)
	simpleFactoryApi3 := design_create.NewSimpleFactoryApi(3)
	fmt.Println(simpleFactoryApi1.Say("simple1"), simpleFactoryApi2.Say("simple2"), simpleFactoryApi3)

	//工厂方法	(类似于子类继承父类方法)
	var computeNormalFactoryPrint = func(factory design_create.OperatorFactory, a, b int) {
		op := factory.CreateOperator()
		op.SetA(a)
		op.SetB(b)
		fmt.Println(op.GetResult())
	}
	normalPlusFactory := design_create.PlusOperatorFactory{}
	normalMinusFactory := design_create.MinusOperatorFactory{}
	computeNormalFactoryPrint(normalPlusFactory, 1, 2)
	computeNormalFactoryPrint(normalMinusFactory, 3, 1)

	//抽象工厂方法	(对象要有关联，用于生成产品族的工厂，否则退化成工厂方法)
	var getMainAndDetail = func(factory design_create.OrderDaoAbstractFactory) {
		factory.CreateOrderMainDao().SaveOrderMain()
		factory.CreateOrderDetailDao().SaveOrderDetail()
	}
	rdbFactory := &design_create.RDBOrderDaoAbstract{}
	getMainAndDetail(rdbFactory)
	xmlFactory := &design_create.XMLOrderDaoAbstract{}
	getMainAndDetail(xmlFactory)

	//创建者  (复杂对象分离成多个简单对象构建组合)
	builder := design_create.Builder1{}
	director := design_create.NewDirector(&builder)
	director.ConstructDirector()
	fmt.Println(builder.GetResult())

	//原型  (克隆对象)
	prototypeManager := design_create.NewPrototypeManager()
	a := &design_create.TestStruct{Id: 10}
	prototypeManager.Set("A", a)
	a1 := prototypeManager.Get("A")
	b := a1.Clone()
	fmt.Println(a == b, b.(*design_create.TestStruct).Id)

	//单例	(只初始化一次)
	singleton := design_create.GetSingletonInstance()
	fmt.Println(singleton)
}
