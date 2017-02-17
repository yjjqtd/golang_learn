package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	name string
}

func (this *MyStruct) GetName() string {
	return this.name
}
/*

*/
func main() {
	a := new(MyStruct)
	a.name = "yejianfeng"
	typ := reflect.TypeOf(a)
	fmt.Println(typ)
	fmt.Println("-------------------")
	var b MyStruct
	b.name = "abc"
	fmt.Println(reflect.TypeOf(b))
	
	var c MyStruct	
	c.name = "xiangli"
	fmt.Println(reflect.ValueOf(c).FieldByName("name").CanSet()) //false
	fmt.Println(reflect.ValueOf(&(c.name)).Elem().CanSet())      //true

	fmt.Println("--------------")
	var str string = "yejianfeng"
	p := reflect.ValueOf(&str)
	fmt.Println(p.CanSet())        //false
	fmt.Println(p.Elem().CanSet()) //true
	p.Elem().SetString("newName")
	fmt.Println(str)
}
/*
  输出结果
  *main.MyStruct
  -------------------
  main.MyStruct
*/
/*
1、TypeOf用来返回变量的类型，ValueOf用来返回变量的方法。
我们在初始化一个结构体的时候，有两种方式，var a Struct跟a :=new(Struct)，其中，前者的类型是struct，后者的类型是指针。

2、reflect.ValueOf(a).FieldByName方法
如果a是结构体，reflect.ValueOf(a).FieldByName("name")等价于a.name。
如果是指针的话ValueOf返回的是指针的Type，它是没有Field的，所以也就不能使用FieldByName

3、CanSet方法
CanSet当Value是可寻址的时候，返回true，否则返回false。
当前面的CanSet是一个指针的时候（p）它是不可寻址的，但是当是p.Elem()(实际上就是*p)，它就是可以寻址的
*/
