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
var (
	a int
	b string
	c []float32
	d func() bool
	e struct {
		x int
	  }
)
 */

/**
字符串：
1.字符串取长度：使用len()可以用来后去切片、字符串、通道等的长度
字符串无汉语 len("pengyu")  有汉语 utf8.RuneCountInString("忍者haha")
2.遍历字符串
 无汉语 theme := "hello World"   for i=0;i<len(theme);i++ {}
 有汉语 theme := "佳期 beautifual" for _,s:=range theme {}
3.字符串是常量，可以通过数组下表访问，但不能通过下表进行修改
4.字符串类型底层实现是一个二元的数据结构，一个是指针指向字节数组的起点，另一个是长度。
	type string Struct{
		str unsafe.Pointer //指向底层字节数组的指针
		len int			//字节数组长度
	}
5.基于字符串创建的切片和原字符串指向相同的底层字符数组，一样不能修改，对字符串的切片
操作返回的子串仍然是string,而非slice.
6.字符串和切片的转换：字符串可以转换为字节数组，也可以转换为Unicode的字数组。
	a := "hello,世界"
	c:=[]rune(a)
    b:=[]byte(a)
7.通过“+”进行连接
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
不要直接修改map value内某个元素的值，如果想修改map的某个键值，则必须整体赋值。
type User struct{
	name string
	age int
}
ma := make(map[int]User)
andes := User{
	name:"Tom",
	age:18,
}
ma[1]=andes
//ma[1].age=19  //Error,不能通过map引用直接修改
andes.age=19
ma[1]=andes //必须整体替换Value
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
2.使用匿名函数实现操作封装
	var skillParam=flag.String("skill","","skill to perform")
	func main(){
		flag.Parse()
		var skill = map[string]func(){
			"fire": func(){
				fmt.Println("chicken fire")
			},
			"run" : func(){
				fmt.Println("soldier run")
			},
			"fly" : func(){
				fmt.Println("angel fly")
			}
		}
		if f,ok := skill[*skillPrama]: ok{
			f()
		} else {
			fmt.Println("skill not found")
		}
	}
运行：go rum main.go --skill=fly
3.函数中允许返回局部变量的地址：
	Go编译器使用“栈逃逸”机制将局部变量的空间分配在堆上。如：
	func sum(a,b  int) *int{
		sum := a+b
		return &sum  //允许，sum会分配在heap上
	}
4.  函数的调用者负责准备环境，包括为参数和返回值开辟栈空间。
	寄存器的保存和恢复也由调用方负责。
	函数调用后回收栈空间，恢复BP也由主调函数负责。
	函数的多指返回实质上是在栈上开辟多个地址分别存放返回值，这个并没有什么特别的地方。
	如果返回值是存放到堆上的，则多了一个复制的动作。
	函数调用前已经为返回值和参数分配了栈空间，分配顺序是从右向左的，先是返回值，然后是
	参数，通用的栈模型如下：
	-------
	返回值y
	-------
	返回值x
	-------
	参数b
	-------
	参数a
	-------
	函数的多值返回时主调函数预先分配好空间来存放返回值，被调函数执行时将返回值复制到
	该返回位置来实现的。

5.无论接收者是什么类型，方法和函数的实参传递都是值拷贝。如果接收者是值类型，则传递的是值的副本；
如果接收者是指针类型，则传递的是指针的副本。
type Int int
func (a Int) Max(b Int) Int{
	if a>=b{
		return a
	}else{
		return b
	}
}
func (i *Int) Set(a Int){
	*i = a
}
func (i Int) Print(){
	fmt.Printf("value=%d\n",i)
}
func main(){
	var a Int=10
	var b Int=20
	c := a.Max(b)
	c.Print() //value=50
	(&c).Print() //value=50,内部被编译器转换为c.Print()
	a.Set(20) // 内部被编译器转化为(&a).Set(20)
	a.Print() //value=30
}
总结：
T类型的方法集是S
*T类型的方法集是S和*S
这里定义了一个新类型Int,新类型的底层类型是int,Int虽然不能继承int的方法，但底层类型支持的操作(
算数运算和赋值运算)可以被上层类型继承。

通过类型变量进行值调用和表达式调用，在这种情况下，使用值调用方式调用时编译器会自动转换，使用表达式
调用方式调用时编译器不会进行转换，会进行严格方法集检查
如：
type Data struct{}
func (Data) TestValue(){}
func (*Data) TestPointer(){}
//声明一个类型变量a
var a Data= struct {}{}
Data.TestValue(a)
Data.TestValue(&a)
Data.TestPointer(&a) //报错
(*Data).TestPointer(&a)
//值调用编译器会自动转换
f := a.TestValue
f()
y := (&a).TestValue //编译器帮助转换为a.TestValue
y()
g := a.TestPointer //会转换为(&a).TestPointer
g()
x :=(&a).TestPointer
x()

Go语言组合方法集
若类型S包含匿名字段T,则S的方法集包含T的方法集
若类型S包含匿名字段*T,则S的方法集包含T和*T方法集
不管类型S中嵌入的匿名字段是T还是*T,*S方法集总是包含T和*T方法集。
例：
type X struct{ a int }
type Y struct{ X }
type Z struct{ *X }
func (x X) Get() int{ return x.a }
func (x *X) Set(i int) { x.a=i }
func main(){
	x := X {a:1}
	y := Y { X:x,}
	println(y.Get())  //1
	//此处编译器做了自动转换
	y.Set(2)
	println(y.Get()) //2
	//为了不让编译器做自动转换，使用方法表达式调用方式
	//Y内嵌字段X,所以type Y的方法集是Get,type *Y的方法是Set Get
	(*Y).Set(&y,3)
	//type Y的方法集合并没有Set方法，所以下一句编译不能通过
	//Y.Set(y,3)
	println(y.Get()) //3
	z :=Z{
		X:&x,
	}
	//按照嵌套字段的方法集的规则
	//Z内嵌字段*X,因此type Z和type *Z方法集都包含类型X定义的方法Get和Set
	//为了不让编译器做自动转换，仍然使用方法表达式调用方式
	Z.Set(z,4)
	println(z.Get())//4
	(*Z).Set(&z,5)
	println(z.Get())//5
}

inject是Go语言依赖注入的实现，它实现了对结构(struct)和函数的依赖注入。

传值还是传引用？
值拷贝：
(1)函数参数传递时使用的是值拷贝。
(2)实例赋值给接口变量，接口对实例的引用是值拷贝。
明明是值拷贝却修改了变量的内容，有以下两种情况：
(1)直接传递的是指针。指针传递同样是值拷贝，但指针和指针副本的值指向的地址是一个地方，所以能修改实参值。
(2)参数是复合数据类型，这些复合数据类型内部有指针类型的元素，此时参数的值拷贝并不影响指针的指向。
Go复合类型中chan、map、slice、interface内部都是通过指针指向具体的数据，这些类型的变量在作为函数参数传递时，
实际上相当于指针的副本。
/*
并发：
time.Sleep(5 * time.Second)
//返回当前程序的goroutine数目
runtime.NumGoroutine()
go后面的函数的返回值会被忽略
调度器不能保证多个goroutine的执行次序。
没有父子goroutine的概念，所有的goroutine是平等地被调度和执行的
Go程序在执行时会单独为main函数创建一个goroutine，遇到其他go关键字时再去创建其他的goroutine.
Go没有暴露goroutine_id给用户，所以不能在一个goroutine里面显式地操作另一个goroutine，不过runtime包
提供了一些函数访问和设置goroutine的相关信息。
func GOMAXPROCS(n int) int 用来设置或查询可以并发执行的goroutine数目，n大于1表示设置GOMAXPROCS值，
否则表示查询当前的GOMAXPROCS值。
获取当前的GOMAXPROCS值=====runtime.GOMAXPROCS(0)
设置GOMAXPROCS的值为2======runtime.GOMAXPROCS(2)
func Goexit()结束当前goroutine的运行，Goexit在结束当前goroutine运行之前会调用当前goroutine已经注册
的defer,Goexit不会产生panic，所以该goroutine defer里面的recover调用都返回nil。
func Gosched()是放弃当前调度执行机会，将当前goroutine放到队列中等待下次被调度。
 */


