package main

// - -! 完全没搞懂go 下面都是猜的,请自行验证
// 通过反射来获取结构内数据 修改结构内数据 包括匿名数据 和调用函数 包括参数传递以及返回值获取
// 如果是匿名数据 reflect FieldByIndex方法获取到对象 然后操作之
//

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Manage struct {
	User
	title string
}

func main() {

	m := Manage{User: User{Id: 1, Name: "Fishe", Age: 1}, title: "hello world go dick!"}
	t := reflect.TypeOf(m)                           //观察类型 结构 相关信息
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1})) //取得User.Name types
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0}))    //Anonymous:true 代表匿名类型

	Set(m) //设置字段的值
	fmt.Println(m)
	//反射来调用方法
	v := reflect.ValueOf(m)       //获得反射值
	mv := v.MethodByName("Hello") //根据名称获得方法

	if !mv.IsValid() {
		fmt.Println("获取方法名错误")
		return
	}

	//配置参数
	// args := []reflect.Value{reflect.ValueOf("monkey")}
	args := []reflect.Value{reflect.ValueOf("monkey")}
	fmt.Println(mv.Call(args)[0])
}

func (u User) Hello(name string) (v uint) {
	fmt.Println("Hello ", name, " my name is ", u.Name)
	v = 18
	return
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() != reflect.Ptr || !v.Elem().CanSet() { //reflect.Valueof 和 v.Elem() 返回的都是value 类型
		fmt.Println("参数不是指针类型.或 cannot set")
		return
	} else {
		v = v.Elem() //类型一样
	}

	//根据字段名来修改 数值
	f := v.FieldByName("Name") //根据字段名 来返回
	if !f.IsValid() {
		fmt.Println("没有找到Name字段,无法设置")
		return
	}
	if f.Kind() == reflect.String { //根据类型判断操作方法
		f.SetString("修改以后的值")
	}
}
