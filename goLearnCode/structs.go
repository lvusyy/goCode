package main

import (
	"fmt"
)

//结构体

type Vertex struct{
	X int
	Y int
}

func main(){


	var sv Vertex
	sv.X=22
	sv.Y=33
	fmt.Println(sv)
	fmt.Println(Vertex{1,2})
	pstruct:=&sv//指针指向结构体
	pstruct.X,pstruct.Y=pstruct.Y,pstruct.X //利用指针来交换数据
	fmt.Println(sv)
	
	// {22 33}
	// {1 2}
	// {33 22}

	var(
		v1=Vertex{1,2} //X:1 Y:2
		v2=Vertex{X:1} //Y:0 被省略
		v3=Vertex{} //X:0 Y:0
		p=&Vertex{2,3} //类型为 *Vertex
	)


}