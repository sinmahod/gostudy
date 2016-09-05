// test.go
package main

import (
	"errors"
	"fmt"
	"gostudy/testa2"
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

//声明数组
var s1 [10]int
var s2 = []int{19, 18, 17, 16}

//多维数组
var s3 = [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}

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

	c1 = make(chan int, 1) //创建一个
	c1 <- 10

	select {
	case i1 = <-c1:
		fmt.Printf("received %d from c1\n", i1)
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
	close(c1) //chan使用完毕需要关闭
	fmt.Println(1)
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
	//判断元素是否存在（判断key是否存在） //v 如果存在则得到key1的值，ok 如果存在则返回true否则false
	var v, ok = mp["key1"]
	fmt.Println(v, ok)
	//删除mp中key=key1的元素
	delete(mp, "key1")
	//记住关键字range
	for key, value := range mp {
		fmt.Println(key, mp[key])
		fmt.Println("Map的value遍历", value)
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

	//测试数组
	for ss := 0; ss < len(s1); ss++ {
		s1[ss] = ss + 5
	}
	for _, ss := range s1 {
		fmt.Println(ss)
	}
	fmt.Println(s2)
	for _, sss := range s2 {
		fmt.Println(sss)
	}
	for ssss := 0; ssss < len(s2); ssss++ {
		fmt.Println(s2[ssss])
	}

	//测试多维数组
	fmt.Println("多维数组s3[2][2]", s3[2][2])

	//测试数组函数
	sspar := sum(s2)
	fmt.Printf("s2的数组之和是%d\n", sspar)

	//测试空指针
	var ptr *int
	fmt.Println("空指针ptr", ptr)
	fmt.Printf("空指针ptr的值为%x\n", ptr)
	fmt.Println("ptr是否为空指针", ptr == nil)

	//测试指向指针的指针
	var ptrss ***int
	var ptrs **int
	var a = 10
	ptr = &a
	fmt.Println(ptr, *ptr)
	ptrs = &ptr
	fmt.Println(ptrs, *ptrs, **ptrs)
	ptrss = &ptrs
	fmt.Println(ptrss, *ptrss, **ptrss, ***ptrss)

	//测试指针函数
	var pa, pb int = 3, 9
	swap(&pa, &pb)
	fmt.Println("pa与pb交换后为", pa, pb)

	//测试结构体
	var pro Prople
	pro.age = 20
	pro.gender = 1
	pro.height = 180
	pro.weight = 62

	fmt.Printf("这个人的性别是%s,年龄是%d,身高是%d,体重是%dkg\n", getGender(pro.gender), pro.age, pro.height, pro.weight)

	//切片测试  5=capacity容量，用cap(slice)获取，切片容量是自增的，超过切片指定容量则增大一倍容量，如果没有指定容量则增加2容量
	var slice = make([]int, 3, 5)
	var slice2 = []int{11, 22, 33, 44, 55}
	fmt.Println("slice=", slice)
	fmt.Println(cap(slice))
	fmt.Println(slice2[4:5])
	slice = append(slice, 1, 1, 1, 1, 1, 1)
	fmt.Println("slice=", slice)
	fmt.Println(cap(slice), len(slice))
	//切片copy
	var slice3 = make([]int, len(slice), cap(slice)*2)
	copy(slice3, slice) //将slice的值copy到slice3
	fmt.Println(cap(slice3), len(slice3), slice3)

	//遍历字符串，得到unicode的值
	for idx, val := range "go" {
		fmt.Println(idx, val)
	}

	fmt.Println(getSum(101))
	fmt.Println(testa2.Testa(8, 9))

	//测试接口
	var ip Phone = new(IPhone)
	ip.num()
	ip = new(Xiaomi)
	ip.num()

	//测试错误
	int2, err2 := sqrt(1)
	fmt.Printf("调用sqrt函数，如果不大于0则返回错误，当前传递的值是%d，返回结果为%d和%s\n", 1, int2, err2)

	//测试错误使用
	if det1, err3 := testa2.Detail(1, 2); err3 == "" {
		fmt.Println("正确：", det1)
	} else {
		fmt.Println("错误：", err3)
	}
}

//定义一个返回错误的函数
func sqrt(n int) (int, error) {
	if n > 0 {
		return n - 1, errors.New("错误，传递的数值必须大于0")
	} else {
		return 0, errors.New("错误，传递的数值必须大于0")
	}
}

//定义一个接口Phone
type Phone interface {
	num()
}

//定义一个结构
type IPhone struct {
}

//定义一个方法
func (iphone IPhone) num() {
	fmt.Println("这个一个定义的方法")
}

//定义第二个结构
type Xiaomi struct {
}

//定义第二个方法
func (xm Xiaomi) num() {
	fmt.Println("这是第二个定义的方法")
}

//测试结构体（类似java的自定义对象）
type Prople struct {
	gender int
	age    int
	height int
	weight int
}

//创建函数 接受Int参数返回string
func test(i int) bool {
	if i < 5 {
		return true
	} else {
		return false
	}
}

func getGender(i int) (gender string) {
	if i == 1 {
		gender = "男"
	} else {
		gender = "女"
	}
	return
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

//创建接受数组返回数组值和的函数
func sum(sz []int) int {
	sumpar := 0
	for _, ssz := range sz {
		sumpar += ssz
	}
	return sumpar
}

//创建指针函数
func swap(ptra *int, ptrb *int) {
	//ptra, ptrb = ptrb, ptra  //这样交换不会改变实际的值
	*ptra, *ptrb = *ptrb, *ptra
}

//递归函数编写：1-n求和
func getSum(n int) int {
	//i := (n+1)*(n/2) + (n%2)*((n+1)/2)
	//fmt.Println(n+1, "*", n/2, "+", (n % 2), "*", (n+1)/2)
	if n == 0 {
		return 0
	}
	i := n + getSum(n-1)
	return i
}
