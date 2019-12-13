package main

import (
	"fmt"
)

/**
1.每个变量都拥有地址，指针的值就是地址
2.对普通变量使用“&”操作符取地址获得这个变量的指针变量，指针变量的值就是指针地址，对指针变量使用"*"操作
   可以获得指针变量指向的原变量的值。
3."*"为右值时是取指向变量的值，“*”为左值时，是将值设置给指向的变量
4.创建指针的另一种方法-new()函数
   str := new (string)
   *str=""pengyu
	fmt.Println(*str)
 */

func swap(a,b *int){
	//取a指针的值，赋给临时变量t
	t := *a
	//取b指针的值，赋给a指针指向的变量
	*a = *b
	//将a指针的值赋给b指针指向的变量
	*b = t

}
func main() {

	x, y := 1,2
	swap(&x,&y)

	// 准备一个字符串类型
	var house = "Malibu Point 10880, 90265"

	// 对字符串取地址，ptr类型为*string
	ptr := &house

	// 打印ptr的类型
	fmt.Printf("ptr type: %T\n", ptr)

	// 打印ptr的指针地址
	fmt.Printf("address: %p\n", ptr)

	// 对指针进行取值操作
	value := *ptr

	// 取值后的类型
	fmt.Printf("value type: %T\n", value)

	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value)

}