/*
defer:
1.使用defer延迟并发解锁
var (
	valueByKey = make(map[string]int)
	valueByKeyGuard sync.Mutex
)
func readValue(key string)int{
	//对共享资源加锁
	valueByKeyGurad.Lock()
	//释放共享资源
	defer valueByKeyGuard.Unlock()
	v := valueByKeyGuard[key]
	return v
}
2.defer后面必须是函数或方法的调用，不能是语句。
3.defer函数的实参是在注册时通过值拷贝的方式传递进去的，不是执行的时候才传进去的。
4.defer语句必须先注册才能执行，如果defer位于return之后，则defer因为没有注册，不会执行
5.主动调用os.Exit(int)退出进程时，defer将不再被执行(即使defer已经提前注册)
6.在打开资源无报错后直接调用defer关闭资源，如：
	src,err := os.Open(src)
	if err != nil {
		return
	}
	defer src.Close()
7.recover()用来捕获panic,阻止panic继续向上传递。recover()和defer一起使用，但是recover()
只有在defer后面的函数体内被直接调用才能捕获panic终止异常，否则返回nil，异常继续向外传递。
捕获失败：
defer func(){
	func(){
		println("defer inner")
		recover()//无效
	}()
}()
捕获成功
defer func(){
	println("defer inner")
	recover()
}()
8.包中init函数引发的panic只能在init函数中捕获，在main中无法被捕获，原因是init函数先于
main执行。
9.函数并不能捕获内容新启动的goroutine所抛出的panic。

*/

