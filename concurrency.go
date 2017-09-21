package main

//Goroutine 通过通信来共享内存,而非通过共享内存来通信
//Channel  阻塞同步

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 10)
	wg := sync.WaitGroup{}
	wg.Add(10)
	// c := make(chan bool,1) //缓存 1 个 异步
	// go func() {
	// 	fmt.Println("Go -  Go GO!!!")
	// 	c <- true
	// 	close(c)
	// }()

	// for v := range c {
	// 	fmt.Println(v)
	// }

	for i := 0; i < 10; i++ {
		go Go(c, i)
		go Go1(&wg, i)

	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		<-c
	}

}

func Go1(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println("sync ", index, a)
	wg.Done() //以指针方式调用, 否则是内存拷贝方式 无法同步状态
}

func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println("chan ", index, a)
	//if index == 9 {
	c <- true
	//}
}
