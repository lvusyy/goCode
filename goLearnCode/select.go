package main

/*
 select 语句的语法：
每个case都必须是一个通信
所有channel表达式都会被求值
所有被发送的表达式都会被求值
如果任意某个通信可以进行，它就执行；其他被忽略。
如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。 否则：
如果有default子句，则执行该语句。
如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。
*/
import (
	"fmt"
)

func main() {

	c1, c2, c3 := make(chan int), make(chan int), make(chan int)
	var i1, i2 int

	go func() {
		i2 = 5525825
		c3 <- 2
		c1 <- 33
		fmt.Println(<-c2)
	}()

	for i := 0; i < 3; i++ { //进行3次select
		select {
		case i1 = <-c1:
			fmt.Println("received ", i1, " from c1\n")
		case c2 <- i2:
			fmt.Println("send ", i2, " to c2\n")

		case i3, ok := (<-c3):
			if ok {
				fmt.Println("received ", i3, " from c3\n")
			} else {
				fmt.Printf("c3 is closed\n")
			}
			// default: //变成异步,不会阻塞了.
			// 	fmt.Println("no communication\n")

		}
	}
}