/**
结构体：
1.基本的实例化
	var ins T
	ins.属性key1=value1
	ins.属性key2=value2
	结构体成员变量的赋值方法与普通变量一致。
1.创建指针类型的结构体：
	ins := new(T)
	T为类型，可以使结构体、整型、字符串等。
	ins: T类型被实例化后保存到ins变量中，ins的类型为*T，属于指针
2.取结构体的地址实例化
	ins := &T{}
	T表示结构体类型
	ins为结构体的实例，类型为*T,是指针类型。
使用键值对填充结构体
	type People struct{
		name string
		child *People
	}

	relation := &People{
		name: "爷爷",
		child: &People{
			name:"爸爸",
			child: &People{
				name :"我",
			},
		}
	}

示例2：
	type Person struct{
		Name string
		Age int
	}
	type Student struct{
		Person *Person
		Number int
	}
	p := &Person{
		Name:"Pengyu",
		Age:12,
	}
	s := Student{
		Person:p,
		Number:110,
	}
多个值列表初始化结构体的例子
type Adress struct{
	Province string
	City string
	Zipcode int
	PhoneNumber string
}

addr := Adress{
	"四川",
	"成都",
	61000,
	"0",
}

模拟构造函数
type Cat struct{
	Color string
	Name string
}

func GetInstance(name string,color string) *Cat{
	return &Cat{
		Name:name,
		Color:color
	}
}

接收器：方法作用的目标
func (接收器变量 接收器类型) 方法名(参数列表) (返回列表){ 函数体 }
type Property struct{
	value int
}
func (p *Property) SetValue(v int) {
	p.value=v
}
func (p *Property) GetValue() int{
	return p.value
}

type Cat struct {
	name string
	Animal
}

func (cat Cat) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)",
		cat.scientificName, cat.Animal.AnimalCategory, cat.name)
}
func (cat *Cat) SetName(name string) {
	cat.name = name
}
那么值方法和指针方法之间有什么不同点呢？
（1）值方法的接收者是该方法所属的那个类型值的一个副本。我们在该方法内对该副本的修改一般都不会体现在原值上，除非这个类型本身是某个引用类型（比如切片或字典）的别名类型。而指针方法的接收者，是该方法所属的那个基本类型值的指针值的一个副本。我们在这样的方法内对该副本指向的值进行修改，却一定会体现在原值上。
（2）一个自定义数据类型的方法集合中仅会包含它的所有值方法，而该类型的指针类型的方法集合却囊括了前者的所有方法，包括所有值方法和所有指针方法。
严格来讲，我们在这样的基本类型的值上只能调用到它的值方法。但是，Go语言会适时地为我们进行自动地转译，使得我们在这样的值上也能调用到它的指针方法。
比如，在Cat类型的变量cat之上，之所以我们可以通过cat.SetName("monster")修改猫的名字，是因为Go语言把它自动转译为了(&cat).SetName("monster")，即：先取cat的指针值，然后在该指针值上调用SetName方法。
（3）在后边你会了解到，一个类型的方法集合中有哪些方法与它能实现哪些接口类型是息息相关的。如果一个基本类型和它的指针类型的方法集合是不同的，那么它们具体实现的接口类型的数量就也会有差异，除非这两个数量都是零。
比如，一个指针类型实现了某某接口类型，但它的基本类型却不一定能够作为该接口的实现类型。

 */

