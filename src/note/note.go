package main

import (
	"bytes"
	"fmt"
)

/**
批量声明变量
整型和浮点型默认值为0，字符串默认值为空字符串，切片、函数、指针默认为nil
整型变量注意事项：
unit8对应byte int16对应short int64对应long
类型别名的写法为：type TypeAlias = Type,例：
type NewInt int 或者 type IntAlias = int
 */
var (
	a int
	b string
	c []float32
	d func() bool
	e struct {
		x int
	  }
)
/**
字符串：
1.字符串取长度：使用len()可以用来后去切片、字符串、通道等的长度
字符串无汉语 len("pengyu")  有汉语 utf8.RuneCountInString("忍者haha")
2.遍历字符串
 无汉语 theme := "hello World"   for i=0;i<len(theme);i++ {}
 有汉语 theme := "佳期 beautifual" for _,s:=range theme {}
 */

/**
指针：
1.每个变量都拥有地址，指针的值就是地址
2.对普通变量使用“&”操作符取地址获得这个变量的指针变量，指针变量的值就是指针地址，对指针变量使用"*"操作
   可以获得指针变量指向的原变量的值。
3."*"为右值时是取指向变量的值，“*”为左值时，是将值设置给指向的变量
	func swap(a,b *int){
		//取a指针的值，赋给临时变量t
		t := *a
		//取b指针的值，赋给a指针指向的变量
		*a = *b
		//将a指针的值赋给b指针指向的变量
		*b = t

	}
4.创建指针的另一种方法-new()函数
   str := new (string)
   *str=""pengyu
	fmt.Println(*str)

 */


/**
切片
从数组或切片生成新的切片 slice [开始位置：结束位置]
切片最后一个元素使用slice[len(slice)]获取。
还可以使用make([] T,size,cap) ,如：a := make([]int,2,10)
使用make函数生成的切片一定发生了内存分配操作
append扩展切片
var car []string
car = append(car,"OldDriver")
car=append(car,"Ice","Sniper","Monk") //切片中添加元素
team := []string{"Pig","Flyingcake","Chicken"}
car = append(car,team...)//合并切片，...必须加，表示将team整个切片合并到car切片
复制切片元素到另一个切片  copy(destSlice,srcSlice []T) int
desData := make([]int,10)
copy(destData,team[1:])
从切片中删除
seq := []string{"a","b","c","d","e"}
index := 2 //指定删除位置
seq := append(seq[:index],seq[index+1:]...)
 */

/**
map:
map[keyType]ValueType
//第一种声明方式
scene := make(map[string]int)
scene["route"]=66
v := scene["route"]
for k,v := range scene{  } //遍历
for _,v := range scene{  }// 取值
for k := range scene{  }// 取键
//第二种
m := map[string]string{
	"name":"pengyu",
	"old":"30",
	"work":"java",
}
删除键值对：delete(m,"name")
并发环境中使用的map——sync.Map
func main(){
	var scene sync.Map
	//将键值对保存到sync.Map
	scene.Store("greece",97)
	scene.Store("london",100)
	scene.Store("egypt",200)
	//从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	//根据key删除
	scenne.Delete("london")
	//遍历
	scene.Range(func(k,v interface{})) bool {
		fmt.Println("iterate:",k,v)
		return true
	})
}
 */

/**
列表list
方式1  a := list.New()
方式2  var 变量名 list.List
a.PushBack("cannon") //插入列表的尾部
a.PushFront(67) // 67将放在first前面
element := a.PushBack("first") //尾部添加后保存元素句柄
a.InsertAfter("high",element)//在first之后添加high
a.InsertBefore("noon",element)//在first之前添加noon
a.remove(element)//删除first
遍历
for i:= a.Front(); i!= nil;i=a.Next(){
	fmt.Println(i.Value)
}
 */

/**
	流程控制
	if 表达式1{

	} else if 表达式2 {

	} else {

	}
	特殊写法：
	if err :=Connect();err != nil {
		fmt.Println(err)
		return
	}

	只有一个循环条件的循环
	var i int
	for i<=10 {
		i++
	}

	遍历通道
	c := make(chan int)
	go func(){
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
    for v := range c { fmt .Println(v)}

    switch分支
	var a= "hello"
	switch a {
		case "hello":
			fmt.Println(1)
			fallthrough
		case "world":
			fmt.Println(2)
		default:
			fmt.Println(0)
	}
case后也可以和if一样添加表达式
	var r int=11
	switch{
		case r >10 && r<20 :
			fmt.Println(r)
	}
	使用goto来实现统一错误处理
	err := firstCheckError()
	if err != nil {
		goto onExit
	}
	err = secondCheckError()
	if err != nil {
		goto onExit
	}
	fmt.Println("done")
	return
onExit:
	fmt.Println(err)
	exitProcess()
 */

/**
函数：
1.匿名函数用作回调函数
		func visit(list []int,f func(int)) {
			for _,v := range list {
				f(v)
			}
		}
		func main() {
			visit([]int{1,2,3,4},func(v int){
					fmt.Println(v)
			})
		}
 */


func main() {
	//字符串转义
	fmt.Println("str:=\"c:\\Go\\bin\\go.exe\"")
	//rune类型的b变量的实际类型是int32,对应的Unicode码就是20320
	/**
	Unicode是字符集，字符集为每个字符分配一个唯一的ID,我们使用到的所有字符在Unicode字符集中都有唯一的
	一个ID对应，UFT-8是编码规则，将Unicode中字符的ID以某种方式进行编码。UFT-8是一种变长编码规则，从1到
	4个字节不等。编码规则如下：
	Oxxxxxx表示文字符号0到127，兼容ASCII字符集
	从128到Ox10ffff表示其他字符。
	中文每个字符占用3个字节
	 */
	var b rune='你'
	fmt.Printf("%d %T\n",b,b)
	//创建切片
	a0:=make([]int ,3)
	a0[0]=1
	//3.除了+之外的，连接字符串
	hammer := "传我一锤"
	sickle := "死吧"
	//声明字节缓冲，把字符串写入缓冲
	var stringBuilder bytes.Buffer
	stringBuilder.WriteString(hammer)
	stringBuilder.WriteString(sickle)
	//将缓冲以字符串形式输出
	fmt.Println(stringBuilder.String())


	
}
