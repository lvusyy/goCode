package main

import "fmt"

func main(){
	// var student map[string]float32=make(map[string]float32)
	//                   key类型 value类型
	student := make(map[string]float32 ) //简短方式
	student["zhangSan"]=18.8
	student["lisi"]=99.9
	delete(student,"lisi")
	fmt.Printf("%v\n",student)
	fmt.Printf("%.2f\n",student["lisi"])
 
}