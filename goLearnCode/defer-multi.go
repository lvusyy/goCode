package main

import (
	"fmt"
)

func main() {

	fmt.Println("conting")
	
	for i:=0;i<10;i++{
		defer fmt.Printf("%d ",i) //defer 语句会被压人栈内,后进先出方式被调用.
	}

	fmt.Println("done")
//输出结果
/*
conting
done
9 8 7 6 5 4 3 2 1 0
*/

}