/**
接口：
一个类型可以实现多个接口，多个类型可以实现相同的接口
type Service interface{
	Start()
	Log(string)
}
type Logger struct{}
func (g *Logger) Log(log string){}
type GameService struct{}
func (g *GameService) Start(){}
 */

/**
包：
  导入包后自定义引用的包名
		customName "path/to/package"
   如：
		import (
			renameLib "chapter08/importadd/mylib"
			"fmt"
		)
	匿名导入包——只导入包但不使用包内类型和数值
	import(
		_ "path/to/package"
	)
	包的注意事项：
(1)一个包可以有多个init函数，包加载会执行全部的init函数，但并不能保证执行顺序，所以不建议在一个包中放入多个
init函数，将需要初始化的逻辑放到一个init函数里面。
(2)包不能出现环形引用。比如包a引用了包b，包b引用了包c，如果包c又引用了包a,则编译不能通过。
(3)包的重复引用是允许的。比如包a引用了包b和包c,包b和包c都引用了包d。这种场景相当于重复引用了d，这种情况是允许
的，并且go编译器保证d的init函数只会执行一次。

在执行main.go之前，Go引导程序会先对整个程序的包进行初始化。
Go包的初始化有如下特点：
(1)包初始化程序从main函数引用的包开始，逐级查找包的引用，直到找到没有引用其他包的包，最终生成一个包引用的有向无环图。
(2)Go编译器会将有向无环图转换为一棵树，然后从树的叶子节点开始逐层向上对包进行初始化。
(3)单个包的初始化过程，先初始化常量，然后是全局变量，最后执行包的init函数(如果有)。
 */

