package main

import "fmt"

func getConst()  {
	const LENGTH = 10
	const WIDTH = 5

	const (
		Unknown = 0
		Female = 1
		Male = 2
	)

	var area int
	const a,b,c =1 ,false,"str" //多重赋值
	area = LENGTH * WIDTH
	fmt.Printf("面积为 ：%d",area)
	println()
	println(a,b,c)
}

func main() {
	getIota()
}
/**
iota  特殊的常量，可以认为是一个可以被编译器修改的常量
iota在const关键字出现的时候将被重置为0 ，const中每新增一行变量声明将使iota计数一次，可以理解为const语句块中的行索引
 */
func getIota() {
	const (
		i=1<<iota
		j=3<<iota
		k
		l
	)
	fmt.Println("i=",i)
	fmt.Println("j=",j)
	fmt.Println("k=",k)
	fmt.Println("l=",l)
}
/**
	iota 表示从0 开始自动加1，所以i=1<<0,j=3<<1 即 i = 1 j = 6,
		i=1：左移 0 位,不变仍为 1;
		j=3：左移 1 位,变为二进制 110, 即 6;
		k=3：左移 2 位,变为二进制 1100, 即 12;
		l=3：左移 3 位,变为二进制 11000,即 24
 */








