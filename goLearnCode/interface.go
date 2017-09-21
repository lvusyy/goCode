package main

import (
	"fmt"
	"math"
)
//接口是一个活多个方法(签名?)的集合

type Abser interface{
	Abs() float64
}
func main(){

	var a Abser
	f:=MyFloat(-math.Sqrt2)
	// v:=Vertex{3,4}

	a=f
	// a=&v

	// a=v

	fmt.Println(a.Abs())

}

type MyFloat float64

func (f MyFloat) Abs() float64{
	if f<0{
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct{
	X,Y float64
}