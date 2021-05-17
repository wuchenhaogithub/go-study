package main

import "fmt"

/**
类型数据转换
将一种数据类型的变量转换为另一种类型的变量
语法 type_name(expression)
type_name 为类型，expression 为表达式。
*/

func main() {
	var sum int
	var count float64
	var mean float64
	sum = 15
	count = 0.5
	mean = float64(sum) / count
	fmt.Println(mean)
}
