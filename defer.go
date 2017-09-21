package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("world ") //defer 语句的参数会立即生成在上层函数返回之后调用

	fmt.Println("hello ")
}