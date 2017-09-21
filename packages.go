package main

//每个GO程序都是由包组成的
//程序运行的入口是 `main` 包
import (
	"bufio"
	"fmt"
	"os"
)

//GOPATH 下src/是必要的.
//go install
//其中首字母大写的可以被导出,也就是可以被其它包访问.

func main() {

	inputReader := bufio.NewReader(os.Stdin)
	outputWriter := bufio.NewWriter(os.Stdout)
	fmt.Fprint(outputWriter, "输入你最喜欢的语言\n")
	outputWriter.Flush()
	// fmt.Println("输入你最喜欢的语言\n")
	language, _ := inputReader.ReadString('\n')
	fmt.Printf("我最喜欢的语言是:%s语言\n", language)
}
