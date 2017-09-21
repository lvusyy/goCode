package main

import (
	"fmt"
)

func main(){
	i,j:=33,2017
	p:=&i //取i地址复制给 指针p .数据类型 通过推导
	fmt.Println(*p) //*p 指针指向地址的数值
	*p=22
	fmt.Println(i)

	p=&j
	*p=*p/22
	fmt.Println(j)
/*
33
22
91
*/
	
}