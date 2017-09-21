package main

import (
	"fmt"
)
 

func main(){
	var z []int //slice的零值是 `nil` 没有初始化
	fmt.Println(z,len(z),cap(z))
	if z == nil{
		fmt.Println("z is nil")
	}

	fmt.Println("--------------func A-----------")
	A()

}

func A(){
	defer func(){
		if err:=recover();err!=nil{//recovery 来恢复panic状态
			fmt.Println(err,"Recovey in B")
		}
	}()

	panic("Panic in B") //触发panic
}