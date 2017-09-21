package main

import (
	"fmt"
)

func main(){

	f:=closure(2)//返回的是一个函数
	fmt.Printf("%T    %p\n",f,&f)
	a:=f(3)//调用返回的函数,a是函数返回的结果
	b:=f(5)
	fmt.Printf("%T %d\n",a,a)
	fmt.Printf("%T %d\n",b,b)
/*
closure x -> 0xc04200e0d0 //
func(int) int    0xc042004028
noname x -> 0xc04200e0d0 // 
noname x -> 0xc04200e0d0 // x的地址都是一样的 
int 5
int 7
*/


fmt.Println("-----------case 2------参数拷贝-----------")
for i:=0;i<3;i++{
	defer func(li int){ //defer 入栈 后入先出...
		fmt.Println(li) //li是参数拷贝
	}(i)//()
}
fmt.Println("-----------case 3------参数引用-----------")
for i:=0;i<3;i++{
	defer func(){
		fmt.Println(i) //i是引用for的变量
	}()
}
}

//   函数名   参数     返回一个匿名函数 一个int类型参数 一个int类型返回值
func closure(x int ) func(int) int{
	fmt.Printf("closure x -> %p\n",&x)

	//	返回 一个函数 参数y int类型 返回值 int类型
	return func(y int) int{ 
		fmt.Printf("noname x -> %p\n",&x)	
		return x+y //返回int类型 因为x和y都是int类型
	}
}