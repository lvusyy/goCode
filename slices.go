package main

import (
	"fmt"
)

func main(){

	p:=[]int{2,3,5,7,11,13}
	fmt.Println("p===",p)


	for i :=0;i<len(p);i++{
		fmt.Printf("p[%d]==%d\n",i,p[i])
	}

	fmt.Println(p[1:4])//从下标1开始 到下标4-1结束
	fmt.Println(p[:4])//从下标0开始 到下标4-1结束
	fmt.Println(p[4:])//从下标4开始 到下标len(s)结束
	
}