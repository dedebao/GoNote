package main

import (
	"fmt"
)
/**
函数类型实现接口——把函数作为接口来调用
 */
// 调用器接口
type Invoker interface {
	// 需要实现一个Call方法，interface{}这种变量表示任意类型的值
	Call(interface{})
}

// 结构体类型
type Struct struct {
}

// 实现Invoker的Call
func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

//函数的声明不能直接实现接口，需要将函数定义为类型后，使用类型实现结构体，当类型方法被
//调用时，还需要调用函数本体
// 函数定义为类型，将func(interface{})定义为FuncCaller类型
type FuncCaller func(interface{})

// FuncCaller的Call()将实现Invoker的Call
func (f FuncCaller) Call(p interface{}) {

	// FuncCaller的Call()被调用与func(interface{})无关，还需要手动调用f函数本体
	f(p)
}

func main() {

	// 声明接口变量
	var invoker Invoker

	// 实例化结构体
	s := new(Struct)

	// 将实例化的结构体赋值到接口
	invoker = s

	// 使用接口调用实例化结构体的方法Struct.Call
	invoker.Call("hello")

	// 将匿名函数转为FuncCaller类型，再赋值给接口
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function", v)
	})

	// 使用接口调用FuncCaller.Call，内部会调用函数本体
	invoker.Call("hello")
}
