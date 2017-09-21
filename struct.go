package main

import (
	"fmt"
)


type person struct{
	Name string
	Age int
}

func main(){
	 a:=person{}
	 //a:=person{Name:"lvusyy",Age:27}
	 a.Age=10
	 a.Name="john"
	 fmt.Println(a)

	 fmt.Println("***********A***********")
	 A(a)
	 fmt.Printf("A  %v\n",a)
	 fmt.Println("***********B**********")
	 B(&a)
	 fmt.Printf("B  %v\n",a)
	 fmt.Println("**********匿名结构**********")
	b:=struct{
		Id int
		Name string
		Contact struct{
			Phone,City string
		}
	}{Id:1,Name:"lvusyy"}
	// {Id:1,Name:"lvusyy",Contact:{Phone:"1213",City:"hangzhou"}}
	b.Contact.City="hangzhou"
	b.Contact.Phone="1351351351351"
	fmt.Println(b)
	fmt.Println("**********匿名字段**********")
	c:=struct{
		int
		string
	}{0,"hello"}
	
	fmt.Println(c,c.int,c.string)//这种也行直接对类型名

	fmt.Println("**********嵌套结构**********")
	d:=struct{
		id int
		level string
		person person //如果定义完整类型就无法 直接使用嵌套属性 如d.Name 只能d.person.Name
		//person      //这样的话,相当于匿名字段? 可以直接使用d.Name访问
	}{id:0,level:"普通会员",person:person{Name:"john",Age:15}}
	fmt.Println(d)
	d.person.Age=18
	fmt.Println(d)
}

func A(pers person){
	pers.Age=15
	fmt.Println("A -> set Age to 15")
}


func B(pers *person){//如果内部想修改 结构对象的值需要传递指针来修改. 直接传递是 对象的拷贝
	pers.Age=18
	fmt.Println("B -> set Age to 18")
}