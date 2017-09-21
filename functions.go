package main

import (
	"fmt"
)

//   函数名 参数名 类型     返回值类型
func add(x uint8,y uint8) uint16 {
//func add(x,y uint8) uint16 { //也可以缩写成这样
	return uint16(x+y)
}
//返回两个字符串
func swap(x,y string) (string,string) {
	return y,x
}
//Go 的返回值可以被命名，并且像变量那样使用。
//返回值的名称应当具有一定的意义，可以作为文档使用。
//没有参数的 return 语句返回结果的当前值。也就是`直接`返回。
func split(sum int) (x,y int){
	x=sum *8/10
	y=sum-x
	return
}

func main(){
	fmt.Println(add(23,44))
	fmt.Println(swap("world ","hello"))
	fmt.Println(split(22))
/*
67
hello world 
17 5
*/
}