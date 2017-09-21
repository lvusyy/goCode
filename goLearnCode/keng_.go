package main

import (
	"fmt"
)

func main() {
	// slice_1_Y()
	slice_length_()
}

func append_slice_y(s *[]int) {
	//s是内存的拷贝
	fmt.Printf("s_point -> %p :%v\n", s, *s)
	*s = append(*s, 2) //
	fmt.Printf("s_point -> %p :%v\n", s, *s)
}

func append_slice(s []int) {
	//s是内存的拷贝
	fmt.Printf("s -> %p :%v\n", &s, s)
	s = append(s, 2) //
	fmt.Printf("s -> %p :%v\n", &s, s)
}

func slice_length_() {
	s := make([]int, 10)
	fmt.Printf("s -> %p :%v\n", &s, s)
	append_slice(s)
	fmt.Printf("s -> %p :%v\n", &s, s)
	/*
		s -> 0xc0420023e0 :[0 0 0 0 0 0 0 0 0 0]
		s -> 0xc042002440 :[0 0 0 0 0 0 0 0 0 0]
		s -> 0xc042002440 :[0 0 0 0 0 0 0 0 0 0 2]
		s -> 0xc0420023e0 :[0 0 0 0 0 0 0 0 0 0]
	*/

	fmt.Printf("s_point -> %p :%v\n", &s, s)
	append_slice_y(&s)
	fmt.Printf("s_point -> %p :%v\n", &s, s)
	/*
		s_point -> 0xc04204e3a0 :[0 0 0 0 0 0 0 0 0 0]
		s_point -> 0xc04204e3a0 :[0 0 0 0 0 0 0 0 0 0]
		s_point -> 0xc04204e3a0 :[0 0 0 0 0 0 0 0 0 0 2]
		s_point -> 0xc04204e3a0 :[0 0 0 0 0 0 0 0 0 0 2]
	*/
}

func slice_1_Y() {
	s := []string{"1", "2", "3"}
	for _, v := range s {
		go func(v string) { //可以是用参数传递,也就是值拷贝
			fmt.Println(v) //v 是值得拷贝了. for 中v变动 不会影响到这样
		}(v)
	}
	select {}
}

func slice_1() {
	s := []string{"1", "2", "3"}
	for _, v := range s {
		go func() { //可以是用参数传递,也就是值拷贝
			fmt.Println(v) //此 v是引用,当goroutine 执行到这里的时候 v的值也许已经变动了.
		}()
	}
	select {}
}
