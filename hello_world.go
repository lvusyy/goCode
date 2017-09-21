package main
	//say hello
import "fmt"




func main(){
	a,b :="sdfas",666 //精简赋值,推算式 只能在函数体内
	var x complex64 =6 +2i //复数计算
	fmt.Println(x*x)
	fmt.Println("hello go world!") //打印行
	fmt.Print("打印中文,测试中\n") //中文打印
	fmt.Printf("%s %d\n",a,b) //格式化打印

	m :=[10]int8{1,2,3,4,5}
	m[0]=1
	m[9]=10
	fmt.Printf("%v",m)
}