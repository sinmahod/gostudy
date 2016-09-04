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
}
