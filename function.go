package main

import (
	"fmt"
)

func main(){
a("hello",3,4,56)
str_:=[]string{"world","hello"}
fmt.Println("****************")
B(str_)
fmt.Println(str_)
fmt.Println("****************")
a,b,c,d,e:=1,2,3,4,5
C(a,b,c,d,e)
fmt.Printf("%p,%d\n|%d,%d,%d,%d,%d\n",&a,a,a,b,c,d,e)
fmt.Println("****************")
/*
****************
0xc042070030,1     //原来a
[1 0 3 4 5]
0xc0420540c8,1     //传递过来的a
|1,2,3,4,5
*/

}

//传递过来的 s
//        / | | \  //内存指向 
//       a  b  c d  //原 a,b,c,d,e内存的拷贝
//       1  2  3 4 //
func C(s ...int){
	s[1]=0
	fmt.Printf("%p,%d\n",&s[0],s[0])
	fmt.Println(s)
}

//传递的是指针形式 
func B(p []string){
	p[1]="c" //把p[1]的地址指向"c"的内存地址
	fmt.Println(p)
}

//不定参数长度
func a(b string,c ...int){
	fmt.Println(b,c)
}