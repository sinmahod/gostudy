package testa2

import (
	"fmt"
)

func Testa(a int, b int) int {
	return a + b
}

//创建一个结构
type DetailError struct {
	par1 int
	par2 int
}

//创建一个开放的方法
func (de *DetailError) Error() (err string) {
	formaterr := "这是一个错误提醒测试，%d."
	return fmt.Sprintf(formaterr, de.par1)
}

//创建一个函数判断传递的值
func Detail(p1 int, p2 int) (p int, err string) {
	if p1 == 0 {
		//创建一个DetailError类型的变量并为变量赋值
		de := DetailError{
			par1: p1,
			par2: p2, //注意这里的逗号
		}
		return 0, de.Error()
	} else {
		return p1, ""
	}
}