/**
通道
通道发送数据的格式： 通道变量 <- 值
chan := make(interface{})
chan <- 0
chan <- "hello"
发送将持续阻塞直到数据被接收 data := <-chan
通道的数据接收可以借用for range语句进行多个元素的接收操作
for data:= range chan{

}
遍历的结果就是接收到的数据，数据类型就是通道的数据类型，通过for遍历获得的变量只有一个，即data.
单向通道：
	只能发送通道：var 通道实例 chan<- 元素类型
	只能接收通道：var 通道实例 <-chan 元素类型
带缓冲的通道：
	缓冲通道在发送时无需等待接收方接收即可完成发送过程，并且不会发生阻塞，只有当存储空间满时才会发生阻塞。
	如果缓冲通道中有数据，接收时将不会发生阻塞，知道通道中没有数据可读时，通道将会再度阻塞。
	无缓冲通道保证的是收发过程同步。
	通道实例 := make(chan 通道类型,缓冲大小)
select 关键字：
	当监听的通道没有状态是可读或可写的，select是阻塞的；只要监听的通道中有一个状态是可读或可写的，则select
就不会阻塞，而是进入处理就绪通道的分支流程。如果监听的通道有多个可读或可写的状态，则select随机选取一个处理。
	可以同时响应多个通道的操作。select的每个case都会对应一个通道的收发过程。当收发完成时，就会触发case中
响应的语句。多个操作在每次select中挑选一个进行相应。
	case <-ch:   接收任意数据
	case d:= <-ch 接收变量
	case ch<- 100 发送数据

	1.向已经关闭的通道写数据会导致panic，建议由写入者关闭通道。
	重复关闭通道会导致panic。
	通道关闭后可以读数据。
	func main(){
		c := make (chan struct{})
		ci := make (chan int,100)
		go func (i chan struct{},j chan int){
			for i:=0;i<10;i++{
				ci <- i
			}
			close(ci)
			//写通道
			c <- struct{}{}
		}(c,ci)
		//读通道c,通过通道进行同步等待
		<-c
		//但通道ci还可以继续读取
		for v:= range ci {
			println(v)
		}
	}
	2.阻塞：
	向未初始化的通道写数据或读数据都会导致当前goroutine的永久阻塞
	向缓冲区已满的通道写入数据会导致goroutine阻塞。
	通道中没有数据，读取该通道会导致goroutine阻塞。
	3.非阻塞：
	读取已经关闭的通道不会引发阻塞，而是立即返回通道元素类型的零值，可以使用comma,ok语法判断通道是否已经关闭。

	4.Go语言的future模式：
future模式的基本工作原理：
	使用chan作为函数参数
	启动goroutine调用函数
	通过chan传入参数
	做其他可以并行处理的事情。
	通过chan异步获取结果。
	type query struct{
		sql chan string //参数Channel
		result chan string //结果channel
	}
	//执行query
	func execQuery(q query){
		//启动线程
		go func(){
			//获取输入
			sql := <-q.sql
			//访问数据库
			//输出结果通道
			q.result <- "result from" +sql
		}()
	}
	func main(){
		//初始化Query
		q :=query{make(chan string,1),make(chan string,1)}
		go execQuery(q)//注意执行的时候无须准备参数
		q.sql <- "select * from table"
		//做其他事情，通过sleep描述
		time.Sleep(1*time.Second)
		//获取结果
		fmt.Println(<-q.result)
	}
	//执行结果
	result from select * from table

 */
/**
互斥锁 sync.Mutex——保证同时只有一个goroutine可以访问共享资源
读写互斥锁sync.RWMutex——在读比写多的环境下比互斥锁更高效，在读多写少的环境中
优先使用。
var (
	count int
	CountGuard sync.RWMutex
)
func GetCount int{
	countGuard.RLock()
	defer countGuard.RUnlock()
	return count
}
sync.WatiGroup——保证在并发环境中完成指定数量的任务
等待组的方法：
(wg *WaitGroup)Add(delta int) 等待组的计数器+1
(wg *WaitGroup)Done()         等待组的计数器-1
(wg *WaitGroup)Wait()         当等待组计数器不等于0时阻塞直到变0
 */

/**
反射：
	反射是指在程序运行期对程序本身进行访问和修改的能力。程序在编译时，变量被转换为内存地址，
变量名不会被编译器写入到可执行部分。在程序运行时，程序无法获取自身的信息。支持反射的语言可以
在程序编译期将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序
提供接口访问反射信息，这样就可以在程序运行期获取类型的反射信息，并且有能力修改它们。
type dog struct{
	LegCount int
}
valueOfDog := reflect.ValueOf(&dog{}) //获取dog实例地址的反射值对象
valueOfDog = valueOfDog.Elem() //取出dog实例地址的元素
vLegCount := valueOfDog.FieldByName("LegCount") //获取legCount字段的值
vLegCount.SetInt(4)
使用反射调用函数：
func add(a,b int) int {
	return a+b
}
func main(){
	funcValue := reflect.ValueOf(add) //将函数包装为反射值对象
	//构造函数参数，传入两个整型值
	paramList := []reflect.Value{reflect.ValueOf(10),reflect.ValueOf(20)}
	//反射调用函数
	retList := funcValue.Call(paramList)
	//获取第一个返回值，取整数值
	fmt.Println(retList[0].Int())
}
 */

