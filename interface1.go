package main

import "fmt"

//空接口
type empty interface {
}

type USB interface {
	Name() string
	Connecter //嵌入式接口
}

type Connecter interface {
	Connect()
}

//PhoneConnecter 实现了USB所有方法
type PhoneConnecter struct {
	name string
}

func (phone *PhoneConnecter) Name() string {
	return phone.name
}

func (phone *PhoneConnecter) Connect() {
	fmt.Println(phone.Name(), " is connect")
}

func Disconnect(usb interface{}) {
	// if pc,ok:=usb.(PhoneConnecter);ok{ //类型判断
	// 	fmt.Println(usb.Name(),"is Disconnect")
	// }
	switch v := usb.(type) {
	case USB:
		fmt.Println(v.Name(), "is Disconnect")
	default:
		fmt.Println("unkonw devices", "is Disconnect")

	}

}

func main() {
	PC := &PhoneConnecter{name: "HuaWei"}
	PC.Connect()
	Disconnect(PC)

	PC1 := &PhoneConnecter{}
	PC1.name = "APPLE"
	PC1.Connect() //也可以直接这样调用
	USB.Connect(PC1)
	Disconnect(PC1)
}
