package main

import (
	"fmt"
)

func main() {
	fmt.Println("**************LABLE1***************")
	var v1 interface{} = 1
	var v2 interface{} = "abc"
	fmt.Println(v1, v2)
	fmt.pl
LABLE1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LABLE1 //跳出for无限循环
				//break //仅跳出第二层for而已
			} else {
				fmt.Println(i)
			}
		}
	}
	fmt.Println("**************LABLE2***************")
LABLE2:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			continue LABLE2 // 跳到第一层for里面
			// continue //只是跳过 continue LABLE2 之后的代码

		}
	}
	fmt.Println("**************LABLE3***************")
	// LABLE3:
	// 	for i:=0;i<10;i++{
	// 		for{
	// 			fmt.Println(i)
	// 			goto LABLE3//调整执行的位置
	// 		}
	// 	}
}
