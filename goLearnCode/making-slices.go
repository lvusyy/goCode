package main

import (
	"fmt"
)

func main(){
	a:=make([]int ,5)//len(a)=5 cap(a)=5
	printSlice("a",a)
	b:=make([]int,0,5)//len(a)=0 cap(a)=5
	printSlice("b",b)
	c:=b[:2]            //len=2 cap=5 [0 0] 
	printSlice(`c`,c)
	d:=c[2:5]           //len=3 cap=3 [0 0 0] 
	printSlice("d",d)


	f:=[...]string{"lvusyy","Penn"}

 

	cc:=make([]string,3,5)
	fmt.Printf("变量cc 地址:%p 内容:%v 长度:%d 容量:%d\n",cc,cc,len(cc),cap(cc))
	cc=append(cc,"s1","s2")
	fmt.Printf("变量cc 地址:%p 内容:%v 长度:%d 容量:%d\n",cc,cc,len(cc),cap(cc))
	cc=append(cc,"s3","s4") //超过容量之后就会从新分配内存用来存储. 从新分配的内存容量是原来的容量倍数
	fmt.Printf("变量cc 地址:%p 内容:%v 长度:%d 容量:%d\n",cc,cc,len(cc),cap(cc))
/*
变量cc 地址:0xc042076000 内容:[  ] 长度:3 容量:5
变量cc 地址:0xc042076000 内容:[   s1 s2] 长度:5 容量:5
变量cc 地址:0xc0420460a0 内容:[   s1 s2 s3 s4] 长度:7 容量:10
*/

}

func printSlice(s string,x []int){
	fmt.Printf("%s len=%d cap=%d %v \n",
	s,len(x),cap(x),x)
}