/**
 * 建议结构体自定义错误处理函数
func (r *RabbitMQ) failOnErr(err error,message string){
   if err != nil {
	log.Fatalf(format:"%s:%s",message,err)
	panic(fmt.Sprintf(format:"%s:%s",message,err))
   }
}

调用示例：
func NewRabbotMQSimple(queueName string) *RabbitMQ{
   rabbitmq := NewRabbitMQ(queueName,exchange:"",key:"")
   var err error
   rabbitmq.conn,err=amqp.Dial(rabbitmq.Mqurl)
   rabbitmq.failOnErr(err,message:"创建连接错误!")
   rabbitmq.channel,err=rabbitmq.conn.Channel()
   rabbitmq.failOnErr(err,message:"获取channel失败")
   return rabbitmq
}
 */

/**
 * 1.属于main代码包，且该文件代码中包含无参数声明和结果声明的main函数，则为命令源码文件，可用go run 命令
 *   直接启动运行。
 * 2.同一个代码包中的所有源码文件，其所属代码包的名称必须一致。如果命令源码文件和库源码文件处于同一个代码包
 *   中，那么在该包中就无法正确执行go build 和 go install 命令。也就是说，这些源码文件将无法通过常规方法
 *   编译和安装。因此，命令源码文件通常会单独放在一个代码包中。
 * 3.同一个代码包中可以有多个命令源码文件，可通过go run 命令分别运行，但这会使go build 和 go install命令
 *   无法编译和安装该代码包。所以，我们也不应该把多个命令源码文件放在同一个代码包中。
 * 4.当代码包中有且仅有一个命令源码文件时，在文件所在目录中执行go build命令，即可在该目录下生成一个与目录
 *   同名的可执行文件；而若使用go install命令，则可在当前工作区的bin目录下生成相应的可执行文件。
 * 5.只有当环境变量GOPATH中只包含一个工作区的目录路径时，go install命令才会把命令源码文件安装到当前工作区的
 *   bin目录下；否则，像这样执行go install 命令就会失败。此时必须设置环境变量GOBIN,该环境变量的值是一个目录
 *   的路径，该目录用于存放所有因安装Go命令源码文件而生成的可执行文件。
 * 6.库源码文件声明的包名会与它直接所属的代码包(目录)名一致，且库源码文件中不包含无参数声明和无结果声明的main
 *   函数。
 * 7.go build:编译指定的代码包或Go语言源码文件。命令源码文件会被编译成可执行文件，并放到命令执行的目录或指定
 *   目录下。而库源码文件被编译后，则不会在非临时目录中留下任何文件。
 *   go clean:用于清除因执行其他go命令而遗留下来的临时目录和文件。
 *   go fmt:用于格式化指定代码包中的Go源码文件。
 *   go get:用于下载、编译并安装指定的代码包及其依赖包。
 *   go install:用于编译并安装指定的代码包及其依赖包。安装命令源码文件后，代码包所在的工作区目录的bin子目录，
 *   或者当前环境变量GOBIN指向的目录中会生成相应的可执行文件。而安装库源码文件后，会在代码包所在的工作区目录
 *   的pkg子目录中生成相应的归档文件。
 *   go list：显示指定代码包的信息。
 *   go run：编译并运行指定的命令源码文件。
 *   go test:用于测试指定的代码包，前提是该代码包目录中必须存在测试源码文件。
 *   命令常用附带参数：
 *   -a 用于强行重新编译所有涉及的Go语言代码包(包括Go语言标准库中的代码包)，即使它们已经是最新的了。该标记可以
 *   让我们有机会通过改动底层的代码包来做一些实验
 *   -n 是命令仅打印其执行过程中用到的所有命令，而不真正执行它们。如果只想查看或验证命令的执行过程，而不想改变任何
 *   东西，使用它正合适。
 *   -race 用于检测并报告指定Go语言程序中存在的数据竞争问题。当用Go语言编写并发程序时，这是很重要的检测手段之一
 *   -v 用于打印命令执行过程中涉及的代码包。这一定包括我们指定的目标代码包，并且有时还会包括该代码包直接或间接
 *   依赖的那些代码包。这会让你知道哪些代码会被命令处理过了。
 *   -x 是命令打印其执行过程中用到的所有命令，同时执行它们。
 *   特殊工具：
 *   pprof:用于以交互的方式访问一些性能概要文件。
 *   trace:读取Go程序踪迹文件，并以图形化的方式展示出来。能让我们深入了解Go程序在运行过程中的内部情况。比如
 *   当前进程中堆的大小及使用情况。又比如，程序中的多个goroutine是怎样被调度的，以及他们在某个时刻被调度的原因。
 *   
 */

