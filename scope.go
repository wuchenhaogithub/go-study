package main

import "fmt"

/**
作用域 全局变量 局部变量 形式参数
局部变量可以和全局变量相同，局部变量优先
*/

var x int //声明全局变量，可以在整个包甚至包外使用

func main() {
	var a, b, c int //声明局部变量，只作用于函数体内
	a, b, c = 1, 2, 3
	x = a + b + c
	fmt.Printf("a = %d  b = %d  c = %d x = %d\n", a, b, c, x)
}
