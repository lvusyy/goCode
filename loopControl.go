package main

import "fmt"
import "runtime"

func main(){

	//if 
	x:=1
	if x==1{
		fmt.Println("x is 1")
	}else{ 
		if x==3{
		fmt.Println("x is 3")
	}else{
		fmt.Println("else...")
		}
	}

	if v:=22;v>20{fmt.Print("v is great 20\n")} //if　可以在判断语句前面加一条语句 和for类似

	// for 的while用法
	s:=0
	for {
		fmt.Printf(`无限循环中的第%d次\n`,s+1) //``中\n的字符不会被转义
		s++
		if s==20{
			fmt.Printf("已经循环了%d次,现在退出循环\n",s)
			break
				}
		}
	//for 的普通用法
	for d:=0;d<2;d++{
		fmt.Println(d)
	}
	//for 的if用法
	for co:=0;co<22;{
		co++
		fmt.Println(co)
	}


	//switch

	switch os:=runtime.GOOS;os{//也可以像if那样在前面加一个语句
	case "darwin":
		fmt.Println("os x")
	case "likelinux":
		fallthrough //会继续执行Linux项下面的代码
	case "linux":
		fmt.Println("linux.")
	default:
		fmt.Println("Windows")
	}

	//switch 的if then else 版本
	os:=runtime.GOOS;
	switch {//也可以像if那样在前面加一个语句
	case os=="darwin":
		fmt.Println("os x")
	case os=="likelinux":
		fallthrough //会继续执行Linux项下面的代码
	case os=="linux":
		fmt.Println("linux.")
	default:
		fmt.Println("Windows")
	}
}