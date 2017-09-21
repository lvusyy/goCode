package main

//select
import (
	"fmt"
)

func main() {
	c1, c2 := make(chan int), make(chan string)

	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					break
				}
				fmt.Println("c1 ", v)
			case v, ok := <-c2:
				if !ok {
					break
				}
				fmt.Println("c2 ", v)
			}
		}
	}()

	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "hello"

	close(c1)
	close(c2)
}
