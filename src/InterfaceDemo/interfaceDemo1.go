package main

import "fmt"

type empty interface {
	//所有类型都实现了空接口
}

type USB interface {
	//接口的签名
	Name() string //返回USB的名称
	Connecter
}
type Connecter interface {
	Connect()
}
type PhoneConnecter struct {
	name string
}

type TVconnecter struct {
	name string
}

func (tv TVconnecter) Connect() {
	fmt.Println("Connected:", tv.name)
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect", pc.name)
}

func main() {
	var a USB
	a = PhoneConnecter{"apple"}
	a.Connect()
	Disconnect(a)
	dis2(a)

	pc := PhoneConnecter{"PH"}
	var d Connecter
	d = Connecter(pc)
	d.Connect()

	pc.name = "pc"
	d.Connect()

	//类型转换,将usb转化为connecter
	var b Connecter
	b = Connecter(a)
	b.Connect()

	//TV类型
	c := PhoneConnecter{"apple123"}
	fmt.Println(c.name)
	c.name = "change"
	fmt.Println(c.name)

}

func Disconnect(usb USB) {
	//进行类型判断，为了取出name
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconnecting..", pc.name)
		return
	}
	fmt.Println("Unkonwn device")
}

func dis2(usb interface{}) {
	//使用了空接口，可以允许任何类型进行传递
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconnecting..", pc.name)
		return
	}
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Dsiconnected:", v.name)
	default:
		fmt.Println("Unknown devices")
	}
}
