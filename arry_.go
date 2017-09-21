package main

import "fmt"

//slice 动态数组
	//var fslice []int
	//s := make([]int,10,20)
func main(){
	x := [10]int{1,2,3,4,5,6,7,8,9,10}

	y := x[1:3] //前闭后开 下标1开始,到3-1结束

	s:=make([]int,3)//分配内存
	fmt.Printf("%v\n%v length:%d\n",y,s,len(s))
	s=append(s,1,3,4,5)//如果s有足够的内存就利用原来的内存,如果不够就会从新申请一个内存地址
	fmt.Printf(`%v`,s) //``和"" 一样
}