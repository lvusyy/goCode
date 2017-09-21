package main

import (
	"fmt"
)

//type 重新定义类型的都可以绑定多个方法.

type person struct{
	Name string
	Age uint8
}

type mint int

//函数
func main(){
	var p person	
	p.Name="alice"
	p.Age=15
	person.m(p,"yellow duck") //Method Expression
	fmt.Println(p.m("duck"))
	var ff mint
	ff=18
	fmt.Println(ff.hello("参数传递"),ff)
	ff.bye()          //方法1 Method Value
	(*mint).bye(&ff) //方法2  Method Expression
}

//   被绑定的类型 绑定的方法 参数  返回的类型
func (p person) m(s string) (status string, errno int){
	fmt.Println("p -> string ",s)
	fmt.Println("p -> ",p)
	return "ok",0
}

//以指针传递的参数可以被修改原值, 数值传递的就是值拷贝.不影响原值
func (m *mint ) hello(arg string)(string){
	fmt.Println("hello")
	*m=99
	fmt.Printf("参数:%s \n",arg)
	return "done"
}

//方法
func (m *mint ) bye(){
	fmt.Println("bye")
}