/**
 * Go的线程实现模型
 * M:machine的缩写，一个M代表一个内核线程，或称“工作线程”
 * P:processor的缩写，一个P代表执行一个Go代码片段所必需的资源(或称"上下文环境")
 * G:goroutine的缩写，一个G代表一个Go代码片段，前者是对后者的一种封装。
 * 简单来说，一个G的运行需要P和M的支持，一个M在与一个P关联之后，就形成了一个有效的G运行环
 * 境(内核线程+上下文环境)。每个P都会包含一个可运行的G的队列，该队列中的G会被依次传递给与
 * 本地P关联的M,并获得运行时机。
 * Go运行时系统可以把一个M和一个G锁定在一起。一旦锁定，这个M就只能运行这个G，这个G也只能由
 * 该M运行。P是G能够在M中运行的关键。Go的运行时系统会适时地让P与不同的M建立或断开关联，以使
 * P中的那些可运行的G能够及时获得运行时机
 *
 * 主goroutine的运作：
 *    封装main函数的goroutine称为主goroutine。主gouroutine会由runtime.m0负责运行。
 *    主gouroutine设定每一个goroutine所能申请的栈空间的最大尺寸。32位计算机系统中此
 * 最大尺寸为250MB,而在64位的此尺寸为1GB。如果某个goroutine的栈空间尺寸大于这个这个
 * 限制，那么运行时系统就会发起一个栈溢出(stack overflow)的运行时恐慌。随机，这个Go
 * 程序的运行也会终止。
 *    在设定好goroutine的最大栈尺寸之后，主goroutine会在当前M的g0上执行系统监测任务。
 * 系统监测任务的作用就是为调度器查漏补缺，这也是让系统监测任务的执行先于main函数的原因
 * 之一。
 *    此后，主goroutine会进行一系列的初始化工作，涉及的工作大致如下：
 *    	1.检查当前M是否是runtime.m0。如果不是，就说明之前的程序出现了某种问题。这时，主goroutine
 *     	会立即抛出异常，这也意味着Go程序启动的失败。
 *      2.创建一个特殊的defer语句，用于在主goroutine退出时做必要的善后处理。因为主goroutine也可能
 *      正常地结束，所以这一点很有必要。
 *      3.启用专用于在后台清扫内存垃圾的goroutine,并设置GC可用的标识。
 *      4.执行main包中的init函数。
 * 如果上述初始化工作成功完成，那么主goroutine就会去执行main函数。在执行完main函数之后，它还会检查
 * 主goroutine是否引发了运行时恐慌，并进行必要的处理。最后，主goroutine会结束自己以及当前进程的运行。
 * 在mian函数执行期间，运行时系统会根据Go程序中的go语句，复用或新建goroutine来封装go函数。这些goroutine
 * 都会放入相应P的可运行G队列中，然后等待调度器的调度。
 * 
 * 
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
