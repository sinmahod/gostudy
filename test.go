// test.go
package main

import (
	"fmt"
	"time"
	"unsafe"
)

var a int = 1
var b string = "你好"
var c bool

//一次声明多个int变量
var d, e, f int

//因式分解形式声明全局变量
var (
	g string
	h float32
	i int
)

//iota ,自增值 同一个()内才生效,下面的如果不初始化值的话就同上面的
const (
	C = iota  //0
	D         //1
	E         //2
	F = "asd" //asd
	G         //asd
	H = iota  //5
	I         //6
)

//定义函数（自定义类型）
type Circle2 struct {
	rad int
}

func main() {
	fmt.Println("Hello World!", a, b, c)

	//给多个变量赋值
	d, e, f = 1, 2, 3
	fmt.Println(d, e, f)
	//相同类型才可以计算，不然就必须转换类型，否则报错
	fmt.Println(d + e + f)
	//得到当前系统时间
	fmt.Println(time.Now())

	h = 10.223
	i = 1
	//强制类型转换
	fmt.Println(h + float32(i))

	//局部变量可以用:=的方式声明，不指定类型的话编译器会根据赋值自动指定类型
	j, k := 123, "这是一个string"

	//&得到变量的内存地址（跟值没关系）
	fmt.Println(&j, k)
	fmt.Println(j == 123)

	i = j
	fmt.Println(i)

	j = 2
	l := 100
	fmt.Println(i, l)
	//交换两个变量的值
	i, j = j, i
	fmt.Println(i, j)

	//_, n = getTwoParams()   这种写法是调用一个函数，这个函数返回两个值，但是其中有一个值你不想要的时候可以抛弃他（因为Go规定得到的局部变量必须要使用，否则报错）

	//定义常量，常量可以不被代码调用
	const A = "常量"
	fmt.Println(a, A)

	const LENGTH, WIDTH = 10, 5
	var area int
	area = LENGTH * WIDTH
	//格式化输出
	fmt.Printf("面积为：%d\n", area)
	//计算字符串长度len()
	fmt.Println(len(A))
	fmt.Println(A, len(A))

	//计算字符串元素个数
	const B = unsafe.Sizeof("AA")
	fmt.Println(B)

	//测试自增
	fmt.Println(C, D, E, F, G, H, I)

	fmt.Printf("得到变量类型：%T\n", A)
	fmt.Printf("得到变量类型：%T\n", B)
	fmt.Printf("得到变量类型：%T\n", C)

	var prt *int //声明指针类型
	prt = &area  //将area的指针位置赋给prt
	fmt.Println("得到变量的指针：", *prt)
	area = 33 //变量area的值改变了那么指针的值就会改变
	fmt.Println("得到变量的指针：", *prt)

	//尝试条件语句if else
	if len(A) > 5 {
		fmt.Println(len(A))
	} else {
		fmt.Println(2)
	}

	//尝试条件语句switch
	switch i {
	case 1, 2, 31:
		fmt.Println(1)
		break
	case 3, 4, 5:
		fmt.Println(2)
		break
	default:
		fmt.Println(3)
	}

	//尝试条件语句select
	var c1, c2, c3 chan int
	var i1, i2 int

	select {
	case i1 = <-c1: //这个不懂，如何使这个成立呢？
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}

	//for循环使用1 类似java for i=2
	for x := 0; x < i; x++ {
		fmt.Println("for循环普通用法", x)
	}
	//for循环使用2 类似java while
	x := 0
	for x < i {
		fmt.Println("for循环while用法", x)
		x++
	}
	//for循环使用3 for each  map用法
	var mp map[string]string = make(map[string]string)
	mp["key1"] = "v1"
	mp["key2"] = "v2"
	//记住关键字range
	for key := range mp {
		fmt.Println(key, mp[key])
	}
	//for循环 无限循环
	for {
		if i > 3 {
			break //还有goto 、continue
		}
		fmt.Println("for循环无限循环", i)
		i++
	}

	//调用测试环境
	fmt.Println("调用函数测试", test(1))
	fmt.Println("调用函数测试", test(10))
	x1, x2 := test2(3, 5)
	x3, x4 := test2(5, 3)
	fmt.Println("调用双返回函数测试", x1, x2)
	fmt.Println("调用双返回函数测试", x3, x4)

	//测试匿名函数，如果不进行声明的话只能得到函数的内存地址
	functest := test3()
	fmt.Println("调用函数闭包，匿名函数", functest())
	fmt.Println("调用函数闭包，匿名函数", functest())
	fmt.Println("调用函数闭包，匿名函数", functest())

	//调用方法
	var cs Circle2
	cs.rad = 100
	fmt.Println("调用方法", cs.getC())
}

//创建函数 接受Int参数返回string
func test(i int) bool {
	if i < 5 {
		return true
	} else {
		return false
	}
}

//创建双返回函数 接受Int参数返回string和int
func test2(i, j int) (bool, int) {
	if i < 5 {
		return true, i + j
	} else {
		return false, i - j
	}
}

//创建函数闭包，测试匿名函数，
func test3() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//创建方法
func (c Circle2) getC() int {
	return c.rad + 10
}
