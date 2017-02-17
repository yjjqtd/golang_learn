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
  TypeOf用来返回变量的类型，ValueOf用来返回变量的方法。
  我们在初始化一个结构体的时候，有两种方式，var a Struct跟a :=new(Struct)，其中，前者的类型是struct，后者的类型是指针。
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
}
/*
  输出结果
  *main.MyStruct
  -------------------
  main.MyStruct
*/
